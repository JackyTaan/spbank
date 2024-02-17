package main

import (
	"database/sql"
	"log"

	"github.com/JackyTaan/spbank/api"
	db "github.com/JackyTaan/spbank/db/sqlc"
	"github.com/JackyTaan/spbank/util"
	_ "github.com/lib/pq"
)

func main() {
	///Load config
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can not load configuration", err)
	}
	///
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
