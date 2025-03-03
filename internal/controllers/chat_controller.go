package controllers

import (
	"net/http"

	"app/internal/entities"
	"app/internal/interfaces"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	openai "github.com/sashabaranov/go-openai"
)

type ChatController struct {
	openaiService  interfaces.OpenAIServiceInterface
	messageService interfaces.MessageServiceInterface
}

// Updated constructor to use interfaces
func NewChatController(openaiService interfaces.OpenAIServiceInterface, messageService interfaces.MessageServiceInterface) *ChatController {
	return &ChatController{
		openaiService:  openaiService,
		messageService: messageService,
	}
}

func (c *ChatController) HandleChatMessage(ctx *gin.Context) {
	var request entities.OpenAIRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	messagesHistory := make([]openai.ChatCompletionMessage, 0)

	conversationID := request.ConversationID
	if conversationID != nil {
		// Get the conversation history from the database
		history, err := c.messageService.GetByConversationID(conversationID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// messagesResponse = history
		for _, message := range history {
			messagesHistory = append(messagesHistory,
				openai.ChatCompletionMessage{
					Role:    openai.ChatMessageRoleUser,
					Content: message.Message,
				},
				openai.ChatCompletionMessage{
					Role:    openai.ChatMessageRoleAssistant,
					Content: *message.Response,
				},
			)
		}
	}

	reply, err := c.openaiService.SendMessage(request.Message, messagesHistory)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create a new UUID for conversation if it's not provided
	if conversationID == nil {
		newID := uuid.New()
		conversationID = &newID
	}

	// Save message and response to database
	err = c.messageService.CreateMessage(&entities.Message{
		ConversationID: conversationID,
		Response:       &reply,
		Message:        request.Message,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save message: " + err.Error()})
		return
	}

	responseMessages, err := c.messageService.GetByConversationID(conversationID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get messages: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, responseMessages)
}
