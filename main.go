package main

import (
	"log"

	"github.com/adityasharma3/go-social/internal/config"
	"github.com/adityasharma3/go-social/internal/handlers"
	"github.com/adityasharma3/go-social/internal/server"
	"github.com/adityasharma3/go-social/internal/store"
)

func main() {
	// Initialize configuration
	cfg := config.NewConfig()

	// Initialize store
	store := store.NewStore()

	// Initialize server
	srv := server.NewServer(cfg.ServerAddr)

	// Initialize handlers
	healthHandler := handlers.NewHealthHandler()
	userHandler := handlers.NewUserHandler(store)

	// Setup routes
	srv.Router().GET("/health", healthHandler.Handle)
	srv.Router().GET("/users", userHandler.GetUsers)
	srv.Router().POST("/users", userHandler.CreateUser)

	// Start server
	if err := srv.Start(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
