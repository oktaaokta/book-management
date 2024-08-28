package main

import (
	"github.com/cosmart/internal/handler"
	repo "github.com/cosmart/internal/repository"
	uc "github.com/cosmart/internal/usecase"
)

func initRepository() *repo.Repository {
	ps := NewPickupSchedules()
	return repo.New(ps)
}

func initUsecase(repo *repo.Repository) *uc.Usecase {
	return uc.New(repo)
}

func initHandler(uc *uc.Usecase) *handler.Handler {
	return handler.New(uc)
}
