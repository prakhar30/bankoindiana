package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prakhar30/bankoindiana/api"
	db "github.com/prakhar30/bankoindiana/db/sqlc"
)

const (
	dbSource      = "postgresql://root:something_secret@localhost:5432/banko_indiana?sslmode=disable"
	serverAddress = "localhost:8080"
)

func main() {
	connPool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(connPool)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
