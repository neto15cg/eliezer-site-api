package containers

import (
	"app/internal_old/handlers"
	"app/internal_old/services"
)

type ChatGPTContainer struct {
	ChatGPTHandler *handlers.ChatGPTHandler
}

func InitializeChatGPTContainer() (*ChatGPTContainer, error) {
	// Initialize service
	chatGPTService := services.NewChatGPTService()

	// Initialize handler
	chatGPTHandler := handlers.NewChatGPTHandler(chatGPTService)

	return &ChatGPTContainer{
		ChatGPTHandler: chatGPTHandler,
	}, nil
}
