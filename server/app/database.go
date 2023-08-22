package app

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/mochammadshenna/aplikasi-po/util/helper"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "root"
	dbname   = "arbrion"
)

func NewDB() *sql.DB {

	connStr := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s ", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	helper.PanicError(err)

	/*
	* handle db close
	**/
	// defer db.Close()
	// err = db.Ping()
	// helper.PanicError(err)

	fmt.Println("Established a successful connection!")

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
