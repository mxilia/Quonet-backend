package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/entities"
)

type PostRepository interface {
	Save(post *entities.Post) error
	/* No private posts involved */
	Find(authorID uuid.UUID, threadID uuid.UUID, title string, offset int, limit int) ([]*entities.Post, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entities.Post, error)
	FindTopLiked(authorID uuid.UUID, threadID uuid.UUID, title string, limit int) ([]*entities.Post, error)
	/* Private posts involved */
	FindPrivate(authorID uuid.UUID, threadID uuid.UUID, title string, offset int, limit int) ([]*entities.Post, error)
	FindPrivateByID(id uuid.UUID) (*entities.Post, error)
	Count(isPrivate bool, authorID uuid.UUID, threadID uuid.UUID, title string) (int64, error)
	FindNoFilterByID(id uuid.UUID) (*entities.Post, error)
	Patch(ctx context.Context, id uuid.UUID, post *entities.Post) error
	Delete(id uuid.UUID) error
}
