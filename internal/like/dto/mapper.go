package dto

import "github.com/mxilia/Conflux-backend/internal/entities"

func ToLikeResponse(like *entities.Like) *LikeResponse {
	return &LikeResponse{
		ID:         like.ID,
		OwnerID:    like.OwnerID,
		ParentID:   like.ParentID,
		ParentType: like.ParentType,
	}
}

func ToLikeResponseList(likes []*entities.Like) []*LikeResponse {
	res := make([]*LikeResponse, 0, len(likes))
	for _, like := range likes {
		res = append(res, ToLikeResponse(like))
	}
	return res
}

func ToLikeCountResponse(count int64) *LikeCountResponse {
	return &LikeCountResponse{
		Count: count,
	}
}

func ToIsLikedResponse(isLiked bool) *IsLikedResponse {
	return &IsLikedResponse{
		IsLiked: isLiked,
	}
}
