package service

import (
	api "github.com/romanrybachok/testapi"
	"github.com/romanrybachok/testapi/pkg/repository"
)

type Authorization interface {
	CreateUser(user api.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	CreateList(userId int, list api.TodoList) (int, error)
	GetAll(userId int) ([]api.TodoList, error)
	GetById(userId, id int) (api.TodoList, error)
	Delete(userId, id int) error
	Update(userId, listId int, input api.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item api.TodoItem) (int, error)
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
