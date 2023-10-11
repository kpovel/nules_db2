package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type ConnectedStation struct {
	IdStation       uint
	City            string
	Name            string
	EnabledFrom     string
	MeasuredTitle   string
	MeasuredUnit    string
	LastMeasurement string
}

func (app *App) connected_stations(w http.ResponseWriter, r *http.Request) {
	connected_stations_query, err := app.DB.Query("select * from connected_stations;")

	if err != nil {
		fmt.Fprintf(w, "Error due selecting connected stations: %v", err)
		log.Println(err)
		return
	}

	var connected_stations []ConnectedStation

	for connected_stations_query.Next() {
		var ct ConnectedStation

		if err := connected_stations_query.Scan(&ct.IdStation, &ct.City, &ct.Name, &ct.EnabledFrom, &ct.MeasuredTitle, &ct.MeasuredUnit, &ct.LastMeasurement); err != nil {
			fmt.Fprintf(w, "Error due to scanning a table: %v", err)
			log.Println(err)
			return
		}

		connected_stations = append(connected_stations, ct)
	}

	cs_template := template.Must(template.ParseFiles("./templates/connected_stations.html"))
	cs_template.Execute(w, connected_stations)
}
