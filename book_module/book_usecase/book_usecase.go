package bookusecase

import (
	bookrepository "mylibary/book_module/book_repository"
	"mylibary/book_module/entities"
)

type BookUsecase interface {
	GetBookAllRepoImpl() ([]entities.Books, error)
	GetBookIdRepoImpl(book_id int) (*entities.Books, error)
	CreateBookRepoImpl(bookNew *entities.Books) error
	DeleteBookIdImpl(book_id int) error
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

func (bookrepo *bookUsecase) CreateBookRepoImpl(book_id int) error
