package main

import (
	"database/sql"
	"log"

	"fidelis.com/simple_bank/api"
	db "fidelis.com/simple_bank/db/sqlc"
	"fidelis.com/simple_bank/util"
	_ "github.com/lib/pq"
)



func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err.Error())
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err.Error())
	}
	
	store := db.NewStore(conn)
	server := api.NewServer(store)

	if err != nil {
		log.Fatal(err.Error())
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}
}
