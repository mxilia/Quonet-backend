package dto

type CreatePostRequest struct {
	Title    string `json:"title" validate:"required"`
	ThreadID string `json:"thread_id" validate:"required"`
	Content  string `json:"content" validate:"required"`
}

type PostPatchRequest struct {
	ThumbnailUrl string `json:"thumbnail_url,omitempty"`
	IsPrivate    bool   `json:"is_private,omitempty"`
}
