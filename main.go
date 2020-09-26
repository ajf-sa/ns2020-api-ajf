package main

import (
	db "api-ajf/db/postgre"
	"api-ajf/handlers"
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	_ "github.com/lib/pq"
)

func conn() *sql.DB {
	dbd, err := sql.Open("postgres", fmt.Sprintf("dbname=%s password=secret user=admin sslmode=disable", "simple_api"))
	if err != nil {
		panic(err)
	}
	return dbd
}

func main() {

	repo := db.NewRepo(conn())

	app := fiber.New()

	handlers := handlers.NewHandlers(repo)
	SetApiV1(app, handlers)

	app.Use(logger.New())
	app.Use(recover.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("hello world"))
	})

	err := app.Listen("127.0.0.1:3000")
	if err != nil {
		panic(err)
	}
}

func SetApiV1(app *fiber.App, handlers *handlers.Handlers) {
	v1 := app.Group("/v1")
	SetTodoRoutes(v1, handlers)
}

func SetTodoRoutes(grp fiber.Router, handlers *handlers.Handlers) {
	todosRoutes := grp.Group("/todos")
	todosRoutes.Get("/", handlers.GetTodos)
	todosRoutes.Post("/", handlers.CreateTodo)
	todosRoutes.Get("/:id", handlers.GetTodo)
	todosRoutes.Delete("/:id", handlers.DeleteTodo)
	todosRoutes.Patch("/:id", handlers.UpdateTodo)

}
