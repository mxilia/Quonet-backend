package dto

import (
	"github.com/google/uuid"
	"github.com/mxilia/Quonet-backend/internal/entities"
)

func ToCommentResponse(comment *entities.Comment) *CommentResponse {
	return &CommentResponse{
		ID:        comment.ID,
		AuthorID:  comment.AuthorID,
		Content:   comment.Content,
		LikeCount: comment.LikeCount,

		ParentID: comment.ParentID,
		RootID:   comment.RootID,

		Author:   comment.Author,
		Likes:    comment.Likes,
		Comments: comment.Comments,

		CreatedAt: comment.CreatedAt,
	}
}

func ToCommentResponseList(comments []*entities.Comment) []*CommentResponse {
	res := make([]*CommentResponse, 0, len(comments))
	for _, comment := range comments {
		res = append(res, ToCommentResponse(comment))
	}
	return res
}

func FromCommentCreateRequest(authorID uuid.UUID, req *CreateCommentRequest) (*entities.Comment, error) {
	var parentID *uuid.UUID
	if req.ParentID != "" {
		id, err := uuid.Parse(req.ParentID)
		if err != nil {
			return nil, err
		}
		parentID = &id
	} else {
		parentID = nil
	}

	rootID, err := uuid.Parse(req.RootID)
	if err != nil {
		return nil, err
	}

	return &entities.Comment{
		AuthorID: &authorID,
		Content:  req.Content,
		ParentID: parentID,
		RootID:   rootID,
	}, nil
}

func FromCommentPatchRequest(req *CommentPatchRequest) *entities.Comment {
	return &entities.Comment{
		Content: req.Content,
	}
}
