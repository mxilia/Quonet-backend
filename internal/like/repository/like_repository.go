package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/mxilia/Conflux-backend/internal/entities"
)

type LikeRepository interface {
	Save(ctx context.Context, like *entities.Like) error
	FindAll(parentType string) ([]*entities.Like, error)
	FindByOwnerID(parentType string, id uuid.UUID) ([]*entities.Like, error)
	FindByParentID(parentType string, id uuid.UUID) ([]*entities.Like, error)
	FindByParentIDAndOwnerID(parentType string, parentID uuid.UUID, ownerID uuid.UUID) (*entities.Like, error)
	FindByID(parentType string, id uuid.UUID) (*entities.Like, error)
	CountByParentID(parentType string, id uuid.UUID) (int64, error)
	IsParentLikedByMe(parentType string, parentID uuid.UUID, ownerID uuid.UUID) (bool, error)
	Delete(ctx context.Context, parentType string, id uuid.UUID) error
}
