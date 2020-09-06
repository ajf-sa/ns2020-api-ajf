package views

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/session"
)

var sessions = session.New(session.Config{Secure: true, Lookup: "cookie:api_c"})

func isLogin(c *fiber.Ctx) bool {
	store := sessions.Get(c)
	islogin := store.Get("login")
	switch islogin.(type) {
	case bool:
		if islogin.(bool) {
			return true
		} else {
			return false
		}
	default:
		return false
	}

}

func Index(c *fiber.Ctx) {

	_ = c.Render("index", fiber.Map{
		"Title": "Hi",
		"Login": isLogin(c),
	}, "layout")

}

func Login(c *fiber.Ctx) {
	store := sessions.Get(c)
	store.Set("login", true)
	defer store.Save()
	c.Redirect("/")
}

func Logout(c *fiber.Ctx) {
	store := sessions.Get(c)
	store.Delete("login")
	defer store.Save()
	c.Redirect("/")
}

func Register(c *fiber.Ctx) {
	c.Send("Register")
}
func About(c *fiber.Ctx) {
	c.Send("about")
}
