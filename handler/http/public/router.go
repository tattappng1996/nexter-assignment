package public

import (
	"github.com/labstack/echo/v4"
)

// NewRouter ..
func NewRouter(e *echo.Group, uc Usecase) *handler {
	h := newHandler(uc)

	e.GET("-one", h.GetExamOne)

	e.POST("-two/v1/cash-register", h.AddCashToCashRegister)
	e.GET("-two/v1/cash-register", h.GetCashRegister)
	e.POST("-two/v1/customer-paid", h.CustomerPaid)

	return h
}
