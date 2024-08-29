package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (hd *Handler) GetBooksList(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Status:  http.StatusOK,
		Message: "Success",
	}

	query := r.URL.Query()
	subject := query.Get("subject")
	if subject == "" {
		response.Status = http.StatusBadRequest
		response.Message = "Missing parameter subject"
		writeResponse(w, response)
		return
	}

	books, err := hd.uc.GetListOfBooks(subject)
	if err != nil {
		log.Println("got error when getting list of books: ", err)
		response.Status = http.StatusInternalServerError
		response.Message = "Failed"

		writeResponse(w, response)
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

	writeResponse(w, response)
}

func (hd *Handler) SubmitBookPickupSchedule(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var requestData PickupScheduleReq

	// Unmarshal the JSON body into the struct
	if err := json.Unmarshal(body, &requestData); err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	err = hd.uc.SubmitBookPickupSchedule(requestData.Edition)
	if err != nil {
		response := Response{
			Message: "Book pickup schedule failed due to: " + err.Error(),
			Status:  http.StatusInternalServerError,
		}
		log.Println("error when submitting the pickup schedule: ", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.Status)

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println("Error when writing response: ", err)
		}
		return
	}

	response := Response{
		Message: "Pickup schedule submitted.",
		Status:  http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Status)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error when writing response: ", err)
	}
}

func writeResponse(w http.ResponseWriter, resp Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Status)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error when writing response: ", err)
	}
}
