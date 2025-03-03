package services_tests

import (
	"app/internal/entities"
	"app/internal/interfaces"
	"app/internal/repositories"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

// MockMessageService is a mock implementation of the MessageService interface
type MockMessageService struct {
	mock.Mock
	repo repositories.MessageRepository
}

// Ensure MockMessageService implements MessageServiceInterface
var _ interfaces.MessageServiceInterface = (*MockMessageService)(nil)

func NewMockMessageService(repo repositories.MessageRepository) *MockMessageService {
	return &MockMessageService{repo: repo}
}

func (m *MockMessageService) GetMessages() ([]entities.Message, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entities.Message), args.Error(1)
}

func (m *MockMessageService) GetMessageById(id uuid.UUID) (*entities.Message, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Message), args.Error(1)
}

func (m *MockMessageService) GetByConversationID(conversationID *uuid.UUID) ([]entities.Message, error) {
	args := m.Called(conversationID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entities.Message), args.Error(1)
}

func (m *MockMessageService) CreateMessage(message *entities.Message) error {
	args := m.Called(message)
	return args.Error(0)
}
