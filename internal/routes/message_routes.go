package routes

import (
	"app/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupMessageRoutes(router *gin.Engine, messageHandler *handlers.MessageHandler) {
	messages := router.Group("/messages")
	{
		messages.GET("/conversation/:conversation_id", messageHandler.GetByConversationId)
	}
}
