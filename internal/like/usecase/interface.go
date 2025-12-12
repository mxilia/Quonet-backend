package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/entities"
)

type LikeUseCase interface {
	CreateLike(ctx context.Context, like *entities.Like) error
	FindLikes(parentType string, ownerID uuid.UUID, parentID uuid.UUID, page int, limit int) ([]*entities.Like, int64, error)
	FindLikeByID(id uuid.UUID) (*entities.Like, error)
	Count(parentType string, ownerID uuid.UUID, parentID uuid.UUID) (int64, error)
	DeleteLike(id uuid.UUID) error
}
