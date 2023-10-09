package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)


type StationL struct {
  ID uint
  Name string
}

func (app *App) stations(w http.ResponseWriter, r *http.Request) {
  stations_query, err := app.DB.Query("select id_station, name from station;")

  if err != nil {
    fmt.Fprintf(w, "Error during data selection: %v", err)
    log.Println(err)
    return
  }

  var stations []StationL

  for stations_query.Next() {
    var station StationL

    if err := stations_query.Scan(&station.ID, &station.Name); err != nil {
      fmt.Fprintf(w, "Error during scanning stations: %v", err)
      log.Println(err)
      return
    }

    stations = append(stations, station)
  }

  template := template.Must(template.ParseFiles("templates/select_station.html"))
  template.Execute(w, stations)
}
