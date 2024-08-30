package usecase

import (
	"errors"
	"reflect"
	"testing"
	"time"

	entity "github.com/cosmart/internal/entities"
	"github.com/cosmart/internal/infrastructure"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var (
	mockTime = time.Date(2024, 1, 1, 1, 1, 1, 1, time.UTC)
)

type mockUsecase struct {
	repo *MockrepositoryInterface
}

func (mu *mockUsecase) toUsecase() *Usecase {
	return &Usecase{
		repo: mu.repo,
	}
}

func newMockUsecase(ctrl *gomock.Controller) *mockUsecase {
	return &mockUsecase{
		repo: NewMockrepositoryInterface(ctrl),
	}
}

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mu := newMockUsecase(ctrl)

	got := New(mu.repo)
	assert.NotNil(t, got)
}

func TestUsecase_GetListOfBooks(t *testing.T) {
	tests := []struct {
		name    string
		subject string
		mockFn  func(*mockUsecase)
		want    []entity.Book
		wantErr bool
	}{
		{
			name:    "case success return books",
			subject: "science",
			mockFn: func(mu *mockUsecase) {
				mu.repo.EXPECT().GetBooksBySubjectFromRepo(gomock.Any()).Return([]entity.Book{
					{
						Title:      "Hello World",
						Authors:    []string{"Mr.A"},
						EditionKey: "edition",
					},
				}, nil)
			},
			want: []entity.Book{
				{
					Title:      "Hello World",
					Authors:    []string{"Mr.A"},
					EditionKey: "edition",
				},
			},
			wantErr: false,
		},
		{
			name:    "error when getting books then return error",
			subject: "science",
			mockFn: func(mu *mockUsecase) {
				mu.repo.EXPECT().GetBooksBySubjectFromRepo(gomock.Any()).Return(nil, errors.New("no books"))
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mu := newMockUsecase(ctrl)
			tt.mockFn(mu)
			uc := mu.toUsecase()

			got, err := uc.GetListOfBooks(tt.subject)
			if (err != nil) != tt.wantErr {
				assert.Equal(t, tt.wantErr, err, tt.name)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				assert.Equal(t, tt.want, got, tt.name)
			}
		})
	}
}

func TestUsecase_SubmitBookPickupSchedule(t *testing.T) {
	tests := []struct {
		name                   string
		subject                string
		pickupDate, returnDate time.Time
		mockFn                 func(*mockUsecase)
		wantErr                bool
	}{
		{
			name:       "case pickup schedules ",
			subject:    "science",
			pickupDate: mockTime,
			returnDate: mockTime.AddDate(0, 0, 1),
			mockFn: func(mu *mockUsecase) {
				mu.repo.EXPECT().GetPickupSchedulesByEdition(gomock.Any()).Return(infrastructure.ScheduleInformation{
					Schedules: []infrastructure.Schedule{
						{
							PickupDate: mockTime.AddDate(0, 0, 1),
							ReturnDate: mockTime.AddDate(0, 0, 2),
						},
					},
					LastWaitlistDate: mockTime.AddDate(0, 0, 2),
				}, nil)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mu := newMockUsecase(ctrl)
			tt.mockFn(mu)
			uc := mu.toUsecase()

			err := uc.SubmitBookPickupSchedule(tt.subject, tt.pickupDate, tt.returnDate)
			if (err != nil) != tt.wantErr {
				assert.Equal(t, tt.wantErr, err, tt.name)
				return
			}
		})
	}
}
