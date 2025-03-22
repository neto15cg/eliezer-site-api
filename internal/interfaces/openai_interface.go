package interfaces

import (
	"app/internal/services"

	"github.com/sashabaranov/go-openai"
)

type OpenAIServiceInterface interface {
	SendMessage(message string, history []openai.ChatCompletionMessage) (string, error)
}

var _ OpenAIServiceInterface = (*services.OpenAIService)(nil)
