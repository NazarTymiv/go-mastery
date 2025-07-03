package db

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func Connect(dsn string) *sqlx.DB {
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Could not open DB: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	log.Println("Connected to MySQL")
	return db
}
