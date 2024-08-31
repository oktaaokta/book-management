package infrastructure

import (
	"time"

	entity "github.com/cosmart/internal/entities"
)

func (ps *PickupSchedules) GetPickupSchedules(edition string) ScheduleInformation {
	if schedule, exists := ps.Info[edition]; exists {
		return schedule
	}

	return ScheduleInformation{}
}

func (ps *PickupSchedules) SetPickupSchedules(edition string, pickupDate, returnDate time.Time, bookInfo entity.BookInformation) {
	schedule, exists := ps.Info[edition]
	if !exists {
		var scheduleInfo ScheduleInformation
		scheduleInfo.Schedules = []Schedule{
			{
				PickupDate: pickupDate,
				ReturnDate: returnDate,
			},
		}
		scheduleInfo.LastWaitlistDate = returnDate
		scheduleInfo.Book = BookInformation{
			Title:   bookInfo.Title,
			Edition: bookInfo.Edition,
			Authors: bookInfo.Authors,
		}
		ps.Info[edition] = scheduleInfo
		return
	}

	schedule.Schedules = append(schedule.Schedules, Schedule{
		PickupDate: pickupDate,
		ReturnDate: returnDate,
	})
	schedule.LastWaitlistDate = returnDate
	ps.Info[edition] = schedule
}
