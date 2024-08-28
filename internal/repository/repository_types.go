package repository

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
	ISBN                *string `json:"isbn"` // Pointer to handle null values
	OCLC                *string `json:"oclc"`
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
	IACollection      []string     `json:"ia_collection"`
	LendingLibrary    bool         `json:"lendinglibrary"`
	PrintDisabled     bool         `json:"printdisabled"`
	LendingEdition    string       `json:"lending_edition"`
	LendingIdentifier string       `json:"lending_identifier"`
	Authors           []Author     `json:"authors"`
	FirstPublishYear  int          `json:"first_publish_year"`
	IA                string       `json:"ia"`
	PublicScan        bool         `json:"public_scan"`
	HasFulltext       bool         `json:"has_fulltext"`
	Availability      Availability `json:"availability"`
}

type OpenLibrarySubjectsResponse struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	SubjectType string `json:"subject_type"`
	WorkCount   int    `json:"work_count"`
	Works       []Work `json:"works"`
}
