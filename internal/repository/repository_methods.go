package repository

import (
	"encoding/json"
	"log"
	"net/http"

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

func (r *Repository) GetBooksByParamFromRepo() (BooksResponse, error) {
	response, err := http.Get(openLibraryAPI + "/search.json?q=the+lord+of+the+rings&fields=key,title,author_name,editions")
	if err != nil {
		return BooksResponse{}, err
	}
	defer response.Body.Close()

	var booksResponse BooksResponse
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&booksResponse); err != nil {
		log.Fatalf("Failed to decode JSON response: %v", err)
	}

	return booksResponse, nil
}

// GetPickupSchedulesByEdition retrieves book pickup schedule by given edition.
func (r *Repository) GetPickupSchedulesByEdition(edition string) infrastructure.ScheduleInformation {
	return r.ps.GetPickupSchedules(edition)
}

func (r *Repository) SetPickupSchedulesByEdition(edition string) {
	r.ps.SetPickupSchedules(edition)
}
