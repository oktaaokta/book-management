package usecase

import "github.com/cosmart/internal/repository"

func (uc *Usecase) GetListOfBooks() (repository.OpenLibrarySubjectsResponse, error) {
	resp, err := uc.repo.GetListOfBooksFromRepo()
	if err != nil {
		return resp, err
	}
	return resp, nil
}
