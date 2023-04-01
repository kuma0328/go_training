package repositry

import "github.com/kuma0328/go_training/domain/entity"

type UserRepositryImpl interface {
	CreateUser(*entity.User) error
	UpdateUser(*entity.User) error
}