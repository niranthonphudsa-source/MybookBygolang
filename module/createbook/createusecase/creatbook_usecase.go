package createusecase

import (
	"mylibary/module/createbook/createrepository"
	"mylibary/module/entities"
	"mylibary/module/getbook/repository"

	"github.com/gofiber/fiber/v2"
)

type CreateBookUsecase interface {
	CreateBook(c *fiber.Ctx) (*entities.Books, error)
}

type createBookUsecaseRepo struct {
	repo createrepository.CreateBookRepository
}

func NewCreateBookUsecase(r createrepository.CreateBookRepository) *createBookUsecaseRepo {
	return &createBookUsecaseRepo{repo: r}
}

type bookUsecase struct {
	repo repository.BookRepository
}

func NewBookUsecase(r repository.BookRepository) *bookUsecase {
	return &bookUsecase{repo: r}
}

func (u *createBookUsecaseRepo) CreateBook(c *fiber.Ctx) (*entities.Books, error) {
	if err := u.repo.CreateBookRepository(c); err != nil {
		return nil, err
	}
	return nil, nil
}
