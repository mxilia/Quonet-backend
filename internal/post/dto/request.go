package dto

type CreatePostRequest struct {
	Title        string `json:"title" validate:"required"`
	AuthorID     string `json:"author_id" validate:"required"`
	ThreadID     string `json:"thread_id" validate:"required"`
	Content      string `json:"content" validate:"required"`
	ThumbnailUrl string `json:"thumbnail_url"`
}

type PostPatchRequest struct {
	ThumbnailUrl string `json:"thumbnail_url"`
	IsPrivate    bool   `json:"is_private"`
}
