package postgres

import (
	"database/sql"

	"app/internal/domain"
	"app/models"

	"github.com/google/uuid"
)

type messageRepository struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) domain.MessageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) Create(message *models.Message) error {
	query := `
		INSERT INTO messages (id, message, created_at, updated_at)
		VALUES ($1, $2, NOW(), NOW())
		RETURNING id, created_at, updated_at`

	return r.db.QueryRow(query, message.ID, message.Message).
		Scan(&message.ID, &message.CreatedAt, &message.UpdatedAt)
}

func (r *messageRepository) List() ([]models.Message, error) {
	query := `
		SELECT id, message, created_at, updated_at
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
		if err := rows.Scan(&m.ID, &m.Message, &m.CreatedAt, &m.UpdatedAt); err != nil {
			return []models.Message{}, err
		}
		messages = append(messages, m)
	}

	return messages, nil
}

func (r *messageRepository) GetByID(id uuid.UUID) (*models.Message, error) {
	var m models.Message
	query := `
		SELECT id, message, created_at, updated_at
		FROM messages
		WHERE id = $1`

	err := r.db.QueryRow(query, id).
		Scan(&m.ID, &m.Message, &m.CreatedAt, &m.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
