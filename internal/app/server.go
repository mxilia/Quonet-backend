package app

import "log"

func Start() {
	db, storage, _, rateLimiter, cfg, err := setupDependencies("dev")
	if err != nil {
		log.Fatalf("Failed to setup dependencies: %v", err)
	}

	app, err := setupRestServer(db, storage, rateLimiter, cfg)
	if err != nil {
		log.Fatalf("Failed to setup Rest server: %v", err)
	}

	app.Listen(":8000")
}
