package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/entities"
)

type UserResponse struct {
	ID          uuid.UUID `json:"id"`
	Handler     string    `json:"handler"`
	Email       string    `json:"email"`
	ProfileUrl  string    `json:"profile_url"`
	Role        string    `json:"role"`
	IsBanned    bool      `json:"is_banned"`
	BannedUntil time.Time `json:"banned_until"`

	Posts    []entities.Post    `json:"posts"`
	Comments []entities.Comment `json:"comments"`
	Likes    []entities.Like    `json:"likes"`

	CreatedAt time.Time `json:"created_at"`
}

type LoginResponse struct {
	ID         uuid.UUID `json:"id"`
	Handler    string    `json:"handler"`
	Email      string    `json:"email"`
	ProfileUrl string    `json:"profile_url"`
	Role       string    `json:"role"`

	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}
