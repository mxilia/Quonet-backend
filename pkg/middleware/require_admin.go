package middleware

import (
	"github.com/gofiber/fiber/v2"
	appError "github.com/mxilia/Conflux-backend/pkg/apperror"
	"github.com/mxilia/Conflux-backend/pkg/responses"
)

func RequireAdmin() fiber.Handler {

	return func(c *fiber.Ctx) error {
		role, ok := c.Locals("role").(string)
		if !ok || (role != "admin" && role != "owner") {
			return responses.ErrorWithMessage(c, appError.ErrForbidden, "forbidden route")
		}
		return c.Next()
	}
}
