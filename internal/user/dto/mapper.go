package dto

import (
	"github.com/mxilia/Conflux-backend/internal/entities"
	"github.com/mxilia/Conflux-backend/pkg/token"
)

func ToUserResponse(user *entities.User) *UserResponse {
	return &UserResponse{
		ID:          user.ID,
		Handler:     user.Handler,
		Email:       user.Email,
		ProfileUrl:  user.ProfileUrl,
		Role:        user.Role,
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

func ToLoginResponse(user *entities.User, accessToken string, acccessClaims *token.UserClaims) *LoginResponse {
	return &LoginResponse{
		Handler:    user.Handler,
		Email:      user.Email,
		ProfileUrl: user.ProfileUrl,
		Role:       user.Role,

		AccessToken:          accessToken,
		AccessTokenExpiresAt: acccessClaims.ExpiresAt.Time,
	}
}

func FromCreateUserByGoogleRequest(req *CreateUserByGoogleRequest) *entities.User {
	return &entities.User{
		Email:      req.Email,
		ProfileUrl: req.ProfileUrl,
	}
}

func FromUserPatchRequest(req *UserPatchRequest) *entities.User {
	return &entities.User{
		Handler:    req.Handler,
		ProfileUrl: req.ProfileUrl,
		Role:       req.Role,
	}
}
