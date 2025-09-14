package main

import (
	"database/sql"
	"github.com/AriSu2904/go-auth/internal/config"
	"github.com/AriSu2904/go-auth/internal/database"
	"github.com/AriSu2904/go-auth/internal/repository"
	"github.com/AriSu2904/go-auth/internal/service"
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
	userRepository := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepository)
	_ = authService

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("Error closing the database: ", err)
		}
	}(db)

	log.Println("Service started...")
}
