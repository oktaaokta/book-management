package handler

type Response struct {
	Message string            `json:"message"`
	Status  int               `json:"status"`
	Books   []BookInformation `json:"books"`
}

type BookInformation struct {
	Title      string   `json:"title"`
	Authors    []string `json:"authors"`
	EditionKey string   `json:"edition"`
}

type PickupScheduleReq struct {
	Edition string `json:"edition"`
}
