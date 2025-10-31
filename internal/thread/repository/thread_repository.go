package repository

import (
	"github.com/google/uuid"
	"github.com/mxilia/Conflux-backend/internal/entities"
)

type ThreadRepository interface {
	Save(thread *entities.Thread) error
	FindAll() ([]*entities.Thread, error)
	FindByID(id uuid.UUID) (*entities.Thread, error)
	Delete(id uuid.UUID) error
}
