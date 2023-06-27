package main

import (
	"context"
	"fmt"
	"log"
	"os/user"

	"github.com/jackc/pgx/v5"
	"github.com/matrosovm/warehouse/internal/pkg/database"
)

const (
	host   = "localhost"
	port   = 5432
	dbname = "db_warehouse"
)

func dbConnect(ctx context.Context) database.Store {
	usr, err := user.Current()
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, usr.Username, usr.Username, dbname,
	)

	conn, err := pgx.Connect(ctx, psqlInfo)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	return database.NewStore(conn)
}
