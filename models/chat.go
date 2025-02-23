package models

import "github.com/google/uuid"

type ChatRequest struct {
	Message        string    `json:"message"`
	ConversationID uuid.UUID `json:"conversation_id"`
}
