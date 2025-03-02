package services_tests

import (
	"os"
	"testing"

	"app/internal/services"

	openai "github.com/sashabaranov/go-openai"
	"github.com/stretchr/testify/assert"
)

// TestNewChatOpenaiService tests the creation of the OpenAI service
func TestNewChatOpenaiService(t *testing.T) {
	// Store original API key
	originalKey := os.Getenv("OPENAI_API_KEY")
	// Set a test key
	os.Setenv("OPENAI_API_KEY", "test-api-key")
	defer os.Setenv("OPENAI_API_KEY", originalKey)

	service := services.NewChatOpenaiService()
	assert.NotNil(t, service)
}

// TestSendMessage tests the SendMessage function
// Note: This is an integration test that requires a real OpenAI API key
// and will make actual API calls. It's recommended to skip this test
// in automated test environments or if no API key is available.
func TestSendMessage(t *testing.T) {
	// Check if we have a real API key for integration testing
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" || apiKey == "test-api-key" {
		t.Skip("Skipping integration test: No valid OpenAI API key available")
	}

	// Store original prompt and set a test prompt
	originalPrompt := os.Getenv("CHATBOT_PROMPT")
	os.Setenv("CHATBOT_PROMPT", "You are a test assistant. Give very short responses for testing purposes only.")
	defer os.Setenv("CHATBOT_PROMPT", originalPrompt)

	service := services.NewChatOpenaiService()

	t.Run("send message with empty history", func(t *testing.T) {
		message := "Say hello for a test"
		history := []openai.ChatCompletionMessage{}

		response, err := service.SendMessage(message, history)

		// We should get a response without error
		assert.NoError(t, err)
		assert.NotEmpty(t, response)
	})

	t.Run("send message with history", func(t *testing.T) {
		message := "What was my previous question?"
		history := []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "What is the capital of France?",
			},
			{
				Role:    openai.ChatMessageRoleAssistant,
				Content: "The capital of France is Paris.",
			},
		}

		response, err := service.SendMessage(message, history)

		// We should get a response without error
		assert.NoError(t, err)
		assert.NotEmpty(t, response)
	})
}
