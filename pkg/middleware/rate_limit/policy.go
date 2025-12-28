package ratelimit

import "time"

type Policy string

const (
	UserRead   Policy = "user-read"
	UserWrite  Policy = "user-write"
	AuthLogin  Policy = "auth-login"
	AuthLogout Policy = "auth-logout"
	PublicRead Policy = "public-read"
	AdminWrite Policy = "admin-write"
)

type Config struct {
	Max    int
	Window time.Duration
}

var Policies = map[Policy]Config{
	UserRead: {
		Max:    500,
		Window: time.Minute,
	},
	UserWrite: {
		Max:    30,
		Window: time.Minute,
	},
	AuthLogin: {
		Max:    5,
		Window: time.Minute,
	},
	AuthLogout: {
		Max:    5,
		Window: time.Minute,
	},
	PublicRead: {
		Max:    500,
		Window: time.Minute,
	},
	AdminWrite: {
		Max:    10,
		Window: time.Minute,
	},
}
