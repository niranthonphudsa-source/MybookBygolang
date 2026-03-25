package loginusecase

import (
	"mylibary/login_module/entities"
	loginrepository "mylibary/login_module/login_repository"
)

type Loginusecase interface {
	LoginUsecase(data *entities.UserLogin) error
}

type loginUsecase struct {
	repo loginrepository.UserLoginRepository
}

func NewLoginUsecase(r loginrepository.UserLoginRepository) *loginUsecase {
	return &loginUsecase{repo: r}
}

func (u *loginUsecase) LoginUsecase(data *entities.UserLogin) error {
	return u.repo.LoginRepo(data)
}
