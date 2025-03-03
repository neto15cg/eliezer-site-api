package interfaces

import (
	"app/internal/entities"
	"app/internal/services"

	"github.com/google/uuid"
)

type MessageServiceInterface interface {
	GetMessages() ([]entities.Message, error)
	GetMessageById(id uuid.UUID) (*entities.Message, error)
	GetByConversationID(id *uuid.UUID) ([]entities.Message, error)
	CreateMessage(message *entities.Message) error
}

var _ MessageServiceInterface = (*services.MessageService)(nil)
