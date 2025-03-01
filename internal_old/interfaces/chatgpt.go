package interfaces

import (
	openai "github.com/sashabaranov/go-openai"
)

type ChatGPTService interface {
	SendMessage(message string, history []openai.ChatCompletionMessage) (string, error)
}
