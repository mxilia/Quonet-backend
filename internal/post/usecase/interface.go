package usecase

import (
	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/entities"
)

type PostUseCase interface {
	CreatePost(post *entities.Post) error
	/* No private posts involved */
	FindAllPosts() ([]*entities.Post, error)
	FindPostsByAuthorID(id uuid.UUID) ([]*entities.Post, error)
	FindPostsByThreadID(id uuid.UUID) ([]*entities.Post, error)
	FindPostByID(id uuid.UUID) (*entities.Post, error)
	FindPostByTitle(title string) (*entities.Post, error)
	/* Private posts involved */
	FindAllPostsCoverPrivate() ([]*entities.Post, error)
	FindAllPrivatePosts() ([]*entities.Post, error)
	FindPostsCoverPrivateByAuthorID(id uuid.UUID) ([]*entities.Post, error)
	FindPrivatePostsByAuthorID(id uuid.UUID) ([]*entities.Post, error)
	FindPostsCoverPrivateByThreadID(id uuid.UUID) ([]*entities.Post, error)
	FindPrivatePostsByThreadID(id uuid.UUID) ([]*entities.Post, error)
	FindPrivatePostByID(id uuid.UUID) (*entities.Post, error)
	FindPrivatePostByTitle(title string) (*entities.Post, error)
	PatchPost(id uuid.UUID, post *entities.Post) error
	DeletePost(id uuid.UUID) error
}
