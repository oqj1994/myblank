package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testqueries *Queries
var dbDriver = "postgres"
var dbSource = "postgresql://root:123456@localhost:5432/simple_bank?sslmode=disable"
var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDb, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	testqueries = New(testDb)
	os.Exit(m.Run())
}
