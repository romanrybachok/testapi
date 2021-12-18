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
	Create(userId int, list api.TodoList) (int, error)
	GetAll(userId int) ([]api.TodoList, error)
	GetById(userId, listId int) (api.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input api.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, list api.TodoItem) (int, error)
	GetAll(userId, listId int) ([]api.TodoItem, error)
	GetById(userId, itemId int) (api.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, listId int, input api.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
