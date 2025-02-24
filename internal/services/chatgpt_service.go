package services

import (
	"context"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func (s *chatGPTService) SendMessage(message string, history []openai.ChatCompletionMessage) (string, error) {
	prompt := os.Getenv("CHATBOT_PROMPT")

	messages := make([]openai.ChatCompletionMessage, 0)
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: prompt,
	})
	messages = append(messages, history...)
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: message,
	})

	resp, err := s.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT4oMini,
			Messages: messages,
		},
	)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
