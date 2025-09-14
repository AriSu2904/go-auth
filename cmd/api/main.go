package main

import (
	"database/sql"
	"github.com/AriSu2904/go-auth/internal/config"
	"github.com/AriSu2904/go-auth/internal/database"
	"log"
)

func main() {
	log.Println("Initialize service...")

	loadedCfg, err := config.LoadConfig()

	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	database.MigrateSchema(loadedCfg.DBSource)

	db := database.ConnectDB(loadedCfg.DBSource)

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("Error closing the database: ", err)
		}
	}(db)

	log.Println("Service started...")
}
