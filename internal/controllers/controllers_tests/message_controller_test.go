package controllers_tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"app/internal/controllers"
	"app/internal/entities"
	"app/internal/services/services_tests"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.Default()
}

func TestGetMessages(t *testing.T) {
	// Setup
	mockService := new(services_tests.MockMessageService)
	controller := controllers.NewMessageController(mockService)
	router := setupRouter()
	router.GET("/messages", controller.GetMessages)

	// Test case: Successful retrieval
	expectedMessages := []entities.Message{
		{ID: uuid.New(), Message: "Test message 1"},
		{ID: uuid.New(), Message: "Test message 2"},
	}
	mockService.On("GetMessages").Return(expectedMessages, nil).Once()

	req := httptest.NewRequest("GET", "/messages", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	var responseMessages []entities.Message
	err := json.Unmarshal(w.Body.Bytes(), &responseMessages)
	assert.NoError(t, err)
	assert.Equal(t, expectedMessages, responseMessages)
	mockService.AssertExpectations(t)

	// Test case: Error retrieving messages
	mockService.On("GetMessages").Return([]entities.Message{}, errors.New("service error")).Once()

	req = httptest.NewRequest("GET", "/messages", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	var errorResponse map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &errorResponse)
	assert.NoError(t, err)
	assert.Contains(t, errorResponse, "error")
	mockService.AssertExpectations(t)
}

func TestCreateMessage(t *testing.T) {
	// Setup
	mockService := new(services_tests.MockMessageService)
	controller := controllers.NewMessageController(mockService)
	router := setupRouter()
	router.POST("/messages", controller.CreateMessage)

	// Test case: Successful message creation
	message := entities.Message{
		Message: "Test message",
	}
	messageBytes, _ := json.Marshal(message)

	mockService.On("CreateMessage", mock.MatchedBy(func(m *entities.Message) bool {
		return m.Message == message.Message
	})).Return(nil).Once()

	req := httptest.NewRequest("POST", "/messages", bytes.NewBuffer(messageBytes))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, w.Code)
	var responseMessage entities.Message
	err := json.Unmarshal(w.Body.Bytes(), &responseMessage)
	assert.NoError(t, err)
	assert.Equal(t, message.Message, responseMessage.Message)
	mockService.AssertExpectations(t)

	// Test case: Invalid JSON
	req = httptest.NewRequest("POST", "/messages", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockService.AssertExpectations(t)

	// Test case: Service error
	mockService.On("CreateMessage", mock.MatchedBy(func(m *entities.Message) bool {
		return m.Message == message.Message
	})).Return(errors.New("service error")).Once()

	req = httptest.NewRequest("POST", "/messages", bytes.NewBuffer(messageBytes))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockService.AssertExpectations(t)
}

func TestGetByConversationID(t *testing.T) {
	// Setup
	mockService := new(services_tests.MockMessageService)
	controller := controllers.NewMessageController(mockService)
	router := setupRouter()
	router.GET("/conversations/:conversation_id/messages", controller.GetByConversationID)

	// Test case: Successful retrieval
	conversationID := uuid.New()
	expectedMessages := []entities.Message{
		{ID: uuid.New(), Message: "Test message 1", ConversationID: &conversationID},
		{ID: uuid.New(), Message: "Test message 2", ConversationID: &conversationID},
	}

	mockService.On("GetByConversationID", mock.MatchedBy(func(id *uuid.UUID) bool {
		return id != nil && *id == conversationID
	})).Return(expectedMessages, nil).Once()

	req := httptest.NewRequest("GET", "/conversations/"+conversationID.String()+"/messages", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	var responseMessages []entities.Message
	err := json.Unmarshal(w.Body.Bytes(), &responseMessages)
	assert.NoError(t, err)
	assert.Equal(t, expectedMessages, responseMessages)
	mockService.AssertExpectations(t)

	// Test case: Invalid UUID format
	req = httptest.NewRequest("GET", "/conversations/invalid-uuid/messages", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockService.AssertExpectations(t)

	// Test case: Service error
	mockService.On("GetByConversationID", mock.MatchedBy(func(id *uuid.UUID) bool {
		return id != nil && *id == conversationID
	})).Return([]entities.Message{}, errors.New("service error")).Once()

	req = httptest.NewRequest("GET", "/conversations/"+conversationID.String()+"/messages", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockService.AssertExpectations(t)
}
