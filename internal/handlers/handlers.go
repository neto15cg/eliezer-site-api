package handlers

import "app/internal/interfaces"

type MessageHandler struct {
	service interfaces.MessageService
}

type ChatGPTHandler struct {
	chatGPTService interfaces.ChatGPTService
}

func NewMessageHandler(service interfaces.MessageService) *MessageHandler {
	return &MessageHandler{service: service}
}

func NewChatGPTHandler(service interfaces.ChatGPTService) *ChatGPTHandler {
	return &ChatGPTHandler{
		chatGPTService: service,
	}
}
