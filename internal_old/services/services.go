package services

import (
	"os"

	"app/internal_old/interfaces"

	openai "github.com/sashabaranov/go-openai"
)

type messageService struct {
	repo interfaces.MessageRepository
}

func NewMessageService(repo interfaces.MessageRepository) interfaces.MessageService {
	return &messageService{repo: repo}
}

type chatGPTService struct {
	client *openai.Client
}

func NewChatGPTService() *chatGPTService {
	apiKey := os.Getenv("OPENAI_API_KEY")
	client := openai.NewClient(apiKey)

	return &chatGPTService{
		client: client,
	}
}
