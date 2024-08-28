package repository

type Repository struct {
	ps pickupSchedulesInterface
}

func New(ps pickupSchedulesInterface) *Repository {
	return &Repository{
		ps: ps,
	}
}
