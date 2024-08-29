package infrastructure

import "time"

func (ps *PickupSchedules) GetPickupSchedules(edition string) ScheduleInformation {
	if schedule, exists := ps.Info[edition]; !exists {
		return schedule
	}

	return ScheduleInformation{}
}

func (ps *PickupSchedules) SetPickupSchedules(edition string) {
	schedule, exists := ps.Info[edition]
	if !exists {
		scheduleInfo := ScheduleInformation{}
		scheduleInfo.Schedules = []Schedule{
			{
				PickupDate: time.Now(),
				ReturnDate: time.Now(),
			},
		}
		scheduleInfo.LastWaitlistDate = time.Now()
		ps.Info[edition] = scheduleInfo
		return
	}
	schedule.Schedules = append(schedule.Schedules, Schedule{
		PickupDate: time.Now(),
		ReturnDate: time.Now(),
	})
	schedule.LastWaitlistDate = time.Now()
	ps.Info[edition] = schedule
}
