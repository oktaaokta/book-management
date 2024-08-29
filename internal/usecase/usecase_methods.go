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
	schedule := uc.repo.GetPickupSchedulesByEdition(edition)
	if len(schedule.Schedules) != 0 && !schedule.LastWaitlistDate.IsZero() {

	}

	booksResponse, err := uc.repo.GetBooksByParamFromRepo()
	if err != nil {
		return err
	}

	fmt.Println(booksResponse)

	return nil
}
