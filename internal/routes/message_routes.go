package routes

import (
	"app/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupMessageRoutes(router *gin.Engine, messageHandler *handler.MessageHandler) {
	messages := router.Group("/messages")
	{
		messages.POST("", messageHandler.Create)
		messages.GET("", messageHandler.List)
		messages.GET("/:id", messageHandler.GetByID)
	}
}
