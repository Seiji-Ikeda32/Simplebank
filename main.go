package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/Seiji-Ikeda32/simplebank/api"
	db "github.com/Seiji-Ikeda32/simplebank/db/sqlc"
	"github.com/Seiji-Ikeda32/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("connot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("connot create server")
	}

	err = server.Start(config.ServerAddress)
	println("server is Running")
	if err != nil {
		log.Fatalln("cannot start server:", err)
	}
}
