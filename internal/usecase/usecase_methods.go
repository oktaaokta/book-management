package usecase

import (
	"fmt"

	entity "github.com/cosmart/internal/entities"
)

func (uc *Usecase) GetListOfBooks(subject string) ([]entity.Book, error) {
	resp, err := uc.repo.GetBooksBySubjectFromRepo(subject)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (uc *Usecase) SubmitBookPickupSchedule(edition string) error {
	booksResponse, err := uc.repo.GetBooksByParamFromRepo()
	if err != nil {
		return err
	}

	fmt.Println(booksResponse)

	return nil
}
