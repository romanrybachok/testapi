package repository

import (
	"github.com/jmoiron/sqlx"
	api "github.com/romanrybachok/testapi"
)

type Authorization interface {
	CreateUser(user api.User) (int, error)
	GetUser(username, password string) (api.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
