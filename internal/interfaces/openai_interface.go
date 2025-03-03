package interfaces

import (
	"app/internal/services"

	"github.com/sashabaranov/go-openai"
)

// OpenAIServiceInterface defines the contract for OpenAI services
type OpenAIServiceInterface interface {
	SendMessage(message string, history []openai.ChatCompletionMessage) (string, error)
}

// Ensure OpenAIService implements OpenAIServiceInterface
var _ OpenAIServiceInterface = (*services.OpenAIService)(nil)
