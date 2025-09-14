package main

import (
	"github.com/AriSu2904/go-auth/internal/config"
	"github.com/AriSu2904/go-auth/internal/database"
	"github.com/AriSu2904/go-auth/internal/handler"
	"github.com/AriSu2904/go-auth/internal/repository"
	"github.com/AriSu2904/go-auth/internal/router"
	"github.com/AriSu2904/go-auth/internal/service"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	log.Println("Initialize service...")

	loadedCfg, err := config.LoadConfig()

	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	database.MigrateSchema(loadedCfg.DBSource)
	db := database.ConnectDB(loadedCfg.DBSource)

	userRepository := repository.NewUserRepository(db, logger)
	authService := service.NewAuthService(userRepository, logger, loadedCfg)
	userService := service.NewUserService(userRepository, logger)
	authHandler := handler.NewAuthHandler(authService, logger)
	userHandler := handler.NewUserHandler(userService, logger)

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
