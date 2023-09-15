package app

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/lib/pq"
	config "github.com/mochammadshenna/aplikasi-po/configs"
)

func TestOpenConnection(t *testing.T) {
	var dbConfig = config.Get().Database

	connStr := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.DbName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

}
