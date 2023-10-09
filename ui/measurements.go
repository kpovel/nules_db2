package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type MeasuredStats struct {
	Value          float64
	Title          string
	Time           string
	IdMeasuredUnit uint
	MinValue       float64
	MaxValue       float64
	AvgValue       float64
}

func (app *App) measurements(w http.ResponseWriter, r *http.Request) {
  stationID := r.FormValue("stationID")
  startDate := r.FormValue("start-date")
  endDate := r.FormValue("end-date")

  if stationID == "" || startDate == "" || endDate == "" {
    fmt.Fprintf(w, "Error: stationID, startDate and endDate are required")
    log.Println("Error: stationID, startDate and endDate are required")
    return
  }

	measurement_query, err := app.DB.Query(`with measured_units as (select * from measured_unit)
  select m.value,
  mu.title,
  m.time,
  m.id_measured_unit,
  last_value(value) over (w_desc) as min_value,
  first_value(value) over (w_asc) as max_value,
  avg(value) over (w_asc)         as avg_value
  from measurement m left join measured_units mu on mu.id_measured_unit = m.id_measured_unit
  where id_station = $1 and time between $2 and $3
  window w_desc as (partition by id_measurement order by value desc),
  w_asc as (partition by id_measurement order by value)
  order by time desc, value
  limit 100;`, stationID, startDate, endDate)

	if err != nil {
		fmt.Fprintf(w, "Query error: %v", err)
		log.Println(err)
		return
	}

	defer measurement_query.Close()

	var measurements []MeasuredStats

	for measurement_query.Next() {
		var measurement MeasuredStats
		err = measurement_query.Scan(&measurement.Value, &measurement.Title, &measurement.Time, &measurement.IdMeasuredUnit, &measurement.MinValue, &measurement.MaxValue, &measurement.AvgValue)

		if err != nil {
			fmt.Fprintf(w, "Scanning error: %v", err)
			log.Println(err)
			return
		}

		measurements = append(measurements, measurement)
	}

	template := template.Must(template.ParseFiles("templates/measurements.html"))
	template.Execute(w, measurements)
}
