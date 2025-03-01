package entities

type ChatGPTRequest struct {
	Message string `json:"message"`
}

type ChatGPTResponse struct {
	Response string `json:"response"`
}
