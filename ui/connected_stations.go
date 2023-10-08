package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type ConnectedStation struct {
	IdStation        uint
	City             string
	Name             string
	FirstMeasured    string
	LastMeasurements string
}

func (app *App) connected_stations(w http.ResponseWriter, r *http.Request) {
	connected_stations_query, err := app.DB.Query(`
  select station.id_station, station.city, station.name, mv_first_measurements.first_measured, mv_last_measurements.last_measurements
  from station
  left join mv_first_measurements on station.id_station = mv_first_measurements.id_station
  left join mv_last_measurements on station.id_station = mv_last_measurements.id_station
  where status = 'enabled'
  order by station.id_station;`)

	if err != nil {
		fmt.Fprintf(w, "Error due selecting connected stations: %v", err)
		log.Println(err)
		return
	}

	var connected_stations []ConnectedStation

	for connected_stations_query.Next() {
		var ct ConnectedStation

		if err := connected_stations_query.Scan(&ct.IdStation, &ct.City, &ct.Name, &ct.FirstMeasured, &ct.LastMeasurements); err != nil {
			fmt.Fprintf(w, "Error due to scanning a table: %v", err)
			log.Println(err)
			return
		}

		connected_stations = append(connected_stations, ct)
	}

	cs_template := template.Must(template.ParseFiles("./templates/connected_stations.html"))
	cs_template.Execute(w, connected_stations)
}
