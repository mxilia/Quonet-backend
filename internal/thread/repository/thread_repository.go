package repository

import "github.com/mxilia/Conflux-backend/internal/entities"

type ThreadRepository interface {
	Save(thread *entities.Thread) error
	FindAll() ([]*entities.Thread, error)
	FindByID(id uint) (*entities.Thread, error)
	Delete(id uint) error
}
