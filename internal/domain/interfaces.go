package domain

import (
	"app/models"

	"github.com/google/uuid"
)

// Repository interfaces
type MessageReader interface {
	List() ([]models.Message, error)
	GetByID(id uuid.UUID) (*models.Message, error)
}

type MessageWriter interface {
	Create(message *models.Message) error
}

// Service interfaces
type MessageService interface {
	CreateMessage(content string) (*models.Message, error)
	ListMessages() ([]models.Message, error)
	GetMessage(id uuid.UUID) (*models.Message, error)
}
