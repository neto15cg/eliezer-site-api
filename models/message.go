package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Validator interface {
	Validate() error
}

type Message struct {
	ID             uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Message        string    `json:"message" binding:"required"`
	ConversationID uuid.UUID `json:"conversation_id" binding:"required"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (m *Message) Validate() error {
	if m.Message == "" {
		return errors.New("message cannot be empty")
	}
	return nil
}

// Factory method for Message
func NewMessage(content string, conversationID uuid.UUID) *Message {
	return &Message{
		ID:             uuid.New(),
		Message:        content,
		ConversationID: conversationID,
	}
}
