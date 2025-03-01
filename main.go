package main

import (
	"fmt"
	"log"

	"app/config/config"
	"app/config/database"
	"app/internal/controllers"
	"app/internal/repositories"
	"app/internal/routes"
	"app/internal/services"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Run migrations with config
	if err := database.RunMigrations(cfg); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Setup database connection
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close() // Close the connection when the application exits

	// Verify connection is active
	if err := db.Ping(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	log.Println("Database connection established successfully")

	messageRepo := repositories.NewMessageRepositoryPq(db)
	messageService := services.NewMessageService(messageRepo)
	messageController := controllers.NewMessageController(messageService)
	router := routes.SetupRoutes(messageController)

	fmt.Printf("Server starting on port %s...\n", cfg.AppPort)
	if err := router.Run(":" + cfg.AppPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
