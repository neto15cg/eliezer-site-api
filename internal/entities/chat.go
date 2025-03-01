package entities

import "github.com/google/uuid"

type OpenAIRequest struct {
	Message        string     `json:"message"`
	ConversationID *uuid.UUID `json:"conversation_id"`
}
