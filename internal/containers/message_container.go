package containers

import (
	"database/sql"

	"app/internal/handlers"
	"app/internal/repositories/postgres"
	"app/internal/services"
)

type MessageContainer struct {
	MessageHandler *handlers.MessageHandler
}

func InitializeMessageContainer(db *sql.DB) (*MessageContainer, error) {
	// Initialize repository
	messageRepo := postgres.NewMessageRepository(db)

	// Initialize service
	messageService := services.NewMessageService(messageRepo)

	// Initialize handler
	messageHandler := handlers.NewMessageHandler(messageService)

	return &MessageContainer{
		MessageHandler: messageHandler,
	}, nil
}
