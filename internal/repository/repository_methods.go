package repository

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	entity "github.com/cosmart/internal/entities"
	"github.com/cosmart/internal/infrastructure"
)

const (
	openLibraryAPI = "https://openlibrary.org"
)

func (r *Repository) GetBooksBySubjectFromRepo(subject string) ([]entity.Book, error) {
	url := openLibraryAPI + "/subjects/" + subject + ".json"
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var workResponse WorksResponse
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&workResponse); err != nil {
		log.Fatalf("Failed to decode JSON response: %v", err)
	}

	return convertWorksResponseToBooks(workResponse), nil
}

func (r *Repository) GetWorkByEdition(edition string) (entity.BookInformation, error) {
	response, err := http.Get(openLibraryAPI + "/works/" + edition + ".json")
	if err != nil {
		return entity.BookInformation{}, err
	}
	defer response.Body.Close()

	var booksResponse BooksResponse
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&booksResponse); err != nil {
		log.Fatalf("Failed to decode JSON response: %v", err)
	}

	if booksResponse.Error != "" {
		return entity.BookInformation{}, errors.New(booksResponse.Error)
	}

	return convertBooksResponseAuthorsToEntity(booksResponse, edition), nil
}

// GetPickupSchedulesByEdition retrieves book pickup schedule by given edition.
func (r *Repository) GetPickupSchedulesByEdition(edition string) infrastructure.ScheduleInformation {
	return r.ps.GetPickupSchedules(edition)
}

func (r *Repository) SetPickupSchedulesByEdition(edition string, pickupDate, returnDate time.Time, bookInfo entity.BookInformation) {
	r.ps.SetPickupSchedules(edition, pickupDate, returnDate, bookInfo)
}
