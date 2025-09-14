package database

import (
	"database/sql"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

func MigrateSchema(dbSource string) {
	migrationPath := "file://internal/database/migration"

	migrated, err := migrate.New(migrationPath, dbSource)

	if err != nil {
		log.Fatal("Could not start migration: ", err)
	}

	err = migrated.Up()

	if err != nil {
		log.Println("Could not run migration: ", err)
	}

	log.Println("Database migrated successfully")
}
