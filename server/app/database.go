package app

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	config "github.com/mochammadshenna/aplikasi-po/configs"
	"github.com/mochammadshenna/aplikasi-po/util/helper"
	"github.com/mochammadshenna/aplikasi-po/util/logger"
)

func NewDb() *sql.DB {
	return newDb(config.Get().Database.DbName)
}

func newDb(dbName string) *sql.DB {

	var dbConfig = config.Get().Database

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s",
		dbConfig.Host, 
		dbConfig.Port, 
		dbConfig.Username, 
		dbConfig.Password, 
		dbName,
	)
	
	db, err := sql.Open("postgres", psqlInfo)
	helper.PanicError(err)
	if err = db.Ping(); err != nil {
		logger.Fatal(context.TODO(), err)
		if err = db.Close(); err != nil {
			logger.Fatal(context.TODO(), err)
		}
	}

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
