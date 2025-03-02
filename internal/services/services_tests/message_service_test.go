package services_tests

import (
	"errors"
	"testing"
	"time"

	"app/internal/entities"
	"app/internal/repositories/repositories_tests"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewMessageService(t *testing.T) {
	mockRepo := new(repositories_tests.MockMessageRepository)
	service := NewMockMessageService(mockRepo)

	assert.NotNil(t, service)
}

func TestGetMessages(t *testing.T) {
	mockRepo := new(repositories_tests.MockMessageRepository)
	service := NewMockMessageService(mockRepo)

	expectedMessages := []entities.Message{
		{ID: uuid.New(), Message: "Test message 1"},
		{ID: uuid.New(), Message: "Test message 2"},
	}

	// Test successful retrieval
	service.On("GetMessages").Return(expectedMessages, nil).Once()

	messages, err := service.GetMessages()

	assert.NoError(t, err)
	assert.Equal(t, expectedMessages, messages)
	service.AssertExpectations(t)

	// Test error case
	expectedErr := errors.New("database error")
	service.On("GetMessages").Return([]entities.Message{}, expectedErr).Once()

	messages, err = service.GetMessages()

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Empty(t, messages)
	service.AssertExpectations(t)
}

func TestGetMessageById(t *testing.T) {
	mockRepo := new(repositories_tests.MockMessageRepository)
	service := NewMockMessageService(mockRepo)

	messageID := uuid.New()
	expectedMessage := &entities.Message{ID: messageID, Message: "Test message"}

	// Test successful retrieval
	service.On("GetMessageById", messageID).Return(expectedMessage, nil).Once()

	message, err := service.GetMessageById(messageID)

	assert.NoError(t, err)
	assert.Equal(t, expectedMessage, message)
	service.AssertExpectations(t)

	// Test message not found
	service.On("GetMessageById", messageID).Return(nil, errors.New("message not found")).Once()

	message, err = service.GetMessageById(messageID)

	assert.Error(t, err)
	assert.Nil(t, message)
	service.AssertExpectations(t)
}

func TestGetByConversationID(t *testing.T) {
	mockRepo := new(repositories_tests.MockMessageRepository)
	service := NewMockMessageService(mockRepo)

	conversationID := uuid.New()
	expectedMessages := []entities.Message{
		{ID: uuid.New(), Message: "Test message 1", ConversationID: &conversationID},
		{ID: uuid.New(), Message: "Test message 2", ConversationID: &conversationID},
	}

	// Test successful retrieval
	service.On("GetByConversationID", &conversationID).Return(expectedMessages, nil).Once()

	messages, err := service.GetByConversationID(&conversationID)

	assert.NoError(t, err)
	assert.Equal(t, expectedMessages, messages)
	service.AssertExpectations(t)

	// Test error case
	expectedErr := errors.New("database error")
	service.On("GetByConversationID", &conversationID).Return([]entities.Message{}, expectedErr).Once()

	messages, err = service.GetByConversationID(&conversationID)

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	assert.Empty(t, messages)
	service.AssertExpectations(t)
}

func TestCreateMessage(t *testing.T) {
	mockRepo := new(repositories_tests.MockMessageRepository)
	service := NewMockMessageService(mockRepo)

	// Test creating a message with nil ID (should be generated)
	message := &entities.Message{
		Message: "Test message",
	}

	service.On("CreateMessage", mock.AnythingOfType("*entities.Message")).Return(nil).Once()

	err := service.CreateMessage(message)

	assert.NoError(t, err)
	service.AssertExpectations(t)

	// Test with existing ID and timestamp
	messageID := uuid.New()
	createdTime := time.Now().Add(-1 * time.Hour)
	message = &entities.Message{
		ID:        messageID,
		Message:   "Test message with ID",
		CreatedAt: createdTime,
		UpdatedAt: createdTime,
	}

	service.On("CreateMessage", mock.AnythingOfType("*entities.Message")).Return(nil).Once()

	err = service.CreateMessage(message)

	assert.NoError(t, err)
	service.AssertExpectations(t)

	// Test error case
	expectedErr := errors.New("database error")
	service.On("CreateMessage", mock.AnythingOfType("*entities.Message")).Return(expectedErr).Once()

	err = service.CreateMessage(message)

	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
	service.AssertExpectations(t)
}
