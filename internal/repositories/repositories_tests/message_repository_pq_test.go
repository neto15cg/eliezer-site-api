package repositories_tests

import (
	"database/sql"
	"regexp"
	"testing"
	"time"

	"app/internal/entities"
	"app/internal/repositories"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMessageRepositoryPg_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := repositories.NewMessageRepositoryPq(db)
	messageID := uuid.New()
	convID := uuid.New()
	now := time.Now()

	hello := "Hello"
	world := "World"
	message := &entities.Message{
		ID:             messageID,
		Message:        hello,
		Response:       &world,
		ConversationID: &convID,
	}

	// Set up expected query and mock response
	mock.ExpectQuery(regexp.QuoteMeta(`
		INSERT INTO messages (id, message, response, conversation_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
		RETURNING id, created_at, updated_at`)).
		WithArgs(messageID, hello, world, &convID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).
			AddRow(messageID, now, now))

	// Call the method being tested
	err = repo.Create(message)

	// Assert results
	assert.NoError(t, err)
	assert.Equal(t, messageID, message.ID)
	assert.NotNil(t, message.CreatedAt)
	assert.NotNil(t, message.UpdatedAt)
}

func TestMessageRepositoryPg_List(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := repositories.NewMessageRepositoryPq(db)
	messageID1 := uuid.New()
	messageID2 := uuid.New()
	convID := uuid.New()
	now := time.Now()

	hello := "Hello"
	world := "World"
	another := "Another"
	response := "Response"

	// Set up expected query and mock response
	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT id, message, response, conversation_id, created_at, updated_at
		FROM messages
		ORDER BY created_at DESC`)).
		WillReturnRows(sqlmock.NewRows([]string{"id", "message", "response", "conversation_id", "created_at", "updated_at"}).
			AddRow(messageID1, hello, world, &convID, now, now).
			AddRow(messageID2, another, response, &convID, now, now))

	// Call the method being tested
	messages, err := repo.List()

	// Assert results
	assert.NoError(t, err)
	assert.Len(t, messages, 2)
	assert.Equal(t, messageID1, messages[0].ID)
	assert.Equal(t, hello, messages[0].Message)
	assert.Equal(t, world, *messages[0].Response)
	assert.Equal(t, messageID2, messages[1].ID)
}

func TestMessageRepositoryPg_GetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := repositories.NewMessageRepositoryPq(db)
	messageID := uuid.New()
	now := time.Now()

	hello := "Hello"
	world := "World"

	// Set up expected query and mock response
	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT id, message, response, created_at, updated_at
		FROM messages
		WHERE id = $1`)).
		WithArgs(messageID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "message", "response", "created_at", "updated_at"}).
			AddRow(messageID, hello, world, now, now))

	// Call the method being tested
	message, err := repo.GetByID(messageID)

	// Assert results
	assert.NoError(t, err)
	assert.NotNil(t, message)
	assert.Equal(t, messageID, message.ID)
	assert.Equal(t, hello, message.Message)
	assert.Equal(t, world, *message.Response)
}

func TestMessageRepositoryPg_GetByID_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := repositories.NewMessageRepositoryPq(db)
	messageID := uuid.New()

	// Set up expected query and mock response for not found
	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT id, message, response, created_at, updated_at
		FROM messages
		WHERE id = $1`)).
		WithArgs(messageID).
		WillReturnError(sql.ErrNoRows)

	// Call the method being tested
	message, err := repo.GetByID(messageID)

	// Assert results
	assert.Error(t, err)
	assert.Nil(t, message)
}

func TestMessageRepositoryPg_GetByConversationID(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := repositories.NewMessageRepositoryPq(db)
	messageID1 := uuid.New()
	messageID2 := uuid.New()
	convID := uuid.New()
	now := time.Now()

	hello := "Hello"
	world := "World"
	another := "Another"
	response := "Response"

	// Set up expected query and mock response
	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT id, message, response, conversation_id, created_at, updated_at
		FROM messages
		WHERE conversation_id = $1
		ORDER BY created_at ASC
		LIMIT 100`)).
		WithArgs(&convID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "message", "response", "conversation_id", "created_at", "updated_at"}).
			AddRow(messageID1, hello, world, &convID, now, now).
			AddRow(messageID2, another, response, &convID, now, now))

	// Call the method being tested
	messages, err := repo.GetByConversationID(&convID)

	// Assert results
	assert.NoError(t, err)
	assert.Len(t, messages, 2)
	assert.Equal(t, messageID1, messages[0].ID)
	assert.Equal(t, hello, messages[0].Message)
	assert.Equal(t, world, *messages[0].Response)
	assert.Equal(t, messageID2, messages[1].ID)
	assert.Equal(t, response, *messages[1].Response)
}

func TestMessageRepositoryPg_GetByConversationID_NilID(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo := repositories.NewMessageRepositoryPq(db)

	// Call the method being tested with nil
	messages, err := repo.GetByConversationID(nil)

	// Assert results
	assert.NoError(t, err)
	assert.Nil(t, messages)

	// Ensure no database queries were executed
	assert.NoError(t, mock.ExpectationsWereMet())
}
