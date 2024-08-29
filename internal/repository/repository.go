package repository

import (
	infra "github.com/cosmart/internal/infrastructure"
)

type Repository struct {
	ps *infra.PickupSchedules
}

func New(ps *infra.PickupSchedules) *Repository {
	return &Repository{
		ps: ps,
	}
}
