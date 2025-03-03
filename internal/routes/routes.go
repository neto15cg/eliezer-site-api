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

		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "ok",
			})
		})
	}

	// Return router with the API group
	return router
}
