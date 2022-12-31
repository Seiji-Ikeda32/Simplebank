package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/Seiji-Ikeda32/simplebank/api"
	db "github.com/Seiji-Ikeda32/simplebank/db/sqlc"
)

const (
	dbDriver     = "postgres"
	dbSource     = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAdress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	defer conn.Close()
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAdress)
	println("server is Running")
	if err != nil {
		log.Fatalln("cannot start server:", err)
	}
}
