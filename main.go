package main

import (
	db "api-ajf/db/postgre"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	_ "github.com/lib/pq"
)

func mapTodo(todo db.Todo) interface{} {

	return struct {
		ID        int64  `json:"id"`
		Name      string `json:"name"`
		Completed bool   `json:"completed"`
	}{
		ID:        todo.ID,
		Name:      todo.Name,
		Completed: todo.Completed.Bool,
	}
}

type Todo struct {
	Id        int    `json:id`
	Name      string `json:name`
	Completed bool   `json:completed`
}

var todos = []*Todo{
	{Id: 1, Name: "make something", Completed: false},
	{Id: 2, Name: "wash Car", Completed: false},
}

type Handlers struct {
	Repo *db.Repo
}

func NewHandlers(repo *db.Repo) *Handlers {
	return &Handlers{Repo: repo}
}

func main() {

	dbd, err := sql.Open("postgres", fmt.Sprintf("dbname=%s password=secret user=admin sslmode=disable", "simple_api"))
	if err != nil {
		panic(err)
	}
	repo := db.NewRepo(dbd)

	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("hello world"))
	})

	handlers := NewHandlers(repo)
	SetApiV1(app, handlers)

	err = app.Listen("127.0.0.1:3000")
	if err != nil {
		panic(err)
	}
}

func SetApiV1(app *fiber.App, handlers *Handlers) {
	v1 := app.Group("/v1")
	SetTodoRoutes(v1, handlers)
}

func SetTodoRoutes(grp fiber.Router, handlers *Handlers) {
	todosRoutes := grp.Group("/todos")
	todosRoutes.Get("/", handlers.GetTodos)
	todosRoutes.Post("/", handlers.CreateTodo)
	todosRoutes.Get("/:id", handlers.GetTodo)
	todosRoutes.Delete("/:id", handlers.DeleteTodo)
	todosRoutes.Patch("/:id", handlers.UpdateTodo)

}

func (h *Handlers) GetTodos(ctx *fiber.Ctx) error {
	todos, err := h.Repo.ListTodos(ctx.Context())
	if err != nil {

		return ctx.Status(fiber.StatusInternalServerError).Send([]byte(err.Error()))
	}

	result := make([]interface{}, len(todos))
	for i, todo := range todos {
		result[i] = mapTodo(todo)
	}

	if err := ctx.Status(fiber.StatusOK).JSON(result); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).Send([]byte(err.Error()))
	}
	return nil
}

func (h *Handlers) GetTodo(ctx *fiber.Ctx) error {
	paramsId := ctx.Params("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
	}

	todo, err := h.Repo.GetTodoById(ctx.Context(), int64(id))
	if err != nil {

		ctx.Status(fiber.StatusNotFound)
		return nil
	}

	if err := ctx.Status(fiber.StatusOK).JSON(mapTodo(todo)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).Send([]byte(err.Error()))
	}
	return nil
}
func (h *Handlers) CreateTodo(ctx *fiber.Ctx) error {
	type request struct {
		Name string `json:"name"`
	}

	var body request

	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})

	}

	if len(body.Name) <= 2 {

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "name not long enough",
		})
	}

	todo, err := h.Repo.CreateTodo(ctx.Context(), body.Name)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).Send([]byte(err.Error()))
	}

	if err := ctx.Status(fiber.StatusCreated).JSON(mapTodo(todo)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).Send([]byte(err.Error()))
	}
	return nil
}

func (h *Handlers) UpdateTodo(ctx *fiber.Ctx) error {
	type request struct {
		Name      *string `json:"name"`
		Completed *bool   `json:"completed"`
	}

	paramsId := ctx.Params("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})

	}

	var body request
	err = ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse body",
		})

	}

	todo, err := h.Repo.GetTodoById(ctx.Context(), int64(id))
	if err != nil {
		ctx.Status(fiber.StatusNotFound)
		return nil

	}

	if body.Name != nil {
		todo.Name = *body.Name
	}

	if body.Completed != nil {
		todo.Completed = sql.NullBool{
			Bool:  *body.Completed,
			Valid: true,
		}
	}

	todo, err = h.Repo.UpdateTodo(ctx.Context(), db.UpdateTodoParams{
		ID:        int64(id),
		Name:      todo.Name,
		Completed: todo.Completed,
	})
	if err != nil {
		return ctx.SendStatus(fiber.StatusNotFound)

	}

	if err := ctx.Status(fiber.StatusOK).JSON(mapTodo(todo)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).Send([]byte(err.Error()))
	}
	return nil
}

func (h *Handlers) DeleteTodo(ctx *fiber.Ctx) error {
	paramsId := ctx.Params("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})

	}

	_, err = h.Repo.GetTodoById(ctx.Context(), int64(id))
	if err != nil {
		ctx.Status(fiber.StatusNotFound)
		return nil

	}

	err = h.Repo.DeleteTodoById(ctx.Context(), int64(id))
	if err != nil {
		return ctx.SendStatus(fiber.StatusNotFound)

	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
