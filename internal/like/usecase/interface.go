package usecase

import (
	"github.com/google/uuid"
	"github.com/mxilia/Conflux-backend/internal/entities"
)

type LikeUseCase interface {
	CreateLike(like *entities.Like) error
	FindAllLikes() ([]*entities.Like, error)
	FindLikesByOwnerID(id uuid.UUID) ([]*entities.Like, error)
	FindLikesByParentID(id uuid.UUID) ([]*entities.Like, error)
	FindLikeByID(id uuid.UUID) (*entities.Like, error)
	LikeCountByParentID(parentType string, id uuid.UUID) (int64, error)
	IsParentLikedByMe(parentType string, parentID uuid.UUID, myID uuid.UUID) (bool, error)
	DeleteLike(id uuid.UUID) error
}
