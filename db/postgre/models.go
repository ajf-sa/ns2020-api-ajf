// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
)

type Todo struct {
	ID        int32          `json:"id"`
	Name      sql.NullString `json:"name"`
	Done      sql.NullBool   `json:"done"`
	CreatedAt sql.NullTime   `json:"created_at"`
}
