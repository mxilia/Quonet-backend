package usecase

import (
	"context"
	"io"

	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/entities"
)

type PostUseCase interface {
	CreatePost(ctx context.Context, post *entities.Post, file io.Reader, filename string, contentType string) error
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
	DeletePost(ctx context.Context, id uuid.UUID) error
}
