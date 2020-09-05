package route

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/session"
)

var sessions = session.New()

func Index(c *fiber.Ctx) {
	c.Send("index")
}

func Login(c *fiber.Ctx) {
	c.Send("Login")
}

func Logout(c *fiber.Ctx) {
	c.Send("Logout")
}

func Register(c *fiber.Ctx) {
	c.Send("Register")
}
func About(c *fiber.Ctx) {
	c.Send("about")
}
