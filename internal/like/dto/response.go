package dto

import "github.com/google/uuid"

type LikeResponse struct {
	ID         uuid.UUID `json:"id"`
	OwnerID    uuid.UUID `json:"owner_id"`
	ParentID   uuid.UUID `json:"parent_id"`
	ParentType string    `json:"parent_type"`
	IsPositive bool      `json:"is_positive"`
}

type LikeCountResponse struct {
	Count int64 `json:"like_count"`
}

type LikeStateResponse struct {
	IsLiked        bool `json:"is_liked"`
	IsLikePositive bool `json:"is_like_positive"`
}
