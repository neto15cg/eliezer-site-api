package services

import (
	"context"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func (s *chatGPTService) SendMessage(message string) (string, error) {
	prompt := os.Getenv("CHATBOT_PROMPT")
	resp, err := s.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4oMini,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: prompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
