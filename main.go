package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/vitaLemoTea/myBank/api"
	"github.com/vitaLemoTea/myBank/config"
	db "github.com/vitaLemoTea/myBank/db/sqlc"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configs:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("cannot connect to db :%v", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server :", err)
	}
}
