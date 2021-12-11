package service

import (
	api "github.com/romanrybachok/testapi"
	"github.com/romanrybachok/testapi/pkg/repository"
)

type Authorization interface {
	CreateUser(user api.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
