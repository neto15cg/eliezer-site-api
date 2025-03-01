package repositories

import "app/internal/entities"

type MessageRepository interface {
	FindAll() ([]entities.Message, error)
	FindById(id string) (*entities.Message, error)
	Save(user *entities.Message) error
}
