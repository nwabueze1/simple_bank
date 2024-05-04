package main

import (
	"database/sql"
	"fidelis.com/simple_bank/api"
	db "fidelis.com/simple_bank/db/sqlc"
	_ "github.com/lib/pq"
	"log"
)

const (
	driverName     = "postgres"
	dataSourceName = "postgresql://postgres:root@localhost:5432/simple_bank?sslmode=disable"
	serverAddress  = "0.0.0.0:8000"
)

func main() {
	conn, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	if err != nil {
		log.Fatal(err.Error())
	}
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}
}
