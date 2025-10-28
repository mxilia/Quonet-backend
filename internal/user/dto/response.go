package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/mxilia/Conflux-backend/internal/entities"
)

type UserResponse struct {
	ID          uuid.UUID `json:"id"`
	Handler     string    `json:"handler"`
	Email       string    `json:"email"`
	ProfileUrl  string    `json:"profile_url"`
	IsAdmin     bool      `json:"is_admin"`
	IsBanned    bool      `json:"is_banned"`
	BannedUntil time.Time `json:"banned_until"`

	Posts    []entities.Post    `json:"posts"`
	Comments []entities.Comment `json:"comments"`

	CreatedAt time.Time `json:"created_at"`
}
