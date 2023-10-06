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

type distinct_stations struct {
	Name string
}

type optimal_value struct {
	Title         string
	Designation   string
	Unit          Units
	Bottom_border uint
	Upper_border  *uint
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
		defer station_query.Close()
		if err != nil {
			fmt.Fprintf(w, "Error: %v", err)
			log.Fatal(err)
			return
		}

		var stations []station

		for station_query.Next() {
			var station station
			if err := station_query.Scan(&station.Name, &station.City); err != nil {
				fmt.Println(w, "Error: %v", err)
				log.Fatal(err)
				return
			}
			stations = append(stations, station)
		}

		station_template := template.Must(template.ParseFiles("./templates/stations.html"))
		station_template.Execute(w, stations)

	case "distinct-stations":
	case "station-info":
	case "obtained-parameters":
	case "optival-value":
	default:
		unknown_select := "Unknown select-data"
		fmt.Fprintf(w, unknown_select)
		log.Printf(unknown_select)
	}
}
