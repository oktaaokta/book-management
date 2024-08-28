package handler

import "github.com/cosmart/internal/repository"

type usecaseInterface interface {
	GetListOfBooks() (repository.OpenLibrarySubjectsResponse, error)
}
