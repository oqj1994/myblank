package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/vitaLemoTea/myBank/config"
)

var testqueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error

	config, err := config.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	testDb, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	testqueries = New(testDb)
	os.Exit(m.Run())
}
