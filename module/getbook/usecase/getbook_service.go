package usecase

import (
	"mylibary/module/entities"
	"mylibary/module/getbook/repository"
)

type BookUsecase interface {
	GetBookAllRepoImpl() ([]entities.Books, error)
}

type BookIdUsecase interface {
	GetBookIdRepoImpl(book_id string) (*entities.Books, error)
}

type bookUsecase struct {
	repo repository.BookRepository
}

type bookIdUsecase struct {
	repoId repository.BookIdRepository
}

func NewBookUsecase(r repository.BookRepository) *bookUsecase {
	return &bookUsecase{repo: r}
}

func NewBookIdUsecase(r repository.BookIdRepository) *bookIdUsecase {
	return &bookIdUsecase{repoId: r}
}

func (u *bookUsecase) GetBookAllRepoImpl() ([]entities.Books, error) {
	books, err := u.repo.GetBookAllRepo()
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (u *bookIdUsecase) GetBookIdRepoImpl(book_id string) (*entities.Books, error) {
	return u.repoId.GetBookIdRepo(book_id)
}
