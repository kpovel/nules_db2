package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type MQTT_Server struct {
	id_server uint
	url       string
	status    string
}

func main() {
	connStr := "postgres://postgres:12345678@localhost:5432/eco-station?sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select * from MQTT_Server")

	defer rows.Close()

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

	for _, server := range servers {
		fmt.Printf("ID: %d, URL: %s, Status: %s\n", server.id_server, server.url, server.status)
	}
}
