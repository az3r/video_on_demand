package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func InitDatabase() {
	log.Print("Connecting to database...")

	conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/school")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	Conn = conn
}
