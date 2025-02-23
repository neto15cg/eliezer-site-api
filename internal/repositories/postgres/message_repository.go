package postgres

import (
	"database/sql"

	"app/internal/interfaces"
	"app/models"

	"github.com/google/uuid"
)

type messageRepository struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) interfaces.MessageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) Create(message *models.Message) error {
	query := `
		INSERT INTO messages (id, message, response, conversation_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
		RETURNING id, created_at, updated_at`

	return r.db.QueryRow(query, message.ID, message.Message, message.Response, message.ConversationID).
		Scan(&message.ID, &message.CreatedAt, &message.UpdatedAt)
}

func (r *messageRepository) List() ([]models.Message, error) {
	query := `
		SELECT id, message, response, conversation_id, created_at, updated_at
		FROM messages
		ORDER BY created_at DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return []models.Message{}, err
	}
	defer rows.Close()

	messages := make([]models.Message, 0)
	for rows.Next() {
		var m models.Message
		if err := rows.Scan(&m.ID, &m.Message, &m.Response, &m.ConversationID, &m.CreatedAt, &m.UpdatedAt); err != nil {
			return []models.Message{}, err
		}
		messages = append(messages, m)
	}

	return messages, nil
}

func (r *messageRepository) GetByID(id uuid.UUID) (*models.Message, error) {
	var m models.Message
	query := `
		SELECT id, message, response, created_at, updated_at
		FROM messages
		WHERE id = $1`

	err := r.db.QueryRow(query, id).
		Scan(&m.ID, &m.Message, &m.Response, &m.CreatedAt, &m.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *messageRepository) GetByConversationID(conversationID uuid.UUID) ([]models.Message, error) {
	query := `
		SELECT id, message, response, conversation_id, created_at, updated_at
		FROM messages
		WHERE conversation_id = $1
		ORDER BY created_at ASC
		LIMIT 100`

	rows, err := r.db.Query(query, conversationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	messages := make([]models.Message, 0)
	for rows.Next() {
		var m models.Message
		if err := rows.Scan(&m.ID, &m.Message, &m.Response, &m.ConversationID, &m.CreatedAt, &m.UpdatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, m)
	}

	return messages, nil
}
