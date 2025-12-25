package app

import "log"

func Start() {
	db, supabase, cfg, err := setupDependencies("dev")
	if err != nil {
		log.Fatalf("Failed to setup dependencies: %v", err)
	}

	app, err := setupRestServer(db, supabase, cfg)
	if err != nil {
		log.Fatalf("Failed to setup Rest server: %v", err)
	}

	app.Listen(":8000")
}
