package repository

import (
	"strings"

	entity "github.com/cosmart/internal/entities"
)

// Define structs to match the JSON response
type Availability struct {
	Status              string  `json:"status"`
	AvailableToBrowse   bool    `json:"available_to_browse"`
	AvailableToBorrow   bool    `json:"available_to_borrow"`
	AvailableToWaitlist bool    `json:"available_to_waitlist"`
	IsPrintdisabled     bool    `json:"is_printdisabled"`
	IsReadable          bool    `json:"is_readable"`
	IsLendable          bool    `json:"is_lendable"`
	IsPreviewable       bool    `json:"is_previewable"`
	Identifier          string  `json:"identifier"`
	OpenlibraryWork     string  `json:"openlibrary_work"`
	OpenlibraryEdition  string  `json:"openlibrary_edition"`
	LastLoanDate        *string `json:"last_loan_date"`
	NumWaitlist         *int    `json:"num_waitlist"`
	LastWaitlistDate    *string `json:"last_waitlist_date"`
	IsRestricted        bool    `json:"is_restricted"`
	IsBrowseable        bool    `json:"is_browseable"`
	Src                 string  `json:"__src__"`
}

type Author struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

type Work struct {
	Key               string       `json:"key"`
	Title             string       `json:"title"`
	EditionCount      int          `json:"edition_count"`
	CoverID           int          `json:"cover_id"`
	CoverEditionKey   string       `json:"cover_edition_key"`
	Subject           []string     `json:"subject"`
	LendingEdition    string       `json:"lending_edition"`
	LendingIdentifier string       `json:"lending_identifier"`
	Authors           []Author     `json:"authors"`
	FirstPublishYear  int          `json:"first_publish_year"`
	HasFulltext       bool         `json:"has_fulltext"`
	Availability      Availability `json:"availability"`
}

type WorksResponse struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	SubjectType string `json:"subject_type"`
	WorkCount   int    `json:"work_count"`
	Works       []Work `json:"works"`
}

type BooksResponse struct {
	Title   string                 `json:"title"`
	Authors []AuthorsBooksResponse `json:"authors"`
	Error   string                 `json:"error"`
}

type AuthorsBooksResponse struct {
	Author Author `json:"author"`
}

type DocsInformation struct {
	AuthorNames  []string `json:"author_name"`
	EditionCount int      `json:"edition_count"`
	Key          string   `json:"key"`
}

func convertWorksResponseToBooks(works WorksResponse) []entity.Book {
	books := make([]entity.Book, len(works.Works))
	for idx, work := range works.Works {
		books[idx] = entity.Book{
			Title:      work.Title,
			Authors:    convertWorksResponseAuthorsToAuthorSlice(work.Authors),
			EditionKey: work.Key,
		}
	}

	return books
}

func convertWorksResponseAuthorsToAuthorSlice(authors []Author) []string {
	authorList := make([]string, len(authors))
	for idx, val := range authors {
		authorList[idx] = val.Name
	}

	return authorList
}

func convertBooksResponseAuthorsToEntity(book BooksResponse, edition string) entity.BookInformation {
	return entity.BookInformation{
		Title:   book.Title,
		Edition: edition,
		Authors: convertBooksResponseAuthorsToEntityAuthorsSlice(book.Authors),
	}
}

func convertBooksResponseAuthorsToEntityAuthorsSlice(authors []AuthorsBooksResponse) []string {
	listAuthors := make([]string, len(authors))

	for idx, author := range authors {
		split := strings.Split(author.Author.Key, "/")
		if len(split) > 2 {
			listAuthors[idx] = split[2]
		}
	}

	return listAuthors

}
