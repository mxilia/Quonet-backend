package dto

type CreateCommentRequest struct {
	AuthorID string `json:"author_id" validate:"required"`
	Content  string `json:"content" validate:"required"`
	ParentID string `json:"parent_id"`
	RootID   string `json:"root_id" validate:"required"`
}

type CommentPatchRequest struct {
	Content string `json:"content" validate:"required"`
}
