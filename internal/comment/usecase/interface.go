package usecase

import (
	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/entities"
)

type CommentUseCase interface {
	CreateComment(comment *entities.Comment) error
	FindAllComments() ([]*entities.Comment, error)
	FindCommentsByAuthorID(id uuid.UUID) ([]*entities.Comment, error)
	FindCommentsByParentID(id uuid.UUID) ([]*entities.Comment, error)
	FindCommentsByRootID(id uuid.UUID) ([]*entities.Comment, error)
	FindCommentByID(id uuid.UUID) (*entities.Comment, error)
	PatchComment(id uuid.UUID, comment *entities.Comment) error
	DeleteComment(id uuid.UUID) error
}
