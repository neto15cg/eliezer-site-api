package containers

import (
	"database/sql"

	"app/internal/handlers"
	"app/internal/repository/postgres"
	"app/internal/services"
)

type Container struct {
	MessageHandler *handlers.MessageHandler
}

func InitializeMessageContainer(db *sql.DB) (*Container, error) {
	// Initialize repository
	messageRepo := postgres.NewMessageRepository(db)

	// Initialize service
	messageService := services.NewMessageService(messageRepo)

	// Initialize handler
	messageHandler := handlers.NewMessageHandler(messageService)

	return &Container{
		MessageHandler: messageHandler,
	}, nil
}
