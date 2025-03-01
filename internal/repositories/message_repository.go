package repositories

import (
	"app/internal/entities"

	"github.com/google/uuid"
)

type MessageRepository interface {
	Create(message *entities.Message) error
	List() ([]entities.Message, error)
	GetByID(id uuid.UUID) (*entities.Message, error)
	GetByConversationID(conversationID *uuid.UUID) ([]entities.Message, error)
}
