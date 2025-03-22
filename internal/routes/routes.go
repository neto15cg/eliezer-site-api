package routes

import (
	"time"

	"app/internal/controllers"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(messageController *controllers.MessageController, openai *controllers.ChatController) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://eliezer-marques.click", "https://eliezer-site.s3.us-east-1.amazonaws.com"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000" || origin == "https://eliezer-marques.click" || origin == "https://eliezer-site.s3.us-east-1.amazonaws.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	api := router.Group("/v1")
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
