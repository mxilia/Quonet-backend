package dto

import "github.com/google/uuid"

type LikeResponse struct {
	ID         uuid.UUID `json:"id"`
	OwnerID    uuid.UUID `json:"owner_id"`
	ParentID   uuid.UUID `json:"parent_id"`
	ParentType string    `json:"parent_type"`
}

type LikeCountResponse struct {
	Count int64 `json:"like_count"`
}

type IsLikedResponse struct {
	IsLiked bool `json:"is_liked"`
}
