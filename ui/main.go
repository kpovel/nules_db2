package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type App struct {
	DB *sql.DB
}

func (app *App) select_servers() []MQTT_Server {
	rows, err := app.DB.Query("select * from MQTT_Server")
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}

	var servers []MQTT_Server

	for rows.Next() {
		var server MQTT_Server
		if err := rows.Scan(&server.id_server, &server.url, &server.status); err != nil {
			log.Fatal(err)
		}
		servers = append(servers, server)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return servers
}

func main() {
	connStr := "postgres://postgres:12345678@localhost:5432/eco-station?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	app := App{DB: db}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/",http.StripPrefix("/", withoutAuth(fs)))

	protected_fs := http.FileServer(http.Dir("./static/protected"))
	http.Handle("/protected/", http.StripPrefix("/protected", requireAuth(protected_fs)))

	css_fs := http.FileServer(http.Dir("./static/css"))
	http.Handle("/css/",http.StripPrefix("/css", css_fs))

	http.HandleFunc("/htmx.min.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "htmx.min.js")
	})

	http.HandleFunc("/api/login", app.login)
	http.HandleFunc("/api/signup", app.signup)

	log.Print("Listening on :42069")

	err = http.ListenAndServe(":42069", nil)
	if err != nil {
		log.Fatal(err)
	}
}
