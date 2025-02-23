package services

import (
	"app/internal/domain"
	"app/models"

	"github.com/google/uuid"
)

type messageService struct {
	repo domain.MessageRepository
}

func NewMessageService(repo domain.MessageRepository) domain.MessageService {
	return &messageService{repo: repo}
}

func (s *messageService) CreateMessage(content string) (*models.Message, error) {
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

func (s *messageService) ListMessages() ([]models.Message, error) {
	return s.repo.List()
}

func (s *messageService) GetMessage(id uuid.UUID) (*models.Message, error) {
	return s.repo.GetByID(id)
}
