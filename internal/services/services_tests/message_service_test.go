package services_tests

import (
	"errors"
	"testing"
	"time"

	"app/internal/entities"
	"app/internal/repositories/repositories_tests"
	"app/internal/services"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewMessageService(t *testing.T) {
	mockRepo := new(repositories_tests.MockMessageRepository)
	service := services.NewMessageService(mockRepo)

	assert.NotNil(t, service)
}

func TestGetMessages(t *testing.T) {
	mockRepo := new(repositories_tests.MockMessageRepository)
	service := services.NewMessageService(mockRepo)

	expectedMessages := []entities.Message{
		{ID: uuid.New(), Message: "Test message 1"},
		{ID: uuid.New(), Message: "Test message 2"},
	}

	// Test successful retrieval
	mockRepo.On("List").Return(expectedMessages, nil).Once()

	messages, err := service.GetMessages()

	assert.NoError(t, err)
	assert.Equal(t, expectedMessages, messages)
	mockRepo.AssertExpectations(t)

	// Test error case
	expectedErr := errors.New("database error")
	mockRepo.On("List").Return([]entities.Message{}, expectedErr).Once()

	messages, err = service.GetMessages()

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Empty(t, messages)
	mockRepo.AssertExpectations(t)
}

func TestGetMessageById(t *testing.T) {
	mockRepo := new(repositories_tests.MockMessageRepository)
	service := services.NewMessageService(mockRepo)

	messageID := uuid.New()
	expectedMessage := &entities.Message{ID: messageID, Message: "Test message"}

	// Test successful retrieval
	mockRepo.On("GetByID", messageID).Return(expectedMessage, nil).Once()

	message, err := service.GetMessageById(messageID)

	assert.NoError(t, err)
	assert.Equal(t, expectedMessage, message)
	mockRepo.AssertExpectations(t)

	// Test message not found
	mockRepo.On("GetByID", messageID).Return(nil, errors.New("message not found")).Once()

	message, err = service.GetMessageById(messageID)

	assert.Error(t, err)
	assert.Nil(t, message)
	mockRepo.AssertExpectations(t)
}

func TestGetByConversationID(t *testing.T) {
	mockRepo := new(repositories_tests.MockMessageRepository)
	service := services.NewMessageService(mockRepo)

	conversationID := uuid.New()
	expectedMessages := []entities.Message{
		{ID: uuid.New(), Message: "Test message 1", ConversationID: &conversationID},
		{ID: uuid.New(), Message: "Test message 2", ConversationID: &conversationID},
	}

	// Test successful retrieval
	mockRepo.On("GetByConversationID", &conversationID).Return(expectedMessages, nil).Once()

	messages, err := service.GetByConversationID(&conversationID)

	assert.NoError(t, err)
	assert.Equal(t, expectedMessages, messages)
	mockRepo.AssertExpectations(t)

	// Test error case
	expectedErr := errors.New("database error")
	mockRepo.On("GetByConversationID", &conversationID).Return([]entities.Message{}, expectedErr).Once()

	messages, err = service.GetByConversationID(&conversationID)

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Empty(t, messages)
	mockRepo.AssertExpectations(t)
}

func TestCreateMessage(t *testing.T) {
	mockRepo := new(repositories_tests.MockMessageRepository)
	service := services.NewMessageService(mockRepo)

	// Test creating a message with nil ID (should be generated)
	message := &entities.Message{
		Message: "Test message",
	}

	mockRepo.On("Create", mock.AnythingOfType("*entities.Message")).Return(nil).Once()

	err := service.CreateMessage(message)

	assert.NoError(t, err)
	assert.NotEqual(t, uuid.Nil, message.ID)
	assert.False(t, message.CreatedAt.IsZero())
	assert.False(t, message.UpdatedAt.IsZero())
	mockRepo.AssertExpectations(t)

	// Test with existing ID and timestamp
	messageID := uuid.New()
	createdTime := time.Now().Add(-1 * time.Hour)
	message = &entities.Message{
		ID:        messageID,
		Message:   "Test message with ID",
		CreatedAt: createdTime,
		UpdatedAt: createdTime,
	}

	mockRepo.On("Create", mock.AnythingOfType("*entities.Message")).Return(nil).Once()

	err = service.CreateMessage(message)

	assert.NoError(t, err)
	assert.Equal(t, messageID, message.ID)
	assert.Equal(t, createdTime, message.CreatedAt)
	mockRepo.AssertExpectations(t)

	// Test error case
	expectedErr := errors.New("database error")
	mockRepo.On("Create", mock.AnythingOfType("*entities.Message")).Return(expectedErr).Once()

	err = service.CreateMessage(message)

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	mockRepo.AssertExpectations(t)
}
