package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/mxilia/Conflux-backend/internal/entities"
)

/* Will modify this */

type PostResponse struct {
	ID           uuid.UUID `json:"id"`
	Title        string    `json:"title"`
	AuthorID     uuid.UUID `json:"author_id"`
	ThreadID     uuid.UUID `json:"thread_id"`
	Content      string    `json:"content"`
	ThumbnailUrl string    `json:"thumbnail_url"`
	IsPrivate    bool      `json:"is_private"`

	Author   entities.User      `json:"author"`
	Likes    []entities.Like    `json:"like"`
	Comments []entities.Comment `json:"comments"`

	CreatedAt time.Time `json:"created_at"`
}
