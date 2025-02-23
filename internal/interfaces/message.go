package interfaces

import (
	"app/models"

	"github.com/google/uuid"
)

type MessageRepository interface {
	Create(message *models.Message) error
	List() ([]models.Message, error)
	GetByID(id uuid.UUID) (*models.Message, error)
	GetByConversationID(conversationID uuid.UUID) ([]models.Message, error)
}

type MessageService interface {
	CreateMessage(content string, conversationID *uuid.UUID) (*models.Message, error)
	ListMessages() ([]models.Message, error)
	GetMessage(id uuid.UUID) (*models.Message, error)
	GetMessagesByConversationID(conversationID uuid.UUID) ([]models.Message, error)
}
