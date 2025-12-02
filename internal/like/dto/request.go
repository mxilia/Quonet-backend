package dto

type CreateLikeRequest struct {
	OwnerID    string `json:"owner_id" validate:"required"`
	ParentID   string `json:"parent_id" validate:"required"`
	ParentType string `json:"parent_type" validate:"required"`
	IsPositive bool   `json:"is_positive" validate:"required"`
}
