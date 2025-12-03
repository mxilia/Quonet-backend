package middleware

import (
	"github.com/gofiber/fiber/v2"
	appError "github.com/mxilia/Quonet-backend/pkg/apperror"
	"github.com/mxilia/Quonet-backend/pkg/responses"
)

func RequireUser() fiber.Handler {

	return func(c *fiber.Ctx) error {
		role, ok := c.Locals("role").(string)
		if ok && (role == "admin" || role == "owner") {
			return c.Next()
		}

		userID, ok := c.Locals("user_id").(string)
		if !ok || userID == "" {
			return responses.ErrorWithMessage(c, appError.ErrUnauthorized, "missing id")
		}

		targetAuthorID := c.Params("id")
		if targetAuthorID != userID {
			return responses.ErrorWithMessage(c, appError.ErrForbidden, "forbidden route")
		}
		return c.Next()
	}
}
