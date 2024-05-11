package main

import (
	"fmt"
	"log"
	"time"

	"fidelis.com/simple_bank/token"
	"fidelis.com/simple_bank/util"
)

// import (
// 	"database/sql"
// 	"log"

// 	"fidelis.com/simple_bank/api"
// 	db "fidelis.com/simple_bank/db/sqlc"
// 	"fidelis.com/simple_bank/util"
// 	_ "github.com/lib/pq"
// )

// func main() {
// 	config, err := util.LoadConfig(".")
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	conn, err := sql.Open(config.DBDriver, config.DBSource)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	store := db.NewStore(conn)
// 	server := api.NewServer(store)

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// 	err = server.Start(config.ServerAddress)
// 	if err != nil {
// 		log.Fatal("Cannot start server", err)
// 	}
// }

func main() {
	maker, err := token.NewJWTMaker(util.RandomString(32))

	if err != nil {
		log.Fatalln(err)
	}

	username := util.GenerateRandomOwner()
	duration := time.Minute
	token, err := maker.CreateToken(username, -duration)

	if err != nil {
		log.Fatalln(err)
	}

	payload, err := maker.VerifyToken(token)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(payload.ExpiresAt.Before(time.Now()))
}
