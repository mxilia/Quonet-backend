package repository

import "github.com/mxilia/Conflux-backend/internal/entities"

type ThreadRepository interface {
	Save(thread *entities.Thread) error
	GetAll() ([]*entities.Thread, error)
	GetByID(id uint) (*entities.Thread, error)
	Delete(id uint) error
}
