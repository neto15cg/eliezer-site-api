package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Validator interface {
	Validate() error
}

type Message struct {
	ID             uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Message        string     `json:"message" binding:"required"`
	Response       *string    `json:"response"`
	ConversationID *uuid.UUID `json:"conversation_id"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

func (m *Message) Validate() error {
	if m.Message == "" {
		return errors.New("message cannot be empty")
	}
	return nil
}

// EnsureID makes sure the message has a valid UUID
func (m *Message) EnsureID() {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
}
