package handler

import "time"

type Response struct {
	Message string            `json:"message"`
	Books   []BookInformation `json:"books"`
}

type BookInformation struct {
	Title      string   `json:"title"`
	Authors    []string `json:"authors"`
	EditionKey string   `json:"edition"`
}

type PickupScheduleReq struct {
	Edition    string    `json:"edition"`
	PickupDate time.Time `json:"pickup_date"`
	ReturnDate time.Time `json:"return_date"`
}
