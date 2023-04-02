package usecase_test

import (
	"fmt"
	"net/http"
	"nexter-assignment/models"
	"nexter-assignment/utils"
)

func (t *TestUsecaseSuite) TestAddCashToCashRegister() {
	t.Run("Should AddCashToCashRegister GetCashRegister StatusInternalServerError", func() {
		t.mockRepo.EXPECT().
			GetCashRegister().
			Return(models.CashRegister{}, fmt.Errorf("error"))

		actualErrMessage := t.publicUC.AddCashToCashRegister(models.CashRegister{})
		expected := utils.HandleError(utils.ErrorMessage{StatusCode: http.StatusInternalServerError,
			ErrorCode: models.ErrorCashRegisterGeneral, Error: "Internal Server Error", Success: false})
		t.Equal(expected, actualErrMessage)
	})

	t.Run("Should AddCashToCashRegister StatusMethodNotAllowed", func() {
		t.mockRepo.EXPECT().
			GetCashRegister().
			Return(models.CashRegister{}, nil)

		actualErrMessage := t.publicUC.AddCashToCashRegister(models.CashRegister{
			BankNoteAndCoins: []models.BankNoteAndCoin{
				{
					Value: -999,
				},
			},
		})
		expected := utils.HandleError(utils.ErrorMessage{StatusCode: http.StatusMethodNotAllowed,
			ErrorCode: models.ErrorNotFoundCashType, Error: "ข้อมูลธนบัตรไม่ถูกต้อง", Success: false})
		t.Equal(expected, actualErrMessage)
	})

	t.Run("Should AddCashToCashRegister SaveCashRegister StatusInternalServerError", func() {
		t.mockRepo.EXPECT().
			GetCashRegister().
			Return(models.CashRegister{}, nil)

		t.mockRepo.EXPECT().
			SaveCashRegister(models.CashRegister{
				BankNoteAndCoins: []models.BankNoteAndCoin{
					{
						Index:       0,
						Value:       1000,
						Quantity:    2,
						MaxQuantity: 10,
					},
					{
						Index:       1,
						Value:       500,
						Quantity:    0,
						MaxQuantity: 20,
					},
					{
						Index:       2,
						Value:       100,
						Quantity:    0,
						MaxQuantity: 15,
					},
					{
						Index:       3,
						Value:       50,
						Quantity:    0,
						MaxQuantity: 20,
					},
					{
						Index:       4,
						Value:       20,
						Quantity:    0,
						MaxQuantity: 30,
					},
					{
						Index:       5,
						Value:       10,
						Quantity:    0,
						MaxQuantity: 20,
					},
					{
						Index:       6,
						Value:       5,
						Quantity:    0,
						MaxQuantity: 20,
					},
					{
						Index:       7,
						Value:       1,
						Quantity:    0,
						MaxQuantity: 20,
					},
					{
						Index:       8,
						Value:       0.25,
						Quantity:    0,
						MaxQuantity: 50,
					},
				},
			}).
			Return(fmt.Errorf("error"))

		actualErrMessage := t.publicUC.AddCashToCashRegister(models.CashRegister{
			BankNoteAndCoins: []models.BankNoteAndCoin{
				{
					Value:    1000,
					Quantity: 2,
				},
			},
		})
		expected := utils.HandleError(utils.ErrorMessage{StatusCode: http.StatusInternalServerError,
			ErrorCode: models.ErrorCashRegisterGeneral, Error: "Internal Server Error", Success: false})
		t.Equal(expected, actualErrMessage)
	})

	t.Run("Should AddCashToCashRegister Success", func() {
		t.mockRepo.EXPECT().
			GetCashRegister().
			Return(models.CashRegister{}, nil)

		t.mockRepo.EXPECT().
			SaveCashRegister(models.CashRegister{
				BankNoteAndCoins: []models.BankNoteAndCoin{
					{
						Index:       0,
						Value:       1000,
						Quantity:    10,
						MaxQuantity: 10,
					},
					{
						Index:       1,
						Value:       500,
						Quantity:    0,
						MaxQuantity: 20,
					},
					{
						Index:       2,
						Value:       100,
						Quantity:    0,
						MaxQuantity: 15,
					},
					{
						Index:       3,
						Value:       50,
						Quantity:    0,
						MaxQuantity: 20,
					},
					{
						Index:       4,
						Value:       20,
						Quantity:    0,
						MaxQuantity: 30,
					},
					{
						Index:       5,
						Value:       10,
						Quantity:    0,
						MaxQuantity: 20,
					},
					{
						Index:       6,
						Value:       5,
						Quantity:    0,
						MaxQuantity: 20,
					},
					{
						Index:       7,
						Value:       1,
						Quantity:    0,
						MaxQuantity: 20,
					},
					{
						Index:       8,
						Value:       0.25,
						Quantity:    0,
						MaxQuantity: 50,
					},
				},
			}).
			Return(nil)

		actualErrMessage := t.publicUC.AddCashToCashRegister(models.CashRegister{
			BankNoteAndCoins: []models.BankNoteAndCoin{
				{
					Value:    1000,
					Quantity: 20,
				},
			},
		})
		var expected *utils.ErrorMessage
		t.Equal(expected, actualErrMessage)
	})
}
