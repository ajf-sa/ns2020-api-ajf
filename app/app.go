package app

import (
	"api-ajf/app/route"

	"github.com/gofiber/fiber"
)

func New() *AppX {
	return &AppX{app: fiber.New()}
}

func (app *AppX) Route() {
	app.app.Get("/login", route.Login)
	app.app.Get("/logout", route.Logout)
	app.app.Get("/register", route.Register)
	app.app.Get("/about", route.About)
	app.app.Get("/", route.Index)
}

func (app *AppX) Listen() error {
	return app.app.Listen("127.0.0.1:3000")
}

type AppX struct {
	app *fiber.App
}
