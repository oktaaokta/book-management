package usecase

import (
	entity "github.com/cosmart/internal/entities"
	repo "github.com/cosmart/internal/repository"
)

type repositoryInterface interface {
	GetBooksByParamFromRepo() (repo.BooksResponse, error)
	GetBooksBySubjectFromRepo(subject string) ([]entity.Book, error)
}
