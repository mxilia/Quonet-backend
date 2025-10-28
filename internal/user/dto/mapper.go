package dto

import "github.com/mxilia/Conflux-backend/internal/entities"

func ToUserResponse(user *entities.User) *UserResponse {
	return &UserResponse{
		ID:          user.ID,
		Handler:     user.Handler,
		Email:       user.Email,
		ProfileUrl:  user.ProfileUrl,
		IsAdmin:     user.IsAdmin,
		IsBanned:    user.IsBanned,
		BannedUntil: user.BannedUntil,

		Posts:    user.Posts,
		Comments: user.Comments,

		CreatedAt: user.CreatedAt,
	}
}

func ToUserResponseList(users []*entities.User) []*UserResponse {
	res := make([]*UserResponse, 0, len(users))
	for _, t := range users {
		res = append(res, ToUserResponse(t))
	}
	return res
}

func FromUserPatchRequest(user *UserPatchRequest) *entities.User {
	return &entities.User{
		Handler:    user.Handler,
		ProfileUrl: user.ProfileUrl,
	}
}
