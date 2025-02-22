package service

import (
	"app/internal/domain"
	"app/models"

	"github.com/google/uuid"
)

type MessageService struct {
	repo domain.MessageRepository
}

func NewMessageService(repo domain.MessageRepository) *MessageService {
	return &MessageService{repo: repo}
}

func (s *MessageService) CreateMessage(content string) (*models.Message, error) {
	message := &models.Message{
		ID:      uuid.New(),
		Message: content,
	}

	if err := message.Validate(); err != nil {
		return nil, err
	}

	if err := s.repo.Create(message); err != nil {
		return nil, err
	}

	return message, nil
}

func (s *MessageService) ListMessages() ([]models.Message, error) {
	return s.repo.List()
}
