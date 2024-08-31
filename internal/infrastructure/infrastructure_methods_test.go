package infrastructure

import (
	"reflect"
	"testing"
	"time"

	entity "github.com/cosmart/internal/entities"
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
	type args struct {
		edition    string
		pickupDate time.Time
		returnDate time.Time
		bookInfo   entity.BookInformation
	}
	tests := []struct {
		name   string
		args   args
		mockFn func(*mockPs)
		want   ScheduleInformation
	}{
		{
			name: "schedule is empty, then add new schedule",
			args: args{
				edition:    "fiction",
				pickupDate: mockTime.AddDate(0, 0, 1),
				returnDate: mockTime.AddDate(0, 0, 2),
				bookInfo: entity.BookInformation{
					Title:   "cat in the hat",
					Authors: []string{"Dr. Seuss"},
					Edition: "1",
				},
			},
			want: ScheduleInformation{
				Schedules: []Schedule{
					{
						PickupDate: mockTime.AddDate(0, 0, 1),
						ReturnDate: mockTime.AddDate(0, 0, 2),
					},
				},
				LastWaitlistDate: mockTime.AddDate(0, 0, 2),
				Book: BookInformation{
					Title:   "cat in the hat",
					Authors: []string{"Dr. Seuss"},
					Edition: "1",
				},
			},
		},
		{
			name: "schedule is already available, append new one and update last waitlist date",
			args: args{
				edition:    "science",
				pickupDate: mockTime.AddDate(0, 0, 3),
				returnDate: mockTime.AddDate(0, 0, 5),
				bookInfo:   entity.BookInformation{},
			},
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

			ps.SetPickupSchedules(tt.args.edition, tt.args.pickupDate, tt.args.returnDate, tt.args.bookInfo)
			got := ps.Info[tt.args.edition]
			if !reflect.DeepEqual(got, tt.want) {
				assert.Equal(t, tt.want, got, tt.name)
			}
		})
	}
}
