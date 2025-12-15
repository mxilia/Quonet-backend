package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/entities"
)

type CommentRepository interface {
	Save(comment *entities.Comment) error
	Find(authorID uuid.UUID, parentID uuid.UUID, rootID uuid.UUID, offset int, limit int) ([]*entities.Comment, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entities.Comment, error)
	Count() (int64, error)
	Patch(ctx context.Context, id uuid.UUID, comment *entities.Comment) error
	Delete(id uuid.UUID) error
}
