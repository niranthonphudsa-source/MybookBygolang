package usecase

import (
	"mylibary/module/entities"
	"mylibary/module/getbook/repository"
)

type BookUsecase interface {
	GetBookAllRepoImpl() ([]entities.Books, error)
}

type bookUsecase struct {
	repo repository.BookRepository
}

func NewBookUsecase(r repository.BookRepository) *bookUsecase {
	return &bookUsecase{repo: r}
}
func (u *bookUsecase) GetBookAllRepoImpl() ([]entities.Books, error) {
	books, err := u.repo.GetBookAllRepo()
	if err != nil {
		return nil, err
	}

	return books, nil
}
