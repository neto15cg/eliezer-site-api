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
	"app/internal/repositories/repositories_tests"
	"app/internal/services/services_tests"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sashabaranov/go-openai"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Create a setup function to initialize the common components
func setupChatTest() (*gin.Engine, *services_tests.MockOpenAIService, *services_tests.MockMessageService) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockOpenAIService := new(services_tests.MockOpenAIService)
	mockMessageRepo := new(repositories_tests.MockMessageRepository)
	mockMessageService := services_tests.NewMockMessageService(mockMessageRepo)

	chatController := controllers.NewChatController(mockOpenAIService, mockMessageService)
	router.POST("/chat", chatController.HandleChatMessage)

	return router, mockOpenAIService, mockMessageService
}

func TestHandleChatMessage(t *testing.T) {
	// Test case: New conversation (no conversation ID)
	t.Run("New Conversation", func(t *testing.T) {
		router, mockOpenAIService, mockMessageService := setupChatTest()

		// Expected response from OpenAI
		aiResponse := "This is the AI response"
		userMessage := "Hello, AI"

		// Setup mocks
		mockOpenAIService.On("SendMessage", userMessage, []openai.ChatCompletionMessage{}).
			Return(aiResponse, nil).Once()

		// Create message match function
		mockMessageService.On("CreateMessage", mock.MatchedBy(func(m *entities.Message) bool {
			return m.Message == userMessage && *m.Response == aiResponse && m.ConversationID != nil
		})).Return(nil).Once()

		// Setup response data for GetByConversationID
		testUUID := uuid.New()
		mockMessageService.On("GetByConversationID", mock.AnythingOfType("*uuid.UUID")).
			Run(func(args mock.Arguments) {
				// Capture the conversation ID for reference
				testUUID = *args.Get(0).(*uuid.UUID)
			}).
			Return([]entities.Message{
				{
					ID:             uuid.New(),
					ConversationID: &testUUID,
					Message:        userMessage,
					Response:       &aiResponse,
				},
			}, nil).Once()

		// Create request
		request := entities.OpenAIRequest{
			Message: userMessage,
		}
		jsonRequest, _ := json.Marshal(request)

		req, _ := http.NewRequest(http.MethodPost, "/chat", bytes.NewBuffer(jsonRequest))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// Perform the request
		router.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusOK, w.Code)

		var response []entities.Message
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Len(t, response, 1)
		assert.Equal(t, userMessage, response[0].Message)
		assert.Equal(t, aiResponse, *response[0].Response)
		assert.NotNil(t, response[0].ConversationID)

		mockOpenAIService.AssertExpectations(t)
		mockMessageService.AssertExpectations(t)
	})

	// Test case: Existing conversation (with history)
	t.Run("Existing Conversation", func(t *testing.T) {
		router, mockOpenAIService, mockMessageService := setupChatTest()

		// Setup test data
		conversationID := uuid.New()
		userMessage := "What is the next step?"
		aiResponse := "Here is the next step."
		previousUserMessage := "How do I start?"
		previousAIResponse := "Start by doing this."

		// Setup history to be returned from the DB
		historyMessages := []entities.Message{
			{
				ID:             uuid.New(),
				ConversationID: &conversationID,
				Message:        previousUserMessage,
				Response:       &previousAIResponse,
			},
		}

		// Setup expected chat history
		expectedChatHistory := []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: previousUserMessage,
			},
			{
				Role:    openai.ChatMessageRoleAssistant,
				Content: previousAIResponse,
			},
		}

		// Mock service calls
		mockMessageService.On("GetByConversationID", mock.MatchedBy(func(id *uuid.UUID) bool {
			return id != nil && *id == conversationID
		})).Return(historyMessages, nil).Once()

		mockOpenAIService.On("SendMessage", userMessage, expectedChatHistory).
			Return(aiResponse, nil).Once()

		mockMessageService.On("CreateMessage", mock.MatchedBy(func(m *entities.Message) bool {
			return m.Message == userMessage &&
				*m.Response == aiResponse &&
				m.ConversationID != nil &&
				*m.ConversationID == conversationID
		})).Return(nil).Once()

		// Updated history for final response
		updatedHistory := append([]entities.Message{}, historyMessages...)
		updatedHistory = append(updatedHistory, entities.Message{
			ID:             uuid.New(),
			ConversationID: &conversationID,
			Message:        userMessage,
			Response:       &aiResponse,
		})

		mockMessageService.On("GetByConversationID", mock.MatchedBy(func(id *uuid.UUID) bool {
			return id != nil && *id == conversationID
		})).Return(updatedHistory, nil).Once()

		// Create request
		request := entities.OpenAIRequest{
			Message:        userMessage,
			ConversationID: &conversationID,
		}
		jsonRequest, _ := json.Marshal(request)

		req, _ := http.NewRequest(http.MethodPost, "/chat", bytes.NewBuffer(jsonRequest))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// Perform the request
		router.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusOK, w.Code)

		var response []entities.Message
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Len(t, response, 2) // Should have both the old and new messages

		mockOpenAIService.AssertExpectations(t)
		mockMessageService.AssertExpectations(t)
	})

	// Test case: Error in binding JSON
	t.Run("Invalid JSON", func(t *testing.T) {
		router, _, _ := setupChatTest()

		// Create invalid JSON request
		req, _ := http.NewRequest(http.MethodPost, "/chat", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// Perform the request
		router.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	// Test case: Error retrieving conversation history
	t.Run("Error Getting History", func(t *testing.T) {
		router, _, mockMessageService := setupChatTest()

		conversationID := uuid.New()
		userMessage := "Hello"

		mockMessageService.On("GetByConversationID", mock.MatchedBy(func(id *uuid.UUID) bool {
			return id != nil && *id == conversationID
		})).Return([]entities.Message{}, errors.New("database error")).Once()

		// Create request
		request := entities.OpenAIRequest{
			Message:        userMessage,
			ConversationID: &conversationID,
		}
		jsonRequest, _ := json.Marshal(request)

		req, _ := http.NewRequest(http.MethodPost, "/chat", bytes.NewBuffer(jsonRequest))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// Perform the request
		router.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		mockMessageService.AssertExpectations(t)
	})

	// Test case: Error in OpenAI service
	t.Run("OpenAI Service Error", func(t *testing.T) {
		router, mockOpenAIService, _ := setupChatTest()

		userMessage := "Hello"

		mockOpenAIService.On("SendMessage", userMessage, []openai.ChatCompletionMessage{}).
			Return("", errors.New("OpenAI API error")).Once()

		// Create request
		request := entities.OpenAIRequest{
			Message: userMessage,
		}
		jsonRequest, _ := json.Marshal(request)

		req, _ := http.NewRequest(http.MethodPost, "/chat", bytes.NewBuffer(jsonRequest))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// Perform the request
		router.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		mockOpenAIService.AssertExpectations(t)
	})

	// Test case: Error creating message
	t.Run("Error Creating Message", func(t *testing.T) {
		router, mockOpenAIService, mockMessageService := setupChatTest()

		userMessage := "Hello"
		aiResponse := "Hi there"

		mockOpenAIService.On("SendMessage", userMessage, []openai.ChatCompletionMessage{}).
			Return(aiResponse, nil).Once()

		mockMessageService.On("CreateMessage", mock.AnythingOfType("*entities.Message")).
			Return(errors.New("database error")).Once()

		// Create request
		request := entities.OpenAIRequest{
			Message: userMessage,
		}
		jsonRequest, _ := json.Marshal(request)

		req, _ := http.NewRequest(http.MethodPost, "/chat", bytes.NewBuffer(jsonRequest))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// Perform the request
		router.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "Failed to save message")
		mockOpenAIService.AssertExpectations(t)
		mockMessageService.AssertExpectations(t)
	})

	// Test case: Error getting messages after creation
	t.Run("Error Getting Messages After Creation", func(t *testing.T) {
		router, mockOpenAIService, mockMessageService := setupChatTest()

		userMessage := "Hello"
		aiResponse := "Hi there"

		mockOpenAIService.On("SendMessage", userMessage, []openai.ChatCompletionMessage{}).
			Return(aiResponse, nil).Once()

		mockMessageService.On("CreateMessage", mock.AnythingOfType("*entities.Message")).
			Return(nil).Once()

		mockMessageService.On("GetByConversationID", mock.AnythingOfType("*uuid.UUID")).
			Return([]entities.Message{}, errors.New("database error")).Once()

		// Create request
		request := entities.OpenAIRequest{
			Message: userMessage,
		}
		jsonRequest, _ := json.Marshal(request)

		req, _ := http.NewRequest(http.MethodPost, "/chat", bytes.NewBuffer(jsonRequest))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		// Perform the request
		router.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "Failed to get messages")
		mockOpenAIService.AssertExpectations(t)
		mockMessageService.AssertExpectations(t)
	})
}
