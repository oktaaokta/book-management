package usecase

import "github.com/cosmart/internal/repository"

type repositoryInterface interface {
	GetListOfBooksFromRepo() (repository.OpenLibrarySubjectsResponse, error)
}
