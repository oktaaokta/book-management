package handler

import (
	entity "github.com/cosmart/internal/entities"
)

type usecaseInterface interface {
	GetListOfBooks(subject string) ([]entity.Book, error)
	SubmitBookPickupSchedule(edition string) error
}
