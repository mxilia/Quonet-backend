package ratelimit

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	appError "github.com/mxilia/Quonet-backend/pkg/apperror"
	"github.com/mxilia/Quonet-backend/pkg/redisclient"
	"github.com/mxilia/Quonet-backend/pkg/responses"
)

type RateLimiter struct {
	store *redisclient.Client
}

func New(redisClient *redisclient.Client) *RateLimiter {
	return &RateLimiter{
		store: redisClient,
	}
}

type KeyFunc func(*fiber.Ctx) (string, bool)

func (rl *RateLimiter) Use(policy Policy, keyFn KeyFunc) fiber.Handler {
	cfg, ok := Policies[policy]
	if !ok {
		log.Println("rate limit policy not found: " + string(policy))
		return func(c *fiber.Ctx) error {
			return c.Next()
		}
	}

	return rl.limit(
		string(policy),
		cfg.Max,
		cfg.Window,
		keyFn,
	)
}

func (rl *RateLimiter) limit(name string, max int, window time.Duration, keyFn KeyFunc) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Method() == fiber.MethodOptions {
			return c.Next()
		}

		rawKey, ok := keyFn(c)
		if !ok {
			return c.Next()
		}

		var (
			key = name + ":" + rawKey
			ctx = c.UserContext()
		)

		count, err := rl.store.Incr(ctx, key, window)
		if err != nil {
			log.Println("Redis rate limiter failed:", err)
			return c.Next()
		}

		if count > int64(max) {
			return responses.ErrorWithMessage(c, appError.ErrLimitExceeded, "too many requests")
		}
		return c.Next()
	}
}
