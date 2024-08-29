package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (hd *Handler) GetBooksList(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Success",
	}

	query := r.URL.Query()
	subject := query.Get("subject")
	if subject == "" {
		response.Message = "Missing parameter subject"
		writeResponse(w, response, http.StatusBadRequest)
		return
	}

	books, err := hd.uc.GetListOfBooks(subject)
	if err != nil {
		log.Println("got error when getting list of books: ", err)
		response.Message = "Failed"

		writeResponse(w, response, http.StatusInternalServerError)
		return
	}

	response.Books = make([]BookInformation, len(books))
	for idx, book := range books {
		response.Books[idx] = BookInformation{
			Title:      book.Title,
			Authors:    book.Authors,
			EditionKey: book.EditionKey,
		}
	}

	writeResponse(w, response, http.StatusOK)
}

func (hd *Handler) SubmitBookPickupSchedule(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		writeResponse(w, Response{Message: "body parameter is not valid"}, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var requestData PickupScheduleReq

	// Unmarshal the JSON body into the struct
	if err := json.Unmarshal(body, &requestData); err != nil {
		writeResponse(w, Response{Message: "error parsing request body"}, http.StatusBadRequest)
		return
	}

	// validate input
	if requestData.Edition == "" || requestData.PickupDate.IsZero() || requestData.ReturnDate.IsZero() {
		writeResponse(w, Response{Message: "missing parameter in body"}, http.StatusBadRequest)
		return
	}

	err = hd.uc.SubmitBookPickupSchedule(requestData.Edition)
	if err != nil {
		response := Response{
			Message: "Book pickup schedule failed due to: " + err.Error(),
		}
		log.Println("error when submitting the pickup schedule: ", err)

		writeResponse(w, response, http.StatusBadRequest)
		return
	}

	response := Response{
		Message: "Pickup schedule submitted.",
	}

	writeResponse(w, response, http.StatusOK)
}

func writeResponse(w http.ResponseWriter, resp Response, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error when writing response: ", err)
	}
}
