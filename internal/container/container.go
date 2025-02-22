package container

import (
	"database/sql"

	"app/internal/handler"
	"app/internal/repository/postgres"
	"app/internal/service"
)

type Container struct {
	MessageHandler *handler.MessageHandler
}

func InitializeMessageContainer(db *sql.DB) (*Container, error) {
	// Initialize repository
	messageRepo := postgres.NewMessageRepository(db)

	// Initialize service
	messageService := service.NewMessageService(messageRepo)

	// Initialize handler
	messageHandler := handler.NewMessageHandler(messageService)

	return &Container{
		MessageHandler: messageHandler,
	}, nil
}
