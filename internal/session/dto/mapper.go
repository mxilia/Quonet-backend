package dto

import "github.com/mxilia/Conflux-backend/pkg/token"

func ToRenewAccessTokenResponse(accessToken string, accessClaims *token.UserClaims) *RenewAccessTokenResponse {
	return &RenewAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessClaims.ExpiresAt.Time,
	}
}
