package repository

import (
	"github.com/google/uuid"
	"github.com/mxilia/Conflux-backend/internal/entities"
)

type CommentRepository interface {
	Save(comment *entities.Comment) error
	FindAll() ([]*entities.Comment, error)
	FindByAuthorID(id uuid.UUID) ([]*entities.Comment, error)
	FindByParentID(id uuid.UUID) ([]*entities.Comment, error)
	FindByRootID(id uuid.UUID) ([]*entities.Comment, error)
	FindByID(id uuid.UUID) (*entities.Comment, error)
	Patch(id uuid.UUID, comment *entities.Comment) error
	Delete(id uuid.UUID) error
}
