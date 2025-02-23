package handlers

import (
	"encoding/json"
	"net/http"

	"app/internal/domain"
	"app/models"
)

type ChatGPTHandler struct {
	chatGPTService domain.ChatGPTService
}

func NewChatGPTHandler(service domain.ChatGPTService) *ChatGPTHandler {
	return &ChatGPTHandler{
		chatGPTService: service,
	}
}

func (h *ChatGPTHandler) SendMessage(w http.ResponseWriter, r *http.Request) {
	var req models.ChatGPTRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := h.chatGPTService.SendMessage(req.Message)
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
