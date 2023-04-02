package usecase

import (
	"nexter-assignment/models"
)

// Repository ..
type Repository interface {
	SaveCashRegister(models.CashRegister) error
	GetCashRegister() (models.CashRegister, error)
}

// NewUsecase ..
func NewUsecase(repo Repository) *usecase {
	return &usecase{
		repo: repo,
	}
}

type usecase struct {
	repo Repository
}
