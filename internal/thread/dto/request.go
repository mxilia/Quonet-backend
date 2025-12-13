package dto

type CreateThreadRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	ImageUrl    string `json:"image_url,omitempty"`
}
