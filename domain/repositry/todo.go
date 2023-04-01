package repositry

import "github.com/kuma0328/go_training/domain/entity"

type TodoRepositryImpl interface {
	CreateTodo(*entity.Todo) error
	UpdateTodo(*entity.Todo) error
	DeleteTodo(id int) error
	GetById(id int) (*entity.Todo, error)
}