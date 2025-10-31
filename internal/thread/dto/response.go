package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/mxilia/Conflux-backend/internal/entities"
)

type ThreadResponse struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`

	Posts []entities.Post `json:"posts"`

	CreatedAt time.Time `json:"created_at"`
}
