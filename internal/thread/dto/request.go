package dto

type CreateThreadRequest struct {
	Title    string `json:"title" validate:"required"`
	ImageUrl string `json:"image_url,omitempty"`
}
