package repository

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func (r *Repository) GetListOfBooksFromRepo() (OpenLibrarySubjectsResponse, error) {
	response, err := http.Get("https://openlibrary.org/subjects/science_fiction.json")
	if err != nil {
		return OpenLibrarySubjectsResponse{}, err
	}
	defer response.Body.Close()

	var subjectResponse OpenLibrarySubjectsResponse
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&subjectResponse); err != nil {
		log.Fatalf("Failed to decode JSON response: %v", err)
	}

	return subjectResponse, errors.New("Hi")
}
