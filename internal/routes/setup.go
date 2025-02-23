package routes

import (
	"app/internal/handlers"
	"app/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Initialize services
	chatGPTService := services.NewChatGPTService()

	// Initialize handlers
	chatGPTHandler := handlers.NewChatGPTHandler(chatGPTService)

	// Setup routes
	SetupChatGPTRoutes(router, chatGPTHandler)
	// ...existing message routes setup...
}
