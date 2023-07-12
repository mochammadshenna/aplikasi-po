package app

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/lib/pq"
)

func TestOpenConnection(t *testing.T) {
	connStr := "host=localhost port=5432 user=root  password=root dbname=arbrion sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

}
