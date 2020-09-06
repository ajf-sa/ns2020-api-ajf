package main

import (
	db "api-ajf/db/postgre"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
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
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.Get("/", func(ctx *fiber.Ctx) {
		ctx.Send("hello world")
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

func (h *Handlers) GetTodos(ctx *fiber.Ctx) {
	todos, err := h.Repo.ListTodos(ctx.Context())
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).Send(err.Error())
		return
	}

	result := make([]interface{}, len(todos))
	for i, todo := range todos {
		result[i] = mapTodo(todo)
	}

	if err := ctx.Status(fiber.StatusOK).JSON(result); err != nil {
		ctx.Status(fiber.StatusInternalServerError).Send(err.Error())
		return
	}
}

func (h *Handlers) GetTodo(ctx *fiber.Ctx) {
	paramsId := ctx.Params("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
		return
	}

	todo, err := h.Repo.GetTodoById(ctx.Context(), int64(id))
	if err != nil {
		ctx.Status(fiber.StatusNotFound)
		return
	}

	if err := ctx.Status(fiber.StatusOK).JSON(mapTodo(todo)); err != nil {
		ctx.Status(fiber.StatusInternalServerError).Send(err.Error())
		return
	}
}
func (h *Handlers) CreateTodo(ctx *fiber.Ctx) {
	type request struct {
		Name string `json:"name"`
	}

	var body request

	err := ctx.BodyParser(&body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
		return
	}

	if len(body.Name) <= 2 {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "name not long enough",
		})
		return
	}

	todo, err := h.Repo.CreateTodo(ctx.Context(), body.Name)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).Send(err.Error())
		return
	}

	if err := ctx.Status(fiber.StatusCreated).JSON(mapTodo(todo)); err != nil {
		ctx.Status(fiber.StatusInternalServerError).Send(err.Error())
		return
	}
}

func (h *Handlers) UpdateTodo(ctx *fiber.Ctx) {
	type request struct {
		Name      *string `json:"name"`
		Completed *bool   `json:"completed"`
	}

	paramsId := ctx.Params("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
		return
	}

	var body request
	err = ctx.BodyParser(&body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse body",
		})
		return
	}

	todo, err := h.Repo.GetTodoById(ctx.Context(), int64(id))
	if err != nil {
		ctx.Status(fiber.StatusNotFound)
		return
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
		ctx.SendStatus(fiber.StatusNotFound)
		return
	}

	if err := ctx.Status(fiber.StatusOK).JSON(mapTodo(todo)); err != nil {
		ctx.Status(fiber.StatusInternalServerError).Send(err.Error())
		return
	}
}

func (h *Handlers) DeleteTodo(ctx *fiber.Ctx) {
	paramsId := ctx.Params("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
		return
	}

	_, err = h.Repo.GetTodoById(ctx.Context(), int64(id))
	if err != nil {
		ctx.Status(fiber.StatusNotFound)
		return
	}

	err = h.Repo.DeleteTodoById(ctx.Context(), int64(id))
	if err != nil {
		ctx.SendStatus(fiber.StatusNotFound)
		return
	}

	ctx.SendStatus(fiber.StatusNoContent)
}
