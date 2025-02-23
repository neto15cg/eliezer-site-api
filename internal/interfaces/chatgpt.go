package interfaces

type ChatGPTService interface {
	SendMessage(message string) (string, error)
}
