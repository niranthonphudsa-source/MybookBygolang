package bookusecase

import (
	"fmt"
	bookrepository "mylibary/book_module/book_repository"
	"mylibary/book_module/entities"
)

type BookUsecase interface {
	GetBookAllRepoImpl() ([]entities.Books, error)
	GetBookIdRepoImpl(book_id int) (*entities.Books, error)
	CreateBookRepoImpl(bookNew *entities.Books) error
	DeleteBookIdImpl(book_id int) error
	UpdateBookIdRepoImpl(book_id int, bookUpdate *entities.Books) (*entities.Books, error)
}

type bookUsecase struct {
	repo bookrepository.BooksRepository
}

func NewBookUsecase(r bookrepository.BooksRepository) *bookUsecase {
	return &bookUsecase{repo: r}
}

func (bookrepo *bookUsecase) GetBookAllRepoImpl() ([]entities.Books, error) {
	books, err := bookrepo.repo.GetBooksAll()
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (bookrepo *bookUsecase) GetBookIdRepoImpl(book_id int) (*entities.Books, error) {
	return bookrepo.repo.GetBookId(book_id)
}

func (bookrepo *bookUsecase) CreateBookRepoImpl(bookNew *entities.Books) error {
	return bookrepo.repo.CreateBook(bookNew)
}

func (bookrepo *bookUsecase) DeleteBookIdImpl(book_id int) error {
	return bookrepo.repo.DeleteBookId(book_id)
}

func (bookrepo *bookUsecase) UpdateBookIdRepoImpl(book_id int, bookUpdate *entities.Books) (*entities.Books, error) {
	fmt.Print("Usecase", book_id)
	return bookrepo.repo.UpdateBookId(book_id, bookUpdate)
}
