package usecase

import (
	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/entities"
)

type CommentUseCase interface {
	CreateComment(comment *entities.Comment) error
	FindComments(authorID uuid.UUID, parentID uuid.UUID, rootID uuid.UUID, page int, limit int) ([]*entities.Comment, int64, error)
	FindCommentByID(id uuid.UUID) (*entities.Comment, error)
	PatchComment(id uuid.UUID, comment *entities.Comment) error
	DeleteComment(id uuid.UUID) error
}
