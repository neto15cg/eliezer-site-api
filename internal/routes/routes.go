package routes

import (
	"app/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(messageController *controllers.MessageController) *gin.Engine {
	router := gin.Default()

	router.GET("/messages", messageController.GetMessages)
	router.POST("/messages", messageController.CreateMessage)
	router.GET("/messages/conversation/:conversation_id", messageController.GetByConversationID)

	return router
}
