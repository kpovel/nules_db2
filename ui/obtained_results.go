package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type RegionMaxPM struct {
	Title    string
	City     string
	MaxValue float64
}

type DailyHarmfulPM2 struct {
	Day                      string
	AvgValue                 float64
	HarmfulMeasurementsCount int
}

type HarmedElement struct {
  Day      string
  AvgValue float64
}

func (app *App) ObtainedResults(w http.ResponseWriter, r *http.Request) {
	formValue := r.FormValue("select-data")

	if formValue == "max-pm" {
		query, err := app.DB.Query(`
select mu.title     as measured_title,
       s.city,
       max(m.value) as max_value
from measurement m
         inner join measured_unit mu on m.id_measured_unit = mu.id_measured_unit
         inner join station s on m.id_station = s.id_station
where (mu.title = 'PM2.5'
    or mu.title = 'PM10')
  and m.time between '2022-02-26' and '2023-04-05'
group by mu.title, s.city;`)

		if err != nil {
			fmt.Fprintf(w, "Query error: %v", err)
			log.Println(err)
			return
		}

		defer query.Close()

		var regionMaxPM []RegionMaxPM

		for query.Next() {
			var maxPM RegionMaxPM

			err := query.Scan(&maxPM.Title, &maxPM.City, &maxPM.MaxValue)

			if err != nil {
				fmt.Fprintf(w, "Query error: %v", err)
				log.Println(err)
				return
			}

			regionMaxPM = append(regionMaxPM, maxPM)
		}

		template := template.Must(template.ParseFiles("templates/region-max-pm.html"))
		template.Execute(w, regionMaxPM)
	} else if formValue == "daily-harmful-pm2" {
		query, err := app.DB.Query(`select date_trunc('day', m.time)                              as day,
       avg(m.value)                                           as avg_value,
       count(case when ov.id_measured_unit >= '4' then 1 end) as harmful_measurements_count
from measurement m
         inner join optimal_value ov on m.id_measured_unit = ov.id_measured_unit
         inner join category c on c.id_category = ov.id_category
where ov.id_category >= '4' -- PM2.5
  and m.id_station = '0002'
group by day
order by day;`)

		if err != nil {
			fmt.Fprintf(w, "Query error: %v", err)
			log.Println(err)
			return
		}

		defer query.Close()

		var dailyHarmfulPM2 []DailyHarmfulPM2

    for query.Next() {
      var dailyPM2 DailyHarmfulPM2

      err := query.Scan(&dailyPM2.Day, &dailyPM2.AvgValue, &dailyPM2.HarmfulMeasurementsCount)

      if err != nil {
        fmt.Fprintf(w, "Query error: %v", err)
        log.Println(err)
        return
      }

      dailyHarmfulPM2 = append(dailyHarmfulPM2, dailyPM2)
    }

    template := template.Must(template.ParseFiles("templates/daily-harmful-pm2.html"))
    template.Execute(w, dailyHarmfulPM2)
	} else if formValue == "sulfur-dioxide" {
    query, err := app.DB.Query(`select date_trunc('day', m.time) as day, avg(m.value) as avg_value
from measurement m
where m.id_measured_unit = '15' -- Sulfur dioxide
group by day
order by day, avg_value;`)

    if err != nil {
      fmt.Fprintf(w, "Query error: %v", err)
      log.Println(err)
      return
    }

    defer query.Close()

    var sulfurDioxide []HarmedElement

    for query.Next() {
      var sulfur HarmedElement

      err := query.Scan(&sulfur.Day, &sulfur.AvgValue)

      if err != nil {
        fmt.Fprintf(w, "Query error: %v", err)
        log.Println(err)
        return
      }

      sulfurDioxide = append(sulfurDioxide, sulfur)
    }

    template := template.Must(template.ParseFiles("templates/harmed-element.html"))
    template.Execute(w, sulfurDioxide)
  } else if formValue == "carbon-monoxide" {
    query, err := app.DB.Query(`select date_trunc('day', m.time) as day, avg(m.value) as avg_value
from measurement m
where m.id_measured_unit = '12' -- Carbon monoxide(CO)
group by day
order by day, avg_value;`)

    if err != nil {
      fmt.Fprintf(w, "Query error: %v", err)
      log.Println(err)
      return
    }

    defer query.Close()

    var carbonMonoxide []HarmedElement
    for query.Next() {
      var carbon HarmedElement

      err := query.Scan(&carbon.Day, &carbon.AvgValue)

      if err != nil {
        fmt.Fprintf(w, "Query error: %v", err)
        log.Println(err)
        return
      }

      carbonMonoxide = append(carbonMonoxide, carbon)
    }

    template := template.Must(template.ParseFiles("templates/harmed-element.html"))
    template.Execute(w, carbonMonoxide)
  }
}
