package usecase

import (
	"errors"
	"fmt"
	"strings"
	"time"

	entity "github.com/cosmart/internal/entities"
	"github.com/cosmart/internal/repository"
)

func (uc *Usecase) GetListOfBooks(subject string) ([]entity.Book, error) {
	resp, err := uc.repo.GetBooksBySubjectFromRepo(subject)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (uc *Usecase) SubmitBookPickupSchedule(edition string, pickupDate, returnDate time.Time) error {
	schedule := uc.repo.GetPickupSchedulesByEdition(edition)
	if len(schedule.Schedules) > 0 && schedule.LastWaitlistDate.After(pickupDate) {
		message := fmt.Sprintf("book is not available during the pickup time. The book will be available at: %v", schedule.LastWaitlistDate)
		return errors.New(message)
	}

	booksResponse, err := uc.repo.GetWorkByEdition(edition)
	if err != nil {
		return err
	}

	if booksResponse.Error != "" {
		return errors.New(booksResponse.Error)
	}

	uc.repo.SetPickupSchedulesByEdition(edition, pickupDate, returnDate, entity.BookInformation{
		Title:   booksResponse.Title,
		Authors: convertBooksResponseAuthorsToEntity(booksResponse.Authors),
	})

	return nil
}

func convertBooksResponseAuthorsToEntity(authors []repository.AuthorsBooksResponse) []string {
	listAuthors := make([]string, len(authors))

	for idx, author := range authors {
		split := strings.Split(author.Author.Key, "/")
		if len(split) > 2 {
			listAuthors[idx] = split[2]
		}
	}

	return listAuthors
}
