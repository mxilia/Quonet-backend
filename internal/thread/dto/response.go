package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/entities"
)

type ThreadResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"image_url"`

	Posts []entities.Post `json:"posts"`

	CreatedAt time.Time `json:"created_at"`
}
