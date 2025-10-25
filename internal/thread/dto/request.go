package dto

type CreateThreadRequest struct {
	Title string `json:"title" validate:"required"`
}
