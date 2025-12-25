package dto

import (
	"github.com/mxilia/Quonet-backend/internal/entities"
	"github.com/mxilia/Quonet-backend/pkg/database"
)

func ToPostResponse(post *entities.Post, storageService *database.StorageService) *PostResponse {
	ThumbnailUrl := post.ThumbnailUrl
	if ThumbnailUrl != "" {
		var err error
		ThumbnailUrl, err = storageService.GetSignedURL(ThumbnailUrl)
		if err != nil {
			ThumbnailUrl = ""
		}
	}

	return &PostResponse{
		ID:           post.ID,
		Title:        post.Title,
		AuthorID:     post.AuthorID,
		ThreadID:     post.ThreadID,
		Content:      post.Content,
		ThumbnailUrl: ThumbnailUrl,
		IsPrivate:    post.IsPrivate,
		LikeCount:    post.LikeCount,

		Author:   post.Author,
		Thread:   post.Thread,
		Likes:    post.Likes,
		Comments: post.Comments,

		CreatedAt: post.CreatedAt,
	}
}

func ToPostResponseList(posts []*entities.Post, supabase *database.StorageService) []*PostResponse {
	res := make([]*PostResponse, 0, len(posts))
	for _, post := range posts {
		res = append(res, ToPostResponse(post, supabase))
	}
	return res
}

func FromPostPatchRequest(req *PostPatchRequest) *entities.Post {
	return &entities.Post{
		ThumbnailUrl: req.ThumbnailUrl,
		IsPrivate:    req.IsPrivate,
	}
}
