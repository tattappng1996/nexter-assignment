package public

import (
	"net/http"
	"nexter-assignment/models"
	"nexter-assignment/utils"

	"github.com/labstack/echo/v4"
)

func (h *handler) AddCashToCashRegister(c echo.Context) error {
	cr := models.CashRegister{}
	if err := c.Bind(&cr); err != nil {
		return utils.BadRequestError(c, nil)
	}

	if err := h.uc.AddCashToCashRegister(cr); err != nil {
		return c.JSON(err.StatusCode, err)
	}

	return utils.SuccessResponseMessage(http.StatusCreated, nil, c)
}

func (h *handler) GetCashRegister(c echo.Context) error {
	cr, err := h.uc.GetCashRegister()
	if err != nil {
		return c.JSON(err.StatusCode, err)
	}

	return utils.SuccessResponseMessage(http.StatusOK, cr, c)
}

func (h *handler) CustomerPaid(c echo.Context) error {
	pi := models.PaidInfo{}
	if err := c.Bind(&pi); err != nil {
		return utils.BadRequestError(c, nil)
	}
	if pi.TotalPaid < 0 || pi.TotalProductPrice < 0 {
		return utils.BadRequestError(c, nil)
	}

	ci, err := h.uc.CustomerPaid(pi)
	if err != nil {
		return c.JSON(err.StatusCode, err)
	}

	return utils.SuccessResponseMessage(http.StatusOK, ci, c)
}
