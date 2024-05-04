package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"fidelis.com/simple_bank/util"
	_ "github.com/lib/pq"
)

var testQuery *Queries
var testDB *sql.DB


func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")


	if err != nil {
		log.Fatal(err.Error())
	}

	
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err.Error())
	}

	testQuery = New(testDB)

	os.Exit(m.Run()) 
}
