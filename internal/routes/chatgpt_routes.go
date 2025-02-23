package routes

import (
	"app/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupChatGPTRoutes(router *gin.Engine, chatGPTHandler *handlers.ChatGPTHandler) {
	chat := router.Group("/chat")
	{
		chat.POST("/send", func(c *gin.Context) {
			chatGPTHandler.SendMessage(c.Writer, c.Request)
		})
	}
}
