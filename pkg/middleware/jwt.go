package middleware

import (
	"github.com/gofiber/fiber/v2"
	appError "github.com/mxilia/Conflux-backend/pkg/apperror"
	"github.com/mxilia/Conflux-backend/pkg/config"
	"github.com/mxilia/Conflux-backend/pkg/responses"
	"github.com/mxilia/Conflux-backend/pkg/token"
)

func JWTMiddleware(cfg *config.Config) fiber.Handler {
	tokenMaker := token.NewJWTMaker(cfg)

	return func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")
		if auth == "" {
			return responses.ErrorWithMessage(c, appError.ErrUnauthorized, "missing token")
		}

		claims, err := tokenMaker.VerifyToken(auth[len("Bearer "):])
		if err != nil {
			return responses.ErrorWithMessage(c, appError.ErrUnauthorized, err.Error())
		}

		c.Locals("user_id", claims.ID)
		c.Locals("role", claims.Role)

		return c.Next()
	}
}
