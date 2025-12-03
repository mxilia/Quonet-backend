package dto

import "github.com/mxilia/Quonet-backend/internal/entities"

func ToPostResponse(post *entities.Post) *PostResponse {
	return &PostResponse{
		ID:           post.ID,
		Title:        post.Title,
		AuthorID:     post.AuthorID,
		ThreadID:     post.ThreadID,
		Content:      post.Content,
		ThumbnailUrl: post.ThumbnailUrl,
		IsPrivate:    post.IsPrivate,
		LikeCount:    post.LikeCount,

		Author:   post.Author,
		Likes:    post.Likes,
		Comments: post.Comments,

		CreatedAt: post.CreatedAt,
	}
}

func ToPostResponseList(posts []*entities.Post) []*PostResponse {
	res := make([]*PostResponse, 0, len(posts))
	for _, post := range posts {
		res = append(res, ToPostResponse(post))
	}
	return res
}

func FromPostPatchRequest(req *PostPatchRequest) *entities.Post {
	return &entities.Post{
		ThumbnailUrl: req.ThumbnailUrl,
		IsPrivate:    req.IsPrivate,
	}
}
