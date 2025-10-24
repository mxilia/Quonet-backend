package app

func Start() {
	app := setUpRestServer()
	app.Listen(":8000")
}
