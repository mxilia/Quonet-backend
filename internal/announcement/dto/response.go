package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/entities"
)

type AnnouncementResponse struct {
	ID        uuid.UUID     `json:"id"`
	AuthorID  uuid.UUID     `json:"author_id"`
	Content   string        `json:"content"`
	Author    entities.User `json:"author"`
	CreatedAt time.Time     `json:"created_at"`
}
