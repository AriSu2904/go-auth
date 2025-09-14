package main

import (
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

	db := database.ConnectDB(loadedCfg.DBSource)

	defer db.Close().Error()

	log.Println("Service started...")
}
