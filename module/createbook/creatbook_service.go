package createbook

import (
	"errors"
)

type CreatbookService interface {
	CreateBook(books Books) error
}

type creatbookServiceImplement struct {
	repo CreatbookRepository
}

func newBookService(repo CreatbookRepository) CreatbookService {
	return &creatbookServiceImplement{repo: repo}
}

func (s *creatbookServiceImplement) CreateBook(books Books) error {
	// Bussines Logic function
	if books.book_id <= 0 {
		return errors.New("Not Fiald")
	}
	if err := s.repo.Save(books); err != nil {
		return err
	}
	return nil
}
