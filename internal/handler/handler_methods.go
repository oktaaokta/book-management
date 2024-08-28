package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (hd *Handler) GetHelloWorld(w http.ResponseWriter, r *http.Request) {
	_, err := hd.uc.GetListOfBooks()
	if err != nil {
		return
	}
	response := Response{
		Message: "hi",
		Status:  http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Status)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Error when writing response: ", err)
	}
}
