package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/entities"
)

type LikeUseCase interface {
	CreateLike(ctx context.Context, like *entities.Like) error
	FindAllLikes(parentType string) ([]*entities.Like, error)
	FindLikesByOwnerID(parentType string, id uuid.UUID) ([]*entities.Like, error)
	FindLikesByParentID(parentType string, id uuid.UUID) ([]*entities.Like, error)
	FindLikeByParentIDAndOwnerID(parentType string, parentID uuid.UUID, ownerID uuid.UUID) (*entities.Like, error)
	FindLikeByID(parentType string, id uuid.UUID) (*entities.Like, error)
	LikeCountByParentID(parentType string, id uuid.UUID) (int64, error)
	IsParentLikedByMe(parentType string, parentID uuid.UUID, ownerID uuid.UUID) (bool, error)
	DeleteLike(parentType string, id uuid.UUID) error
}
