package modules

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

var Db *pgx.Conn

func InitDatabase() {
	log.Print("Connecting to database...")

	conn, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/school")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	query, err := conn.Query(context.Background(), "select id, name from account_entity")
	if err != nil {
		log.Fatalf("Failed to query: %v", err)
	}

	var id string
	var name string

	if query.Next() {
		query.Scan(&id, &name)
		log.Printf("find account: id %s, name %s", id, name)
	}

	Db = conn
}
