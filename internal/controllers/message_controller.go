package controllers

import (
	"net/http"

	"app/internal/entities"
	"app/internal/interfaces"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MessageController struct {
	service interfaces.MessageServiceInterface
}

func NewMessageController(service interfaces.MessageServiceInterface) *MessageController {
	return &MessageController{service: service}
}

func (c *MessageController) GetMessages(ctx *gin.Context) {
	messages, err := c.service.GetMessages()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, messages)
}

func (c *MessageController) CreateMessage(ctx *gin.Context) {
	var message entities.Message
	if err := ctx.ShouldBindJSON(&message); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateMessage(&message); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, message)
}

func (c *MessageController) GetByConversationID(ctx *gin.Context) {
	conversationID, err := uuid.Parse(ctx.Param("conversation_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid conversation id format"})
		return
	}
	messages, err := c.service.GetByConversationID(&conversationID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, messages)
}
