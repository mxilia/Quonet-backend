package repository

import (
	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/entities"
)

type PostRepository interface {
	Save(post *entities.Post) error
	/* No private posts involved */
	Find(authorID uuid.UUID, threadID uuid.UUID, title string, offset int, limit int) ([]*entities.Post, error)
	FindByID(id uuid.UUID) (*entities.Post, error)
	/* Private posts involved */
	FindPrivate(authorID uuid.UUID, threadID uuid.UUID, title string, offset int, limit int) ([]*entities.Post, error)
	FindPrivateByID(id uuid.UUID) (*entities.Post, error)
	Count(isPrivate bool, authorID uuid.UUID, threadID uuid.UUID, title string) (int64, error)
	Patch(id uuid.UUID, post *entities.Post) error
	Delete(id uuid.UUID) error
}
