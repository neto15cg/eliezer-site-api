package containers

import (
	"database/sql"

	"app/internal/handlers"
	"app/internal/repositories/postgres"
	"app/internal/services"
)

type ChatContainer struct {
	ChntHandler *handlers.ChatHandler
}

func InitializeChatContainer(db *sql.DB) (*ChatContainer, error) {
	messageRepo := postgres.NewMessageRepository(db)
	messageService := services.NewMessageService(messageRepo)

	chatGPTService := services.NewChatGPTService()

	chatHandler := handlers.NewChatHandler(chatGPTService, messageService)

	return &ChatContainer{
		ChntHandler: chatHandler,
	}, nil
}
