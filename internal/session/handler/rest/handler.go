package rest

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mxilia/Quonet-backend/internal/session/dto"
	"github.com/mxilia/Quonet-backend/internal/session/usecase"
	userUseCase "github.com/mxilia/Quonet-backend/internal/user/usecase"
	appError "github.com/mxilia/Quonet-backend/pkg/apperror"
	"github.com/mxilia/Quonet-backend/pkg/config"
	"github.com/mxilia/Quonet-backend/pkg/responses"
	"github.com/mxilia/Quonet-backend/pkg/token"
)

type HttpSessionHandler struct {
	usecase     usecase.SessionUseCase
	userUseCase userUseCase.UserUseCase
	tokenMaker  *token.JWTMaker
}

func NewHttpSessionHandler(usecase usecase.SessionUseCase, userUseCase userUseCase.UserUseCase, cfg *config.Config) *HttpSessionHandler {
	return &HttpSessionHandler{
		usecase:     usecase,
		userUseCase: userUseCase,
		tokenMaker:  token.NewJWTMaker(cfg),
	}
}

func removeToken(c *fiber.Ctx) {
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		Domain:   os.Getenv("DOMAIN"),
	})
}

func (h *HttpSessionHandler) RenewToken(c *fiber.Ctx) error {
	tokenStr := c.Cookies("token")
	claims, err := h.tokenMaker.VerifyToken(tokenStr)
	if err != nil {
		removeToken(c)
		return responses.ErrorWithMessage(c, appError.ErrUnauthorized, "failed to parse "+tokenStr)
	}

	session, err := h.usecase.FindSessionByID(claims.RegisteredClaims.ID)
	if err != nil {
		removeToken(c)
		return responses.ErrorWithMessage(c, appError.ErrUnauthorized, "failed to get session")
	}
	if session.IsRevoked {
		removeToken(c)
		return responses.ErrorWithMessage(c, appError.ErrUnauthorized, "session is revoked")
	}
	if session.UserEmail != claims.Email {
		removeToken(c)
		return responses.ErrorWithMessage(c, appError.ErrUnauthorized, "invalid email")
	}

	user, err := h.userUseCase.FindUserByEmail(claims.Email)
	if user == nil || err != nil {
		removeToken(c)
		return responses.Error(c, appError.ErrUnauthorized)
	}

	accessToken, accessClaims, err := h.tokenMaker.CreateToken(claims.ID, claims.Email, claims.Role, 10*time.Minute)
	if err != nil {
		return responses.Error(c, appError.ErrInternalServer)
	}
	return c.JSON(dto.ToRenewAccessTokenResponse(accessToken, accessClaims))
}

func (h *HttpSessionHandler) Logout(c *fiber.Ctx) error {
	tokenStr := c.Cookies("token")
	claims, err := h.tokenMaker.VerifyToken(tokenStr)
	if err != nil {
		removeToken(c)
		return responses.Error(c, appError.ErrUnauthorized)
	}

	if err := h.usecase.DeleteSession(claims.RegisteredClaims.ID); err != nil {
		return responses.Error(c, appError.ErrInternalServer)
	}
	removeToken(c)
	return responses.Message(c, fiber.StatusOK, "logged out successfully")
}

func (h *HttpSessionHandler) RevokeToken(c *fiber.Ctx) error {
	email := c.Params("email")
	if email == "" {
		return responses.Error(c, appError.ErrInvalidData)
	}

	if err := h.usecase.RevokeSession(email); err != nil {
		return responses.Error(c, err)
	}
	return responses.Message(c, fiber.StatusOK, "revoke successfully")
}
