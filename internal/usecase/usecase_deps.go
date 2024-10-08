package usecase

import (
	"time"

	entity "github.com/cosmart/internal/entities"
	"github.com/cosmart/internal/infrastructure"
)

//go:generate mockgen -build_flags=--mod=mod -package=usecase -destination=usecase_deps_mock_test.go -source=usecase_deps.go

type repositoryInterface interface {
	// Books methods
	GetWorkByEdition(edition string) (entity.BookInformation, error)
	GetBooksBySubjectFromRepo(subject string) ([]entity.Book, error)

	// Pickup schedule methods
	GetPickupSchedulesByEdition(edition string) infrastructure.ScheduleInformation
	SetPickupSchedulesByEdition(edition string, pickupDate, returnDate time.Time, bookInfo entity.BookInformation)
}
