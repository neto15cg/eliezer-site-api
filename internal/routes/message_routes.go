package routes

import (
	"app/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupMessageRoutes(router *gin.Engine, messageHandler *handlers.MessageHandler) {
	messages := router.Group("/messages")
	{
		messages.POST("", messageHandler.Create)
		messages.GET("", messageHandler.List)
		messages.GET("/:id", messageHandler.GetByID)
		messages.GET("/conversation/:conversation_id", messageHandler.GetByConversationId)
	}
}
