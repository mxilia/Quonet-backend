package dto

type CreateLikeRequest struct {
	ParentID   string `json:"parent_id" validate:"required"`
	ParentType string `json:"parent_type" validate:"required"`
	IsPositive bool   `json:"is_positive" validate:"required"`
}
