package services

import (
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

func (s *MessageService) GetByConversationID(id uuid.UUID) ([]entities.Message, error) {
	return s.repo.GetByConversationID(id)
}

func (s *MessageService) CreateMessage(message *entities.Message) error {
	return s.repo.Create(message)
}
