package containers

import "database/sql"

type AppContainer struct {
	Messages *MessageContainer
	ChatGPT  *ChatGPTContainer
	Chat     *ChatContainer
}

func Initialize(db *sql.DB) (*AppContainer, error) {
	messageContainer, err := InitializeMessageContainer(db)
	if err != nil {
		return nil, err
	}

	chatGPTContainer, err := InitializeChatGPTContainer()
	if err != nil {
		return nil, err
	}

	chatContainer, err := InitializeChatContainer(db)
	if err != nil {
		return nil, err
	}

	return &AppContainer{
		Messages: messageContainer,
		ChatGPT:  chatGPTContainer,
		Chat:     chatContainer,
	}, nil
}
