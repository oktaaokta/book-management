package handler

import (
	"time"

	entity "github.com/cosmart/internal/entities"
)

//go:generate mockgen -build_flags=--mod=mod -package=handler -destination=handler_deps_mock_test.go -source=handler_deps.go

type usecaseInterface interface {
	GetListOfBooks(subject string) ([]entity.Book, error)
	SubmitBookPickupSchedule(edition string, pickupDate, returnDate time.Time) error
}
