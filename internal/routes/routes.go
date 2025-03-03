package routes

import (
	"app/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(messageController *controllers.MessageController, openai *controllers.ChatController) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/messages/conversation/:conversation_id", messageController.GetByConversationID)

		api.POST("/chat", openai.HandleChatMessage)
	}

	// Return router with the API group
	return router
}
