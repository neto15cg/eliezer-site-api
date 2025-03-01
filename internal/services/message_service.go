package services

import (
	"time"

	"app/internal/entities"
	"app/internal/repositories"

	"github.com/google/uuid"
)

type MessageService struct {
	repo repositories.MessageRepository
}

func NewMessageService(repo repositories.MessageRepository) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) GetMessages() ([]entities.Message, error) {
	return s.repo.List()
}

func (s *MessageService) GetMessageById(id uuid.UUID) (*entities.Message, error) {
	return s.repo.GetByID(id)
}

func (s *MessageService) GetByConversationID(id *uuid.UUID) ([]entities.Message, error) {
	return s.repo.GetByConversationID(id)
}

func (s *MessageService) CreateMessage(message *entities.Message) error {
	// Ensure the message has an ID before saving
	message.EnsureID()

	// Set timestamps if they're empty
	if message.CreatedAt.IsZero() {
		now := time.Now()
		message.CreatedAt = now
		message.UpdatedAt = now
	}

	return s.repo.Create(message)
}
