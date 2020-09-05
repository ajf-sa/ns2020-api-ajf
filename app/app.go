package app

import (
	"api-ajf/app/views"

	"github.com/gofiber/fiber"
	"github.com/gofiber/template/html"
)

func New() *AppX {
	engine := html.New("./templates", ".html")
	app := fiber.New(&fiber.Settings{
		Views:     engine,
		BodyLimit: (15 * 1024 * 1024 * 1024),
	})
	app.Static("/image", "./image")
	app.Static("/static", "./static")

	return &AppX{app: app}
}

func (app *AppX) Route() {
	app.app.Get("/login", views.Login)
	app.app.Get("/logout", views.Logout)
	app.app.Get("/register", views.Register)
	app.app.Get("/about", views.About)
	app.app.Get("/", views.Index)
}

func (app *AppX) Listen() error {
	return app.app.Listen("127.0.0.1:3000")
}

type AppX struct {
	app *fiber.App
}
