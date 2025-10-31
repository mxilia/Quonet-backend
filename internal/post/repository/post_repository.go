package repository

import (
	"github.com/google/uuid"
	"github.com/mxilia/Conflux-backend/internal/entities"
)

type PostRepository interface {
	Save(post *entities.Post) error
	/* No private posts involved */
	FindAll() ([]*entities.Post, error)
	FindByAuthorID(id uuid.UUID) ([]*entities.Post, error)
	FindByThreadID(id uuid.UUID) ([]*entities.Post, error)
	FindByID(id uuid.UUID) (*entities.Post, error)
	FindByTitle(title string) (*entities.Post, error)
	/* Private posts involved */
	FindAllCoverPrivate() ([]*entities.Post, error)
	FindAllPrivate() ([]*entities.Post, error)
	FindCoverPrivateByAuthorID(id uuid.UUID) ([]*entities.Post, error)
	FindPrivateByAuthorID(id uuid.UUID) ([]*entities.Post, error)
	FindCoverPrivateByThreadID(id uuid.UUID) ([]*entities.Post, error)
	FindPrivateByThreadID(id uuid.UUID) ([]*entities.Post, error)
	FindPrivateByID(id uuid.UUID) (*entities.Post, error)
	FindPrivateByTitle(title string) (*entities.Post, error)
	Patch(id uuid.UUID, post *entities.Post) error
	Delete(id uuid.UUID) error
}
