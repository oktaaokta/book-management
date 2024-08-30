package infrastructure

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var (
	mockTime = time.Date(2024, 1, 1, 1, 1, 1, 1, time.UTC)
)

type mockPs struct{}

func (mp *mockPs) toPickupSchedules() *PickupSchedules {
	return &PickupSchedules{
		Info: map[string]ScheduleInformation{
			"science": {
				Schedules: []Schedule{
					{
						PickupDate: mockTime.AddDate(0, 0, 1),
						ReturnDate: mockTime.AddDate(0, 0, 2),
					},
				},
				LastWaitlistDate: mockTime.AddDate(0, 0, 2),
			},
		},
	}
}

func newMockPickupSchedule() *mockPs {
	return &mockPs{}
}

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	got := NewPickupSchedules()
	assert.NotNil(t, got)
}

func TestUsecase_GetPickupSchedules(t *testing.T) {
	tests := []struct {
		name    string
		edition string
		mockFn  func(*mockPs)
		want    ScheduleInformation
	}{
		{
			name:    "no schedules found then return empty",
			edition: "fiction",
			want:    ScheduleInformation{},
		},
		{
			name:    "schedule is found then return schedule",
			edition: "science",
			want: ScheduleInformation{
				Schedules: []Schedule{
					{
						PickupDate: mockTime.AddDate(0, 0, 1),
						ReturnDate: mockTime.AddDate(0, 0, 2),
					},
				},
				LastWaitlistDate: mockTime.AddDate(0, 0, 2),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			pickup := newMockPickupSchedule()
			ps := pickup.toPickupSchedules()

			got := ps.GetPickupSchedules(tt.edition)
			if !reflect.DeepEqual(got, tt.want) {
				assert.Equal(t, tt.want, got, tt.name)
			}
		})
	}
}

func TestUsecase_SetPickupSchedules(t *testing.T) {
	tests := []struct {
		name       string
		edition    string
		pickupDate time.Time
		returnDate time.Time
		mockFn     func(*mockPs)
		want       ScheduleInformation
	}{
		{
			name:       "then add new schedule",
			edition:    "fiction",
			pickupDate: mockTime.AddDate(0, 0, 1),
			returnDate: mockTime.AddDate(0, 0, 2),
			want: ScheduleInformation{
				Schedules: []Schedule{
					{
						PickupDate: mockTime.AddDate(0, 0, 1),
						ReturnDate: mockTime.AddDate(0, 0, 2),
					},
				},
				LastWaitlistDate: mockTime.AddDate(0, 0, 2),
			},
		},
		{
			name:       "schedule is already available, append new one and update last waitlist date",
			edition:    "science",
			pickupDate: mockTime.AddDate(0, 0, 3),
			returnDate: mockTime.AddDate(0, 0, 5),
			want: ScheduleInformation{
				Schedules: []Schedule{
					{
						PickupDate: mockTime.AddDate(0, 0, 1),
						ReturnDate: mockTime.AddDate(0, 0, 2),
					},
					{
						PickupDate: mockTime.AddDate(0, 0, 3),
						ReturnDate: mockTime.AddDate(0, 0, 5),
					},
				},
				LastWaitlistDate: mockTime.AddDate(0, 0, 5),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			pickup := newMockPickupSchedule()
			ps := pickup.toPickupSchedules()

			ps.SetPickupSchedules(tt.edition, tt.pickupDate, tt.returnDate)
			got := ps.Info[tt.edition]
			if !reflect.DeepEqual(got, tt.want) {
				assert.Equal(t, tt.want, got, tt.name)
			}
		})
	}
}
