package usecase

import (
	"fmt"

	"github.com/kuma0328/go_training/domain/entity"
	"github.com/kuma0328/go_training/domain/repositry"
)

var _TodoUsecaseImpl = &TodoUsecase{}

var TitleEmptyError = "Todo title must not be empty"
var IdEmptyError = "Todo id must not be empty"

type TodoUsecaseImpl interface {
	CreateTodo(*entity.Todo) error
	UpdateTodo(*entity.Todo) error
	DeleteTodo(id int) error
	GetById(id int) (*entity.Todo, error)
}

type TodoUsecase struct {
	repo repositry.TodoRepositryImpl
}

func NewTodoUsecase(repo repositry.TodoRepositryImpl) *TodoUsecase {
	return &TodoUsecase {
		repo : repo,
	}
}

func (tu *TodoUsecase) CreateTodo(todo *entity.Todo) error {
	if todo.Title == "" {
		return fmt.Errorf(TitleEmptyError)
	}
	err := tu.repo.CreateTodo(todo)
	return err
}

func (tu *TodoUsecase) UpdateTodo(todo *entity.Todo) error {
	if todo.Title == "" {
		return fmt.Errorf(TitleEmptyError)
	}
	err := tu.UpdateTodo(todo)
	return err
}

func (tu *TodoUsecase) DeleteTodo(id int) error {
	if id == 0 {
		return fmt.Errorf(IdEmptyError)
	}
	err := tu.DeleteTodo(id)
	return err
}

func (tu *TodoUsecase) GetById(id int) (*entity.Todo, error) {
	if id == 0 {
		return nil, fmt.Errorf(IdEmptyError)
	}
	todo, err := tu.GetById(id)
	return todo, err
}