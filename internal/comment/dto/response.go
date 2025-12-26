package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/entities"
)

type CommentResponse struct {
	ID        uuid.UUID  `json:"id"`
	AuthorID  *uuid.UUID `json:"author_id"`
	Content   string     `json:"content"`
	LikeCount int64      `json:"like_count"`

	ParentID *uuid.UUID `json:"parent_id"`
	RootID   uuid.UUID  `json:"root_id"`

	Author   *entities.User     `json:"author"`
	Likes    []entities.Like    `json:"likes"`
	Comments []entities.Comment `json:"comments"`

	CreatedAt time.Time `json:"created_at"`
}
