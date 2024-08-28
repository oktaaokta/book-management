package usecase

type Usecase struct {
	repo repositoryInterface
}

func New(repo repositoryInterface) *Usecase {
	return &Usecase{
		repo: repo,
	}
}
