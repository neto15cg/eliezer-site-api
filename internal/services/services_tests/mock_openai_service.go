package services_tests

import (
	"app/internal/interfaces"

	"github.com/sashabaranov/go-openai"
	"github.com/stretchr/testify/mock"
)

// MockOpenAIService is a mock implementation of the OpenAI service
type MockOpenAIService struct {
	mock.Mock
}

// Ensure MockOpenAIService implements OpenAIServiceInterface
var _ interfaces.OpenAIServiceInterface = (*MockOpenAIService)(nil)

// NewMockOpenAIService creates a new instance of MockOpenAIService
func NewMockOpenAIService() *MockOpenAIService {
	return &MockOpenAIService{}
}

// SendMessage mocks sending a message to OpenAI and getting a response
func (m *MockOpenAIService) SendMessage(message string, history []openai.ChatCompletionMessage) (string, error) {
	args := m.Called(message, history)
	return args.String(0), args.Error(1)
}
