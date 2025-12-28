package ratelimit

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UserKey(c *fiber.Ctx) (string, bool) {
	id, ok := c.Locals("user_id").(uuid.UUID)
	if !ok {
		return "", false
	}
	return "user:" + id.String(), true
}

func IPKey(c *fiber.Ctx) (string, bool) {
	return "ip:" + c.IP(), true
}

func UserOrIPKey(c *fiber.Ctx) (string, bool) {
	if key, ok := UserKey(c); ok {
		return key, true
	}
	return IPKey(c)
}
