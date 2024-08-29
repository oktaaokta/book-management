package infrastructure

import "time"

type PickupSchedules struct {
	Info map[string]ScheduleInformation
}

type ScheduleInformation struct {
	Schedules        []Schedule
	Book             BookInformation
	LastWaitlistDate time.Time
}

type Schedule struct {
	PickupDate time.Time
	ReturnDate time.Time
}

type BookInformation struct {
	Title   string
	Authors []string
	Edition string
}

func NewPickupSchedules() *PickupSchedules {
	return &PickupSchedules{}
}
