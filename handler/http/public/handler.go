package public

import (
	"nexter-assignment/models"
	"nexter-assignment/utils"
)

// Usecase ..
type Usecase interface {
	AddCashToCashRegister(models.CashRegister) *utils.ErrorMessage
	GetCashRegister() (models.CashRegister, *utils.ErrorMessage)
	CustomerPaid(models.PaidInfo) (models.ChangeInfo, *utils.ErrorMessage)
}

type handler struct {
	uc Usecase
}

func newHandler(usecase Usecase) *handler {
	return &handler{
		uc: usecase,
	}
}
