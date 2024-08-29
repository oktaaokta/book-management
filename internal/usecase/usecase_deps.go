package usecase

import (
	entity "github.com/cosmart/internal/entities"
	"github.com/cosmart/internal/infrastructure"
	repo "github.com/cosmart/internal/repository"
)

type repositoryInterface interface {
	// Books methods
	GetBooksByParamFromRepo() (repo.BooksResponse, error)
	GetBooksBySubjectFromRepo(subject string) ([]entity.Book, error)

	// Pickup schedule methods
	GetPickupSchedulesByEdition(edition string) infrastructure.ScheduleInformation
	SetPickupSchedulesByEdition(edition string)
}
