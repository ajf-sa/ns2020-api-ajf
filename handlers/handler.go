package handlers

import (
	db "api-ajf/db/postgre"
)

type Handlers struct {
	Repo *db.Repo
}

func NewHandlers(repo *db.Repo) *Handlers {
	return &Handlers{Repo: repo}
}
