package main

import (
	"database/sql"
	"fmt"
	"log"

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

	servers := app.select_servers()

	for _, server := range servers {
		fmt.Printf("ID: %d, URL: %s, Status: %s\n", server.id_server, server.url, server.status)
	}
}
