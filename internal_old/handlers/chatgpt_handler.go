package handlers

import (
	"encoding/json"
	"net/http"

	"app/models"

	openai "github.com/sashabaranov/go-openai"
)

func (h *ChatGPTHandler) SendMessage(w http.ResponseWriter, r *http.Request) {
	var req models.ChatGPTRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	messagesHistory := make([]openai.ChatCompletionMessage, 0)

	response, err := h.chatGPTService.SendMessage(req.Message, messagesHistory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := models.ChatGPTResponse{
		Response: response,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
