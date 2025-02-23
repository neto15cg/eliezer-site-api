package handlers

import (
	"net/http"

	"app/internal/interfaces"
	"app/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ChatHandler struct {
	chatGPTService interfaces.ChatGPTService
	messageService interfaces.MessageService
}

func NewChatHandler(chatService interfaces.ChatGPTService, msgService interfaces.MessageService) *ChatHandler {
	return &ChatHandler{
		chatGPTService: chatService,
		messageService: msgService,
	}
}

func (h *ChatHandler) SendConversationMessage(c *gin.Context) {
	var req models.ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate conversation ID if not provided
	if req.ConversationID == uuid.Nil {
		req.ConversationID = uuid.New()
	}

	// Send message to ChatGPT
	response, err := h.chatGPTService.SendMessage(req.Message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Save both question and response
	message, err := h.messageService.CreateMessage(req.Message, response, &req.ConversationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, message)
}
