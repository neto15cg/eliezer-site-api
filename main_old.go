package main

// import (
// 	"fmt"
// 	"log"
// 	"time"

// 	"app/internal_old/containers"
// 	"app/internal_old/routes"
// 	"app/pkg/config"
// 	"app/pkg/database"

// 	"github.com/gin-gonic/gin"
// 	_ "github.com/lib/pq"
// )

// func main_old() {
// 	// Load configuration
// 	cfg, err := config.LoadConfig()
// 	if err != nil {
// 		log.Fatalf("Failed to load config: %v", err)
// 	}

// 	// Run migrations with config
// 	if err := database.RunMigrations(cfg); err != nil {
// 		log.Fatalf("Failed to run migrations: %v", err)
// 	}

// 	// Setup database connection
// 	db, err := database.Connect(cfg)
// 	if err != nil {
// 		log.Fatalf("Failed to connect to database: %v", err)
// 	}
// 	defer db.Close()

// 	// Initialize containers
// 	messageContainer, err := containers.Initialize(db)
// 	if err != nil {
// 		log.Fatalf("Failed to initialize container: %v", err)
// 	}

// 	// Initialize Gin
// 	r := gin.Default()

// 	// Setup routes
// 	routes.SetupConversationRoutes(r, messageContainer.Chat.ChntHandler, messageContainer.Messages.MessageHandler)

// 	// Add health check endpoint
// 	r.GET("/health", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"status": "ok",
// 			"time":   time.Now(),
// 		})
// 	})

// 	fmt.Printf("Server starting on port %s...\n", cfg.AppPort)
// 	if err := r.Run(":" + cfg.AppPort); err != nil {
// 		panic(err)
// 	}
// }
