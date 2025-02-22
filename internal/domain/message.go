package domain

import (
	"app/models"

	"github.com/google/uuid"
)

type MessageRepository interface {
	Create(message *models.Message) error
	List() ([]models.Message, error)
	GetByID(id uuid.UUID) (*models.Message, error)
}
