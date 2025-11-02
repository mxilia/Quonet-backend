package repository

import (
	"github.com/google/uuid"
	"github.com/mxilia/Conflux-backend/internal/entities"
)

type LikeRepository interface {
	Save(like *entities.Like) error
	FindAll() ([]*entities.Like, error)
	FindByOwnerID(id uuid.UUID) ([]*entities.Like, error)
	FindByParentID(id uuid.UUID) ([]*entities.Like, error)
	FindByID(id uuid.UUID) (*entities.Like, error)
	CountByParentID(parentType string, id uuid.UUID) (int64, error)
	IsParentLikedByMe(parentType string, parentID uuid.UUID, ownerID uuid.UUID) (bool, error)
	Delete(id uuid.UUID) error
}
