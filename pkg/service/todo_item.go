package service

import (
	api "github.com/romanrybachok/testapi"
	"github.com/romanrybachok/testapi/pkg/repository"
)

type TodoItemService struct {
	itemRepo repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{itemRepo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId, listId int, item api.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		// list does not exists or does not belongs to user
		return 0, err
	}

	return s.itemRepo.Create(listId, item)
}

func (s *TodoItemService) GetAll(userId, listId int) ([]api.TodoItem, error) {
	return s.itemRepo.GetAll(userId, listId)
}
