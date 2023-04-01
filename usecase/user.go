package usecase

import (
	"fmt"

	"github.com/kuma0328/go_training/domain/entity"
	"github.com/kuma0328/go_training/domain/repositry"
)

var _UserUsacaseImpl = &UserUsecase{}

type UsecaseImpl interface {
	CreateUser(*entity.User) error
	UpdateUser(*entity.User) error
}

type UserUsecase struct {
	repo repositry.UserRepositryImpl
}

func NewUserUsecase(repo repositry.UserRepositryImpl) *UserUsecase {
	return &UserUsecase {
		repo : repo,
	}
}

func (uu *UserUsecase) CreateUser(user *entity.User) error {
	if user.ID == 0 {
		fmt.Errorf("User id must not be empty")
	}
	err := uu.CreateUser(user)
	return err
}