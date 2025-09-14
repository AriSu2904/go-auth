package main

import (
	"github.com/AriSu2904/go-auth/internal/config"
	"github.com/AriSu2904/go-auth/internal/database"
	"github.com/AriSu2904/go-auth/internal/handler"
	"github.com/AriSu2904/go-auth/internal/repository"
	"github.com/AriSu2904/go-auth/internal/router"
	"github.com/AriSu2904/go-auth/internal/service"
	"log"
	"net/http"
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
	authService := service.NewAuthService(userRepository, loadedCfg)
	userService := service.NewUserService(userRepository)
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)

	chiRouter := router.NewRouter(authHandler, userHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: chiRouter,
	}

	log.Println("Starting service on port 8080...")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
