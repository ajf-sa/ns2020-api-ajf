package app

import (
	"github.com/gofiber/fiber"
)

func New() *fiber.App {
	return fiber.New()
}

func Route(app *fiber.App) {
	app.Get("/", index)
	app.Get("/about", about)
}

func index(c *fiber.Ctx) {
	c.Send("index")
}

func about(c *fiber.Ctx) {
	c.Send("about")
}

// type App struct {
// 	Fiber *fiber.App
// }

// func (a *App) index(c *fiber.Ctx) {

// 	c.Send("index")

// }
