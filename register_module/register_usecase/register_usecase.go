package registerusecase

import (
	"mylibary/register_module/entities"
	registerrepository "mylibary/register_module/register_repository"
)

type RegisterUsecase interface {
	RegisterUsecase(data *entities.Users) error
}

type registerUsecase struct {
	repo registerrepository.RegisterRepository
}

func NewRegisterUsecaseImplrepo(r registerrepository.RegisterRepository) *registerUsecase {
	return &registerUsecase{repo: r}
}

func (r *registerUsecase) RegisterUsecase(data *entities.Users) error {
	return r.repo.RegisterRepository(data)
}
