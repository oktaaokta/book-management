package entities

type Book struct {
	Title      string
	Authors    []string
	EditionKey string
}

type PickupSchedules struct {
	Schedules []BookInformation
}

type BookInformation struct {
	Title   string
	Authors []string
	Edition string
}
