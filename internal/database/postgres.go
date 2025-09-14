package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func ConnectDB(dbSource string) *sql.DB {
	db, err := sql.Open("postgres", dbSource)

	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	err = db.Ping()

	if err != nil {
		log.Fatal("Error pinging the database: ", err)
	}

	log.Println("Successfully connected to the database")
	return db
}
