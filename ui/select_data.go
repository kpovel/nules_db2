package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type station struct {
	Name string
	City string
}

type city struct {
	Name string
}

type optimal_value struct {
	Title         string
	Designation   string
	Unit          Units
	Bottom_border uint
	Upper_border  *uint
}

type obtained_parameter struct {
  ID_Measurement uint
  ID_Station uint
  ID_Measured_Unit uint
  Time string
  Value float64
}

func (app *App) select_data(w http.ResponseWriter, r *http.Request) {
	select_data := r.URL.Query().Get("select-data")

	switch select_data {
	case "station-list":
		station_query, err := app.DB.Query(`with stationFirstMeasurment as (select id_station, min(time) as firstTimeMeasurement
                                from measurement
                                group by id_station)
                                select name, city
                                from station
                                join stationFirstMeasurment sfm on station.id_station = sfm.id_station
                                where extract(year from sfm.firstTimeMeasurement) >= 2021;`)
		if err != nil {
			fmt.Fprintf(w, "Error: %v", err)
			log.Println(err)
			return
		}

		defer station_query.Close()

		var stations []station

		for station_query.Next() {
			var station station
			if err := station_query.Scan(&station.Name, &station.City); err != nil {
				fmt.Println(w, "Error due to scanning a table: %v", err)
				log.Println(err)
				return
			}
			stations = append(stations, station)
		}

		station_template := template.Must(template.ParseFiles("./templates/stations.html"))
		station_template.Execute(w, stations)

	case "cities-with-station":
		cities_query, err := app.DB.Query("select distinct city from station;")

		if err != nil {
			fmt.Fprintf(w, "Error: %v", err)
			log.Println(err)
			return
		}
		defer cities_query.Close()

		var cities []city
		for cities_query.Next() {
			var city city
			if err := cities_query.Scan(&city.Name); err != nil {
				fmt.Fprintf(w, "Error due to scanning a table: %v", err)
				log.Println(err)
				return
			}
			cities = append(cities, city)
		}

		cities_table := template.Must(template.ParseFiles("./templates/cities.html"))
		cities_table.Execute(w, cities)

	case "station-info":
		station_query := app.DB.QueryRow("select * from station where coordinates <-> point(35.058606, 48.44803) = 0;")

		var station Station
		if err := station_query.Scan(&station.ID_Station, &station.City, &station.Name, &station.Status, &station.ID_SaveEcoBot, &station.ID_Server, &station.Coordinates); err != nil {
			fmt.Fprintf(w, "Error due to scanning a table: %v", err)
			log.Println(err)
			return
		}

		station_template := template.Must(template.ParseFiles("./templates/station.html"))
		station_template.Execute(w, station)

	case "obtained-parameters":
		measurement_query, err := app.DB.Query(`with KyivsStations as (select id_station from station where city = 'Kyiv')
    select measurement.id_measurement, measurement.id_station, measurement.id_measured_unit, measurement.time, measurement.value
    from measurement join KyivsStations on KyivsStations.id_station = measurement.id_station
    where date(time) between '2022-01-01' and '2022-02-20'
    group by measurement.id_station, measurement.id_measured_unit, measurement.value, measurement.time, measurement.id_measurement
    limit 500;`)

    if err != nil {
      fmt.Fprintf(w, "Error: %v", err)
      log.Println(err)
      return
    }

    var parameters []obtained_parameter

    for measurement_query.Next() {
      var parameter obtained_parameter
      if err := measurement_query.Scan(&parameter.ID_Measurement, &parameter.ID_Station, &parameter.ID_Measured_Unit, &parameter.Time, &parameter.Value); err != nil {
        fmt.Fprintf(w, "Error due to scanning a table: %v", err)
        log.Println(err)
        return
      }

      parameters = append(parameters, parameter)
    }

    parameters_table := template.Must(template.ParseFiles("./templates/obtain_parameters.html"))
    parameters_table.Execute(w, parameters)

	case "optimal-value":
	default:
		unknown_select := "Unknown select-data"
		fmt.Fprintf(w, unknown_select)
		log.Printf(unknown_select)
	}
}
