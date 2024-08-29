package handler

import (
	"time"

	entity "github.com/cosmart/internal/entities"
)

type usecaseInterface interface {
	GetListOfBooks(subject string) ([]entity.Book, error)
	SubmitBookPickupSchedule(edition string, pickupDate, returnDate time.Time) error
}
