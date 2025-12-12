package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/entities"
)

type LikeRepository interface {
	Save(ctx context.Context, like *entities.Like) error
	Find(parentType string, ownerID uuid.UUID, parentID uuid.UUID, offset int, limit int) ([]*entities.Like, error)
	FindByID(id uuid.UUID) (*entities.Like, error)
	Count(parentType string, ownerID uuid.UUID, parentID uuid.UUID) (int64, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
