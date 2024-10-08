package main

import (
	"fmt"
	"net/http"

	"github.com/cosmart/internal/handler"
	"github.com/gorilla/mux"
)

func serveHTTP(hd *handler.Handler) {
	router := mux.NewRouter()
	router.HandleFunc("/api/books", hd.GetBooksList).Methods(http.MethodGet)
	router.HandleFunc("/api/books/pickup", hd.SubmitBookPickupSchedule).Methods(http.MethodPost)

	fmt.Println("Listening on port :8000")
	// Start the server and pass the router
	if err := http.ListenAndServe(":8000", router); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
