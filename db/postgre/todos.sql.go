// Code generated by sqlc. DO NOT EDIT.
// source: todos.sql

package db

import (
	"context"
	"database/sql"
)

const createTodo = `-- name: CreateTodo :one
INSERT INTO todos (name) VALUES ($1) RETURNING id, name, completed, created_at
`

func (q *Queries) CreateTodo(ctx context.Context, name string) (Todo, error) {
	row := q.queryRow(ctx, q.createTodoStmt, createTodo, name)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Completed,
		&i.CreatedAt,
	)
	return i, err
}

const deleteTodoById = `-- name: DeleteTodoById :exec
DELETE FROM todos WHERE id = $1
`

func (q *Queries) DeleteTodoById(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteTodoByIdStmt, deleteTodoById, id)
	return err
}

const getTodoById = `-- name: GetTodoById :one
SELECT id, name, completed, created_at FROM todos WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTodoById(ctx context.Context, id int64) (Todo, error) {
	row := q.queryRow(ctx, q.getTodoByIdStmt, getTodoById, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Completed,
		&i.CreatedAt,
	)
	return i, err
}

const listTodos = `-- name: ListTodos :many
SELECT id, name, completed, created_at FROM todos
ORDER BY id
`

func (q *Queries) ListTodos(ctx context.Context) ([]Todo, error) {
	rows, err := q.query(ctx, q.listTodosStmt, listTodos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Completed,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTodo = `-- name: UpdateTodo :one
UPDATE todos SET name = $2, completed = $3 WHERE id = $1 RETURNING id, name, completed, created_at
`

type UpdateTodoParams struct {
	ID        int64
	Name      string
	Completed sql.NullBool
}

func (q *Queries) UpdateTodo(ctx context.Context, arg UpdateTodoParams) (Todo, error) {
	row := q.queryRow(ctx, q.updateTodoStmt, updateTodo, arg.ID, arg.Name, arg.Completed)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Completed,
		&i.CreatedAt,
	)
	return i, err
}
