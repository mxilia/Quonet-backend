package usecase

import (
	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/entities"
)

type PostUseCase interface {
	CreatePost(post *entities.Post) error
	/* No private posts involved */
	FindPosts(authorID uuid.UUID, threadID uuid.UUID, title string, page int, limit int) ([]*entities.Post, int64, error)
	FindPostByID(id uuid.UUID) (*entities.Post, error)
	/* Private posts involved */
	FindPrivatePosts(authorID uuid.UUID, threadID uuid.UUID, title string, page int, limit int) ([]*entities.Post, int64, error)
	FindPrivatePostByID(id uuid.UUID) (*entities.Post, error)
	/* General posts */
	FindTopLikedPosts(authorID uuid.UUID, threadID uuid.UUID, title string, limit int) ([]*entities.Post, error)
	FindNoFilterPostByID(id uuid.UUID) (*entities.Post, error)
	PatchPost(id uuid.UUID, post *entities.Post) error
	DeletePost(id uuid.UUID) error
}
