package repositories_tests

import (
	"app/internal/entities"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

// MockMessageRepository is a mock implementation of the MessageRepository interface
type MockMessageRepository struct {
	mock.Mock
}

func NewMockMessageRepository() *MockMessageRepository {
	return &MockMessageRepository{}
}

func (m *MockMessageRepository) Create(message *entities.Message) error {
	args := m.Called(message)
	return args.Error(0)
}

func (m *MockMessageRepository) List() ([]entities.Message, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entities.Message), args.Error(1)
}

func (m *MockMessageRepository) GetByID(id uuid.UUID) (*entities.Message, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Message), args.Error(1)
}

func (m *MockMessageRepository) GetByConversationID(conversationID *uuid.UUID) ([]entities.Message, error) {
	args := m.Called(conversationID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entities.Message), args.Error(1)
}
