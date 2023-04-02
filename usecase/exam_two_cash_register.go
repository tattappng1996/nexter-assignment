package usecase

import (
	"net/http"
	"nexter-assignment/models"
	"nexter-assignment/utils"
)

func (u *usecase) AddCashToCashRegister(cr models.CashRegister) *utils.ErrorMessage {
	oldCR, err := u.repo.GetCashRegister()
	if err != nil {
		return utils.HandleError(utils.ErrorMessage{StatusCode: http.StatusInternalServerError,
			ErrorCode: models.ErrorCashRegisterGeneral, Error: err.Error(), Success: false})
	}
	if len(oldCR.BankNoteAndCoins) == 0 {
		oldCR.BankNoteAndCoins = make([]models.BankNoteAndCoin, len(models.CashRegisterIndexMap))
	}

	for _, bankNoteAndCoin := range cr.BankNoteAndCoins {
		if _, found := models.CashRegisterIndexMap[bankNoteAndCoin.Value]; !found {
			return utils.HandleError(utils.ErrorMessage{StatusCode: http.StatusMethodNotAllowed,
				ErrorCode: models.ErrorNotFoundCashType, Error: "ข้อมูลธนบัตรไม่ถูกต้อง", Success: false})
		}

		cri := models.CashRegisterIndexMap[bankNoteAndCoin.Value]

		oldCR.BankNoteAndCoins[cri.Index].MaxQuantity = cri.MaxQuantity
		oldCR.BankNoteAndCoins[cri.Index].Value = cri.Value
		oldCR.BankNoteAndCoins[cri.Index].Index = cri.Index

		oldCR.BankNoteAndCoins[cri.Index].Quantity += bankNoteAndCoin.Quantity
		if oldCR.BankNoteAndCoins[cri.Index].Quantity > cri.MaxQuantity {
			oldCR.BankNoteAndCoins[cri.Index].Quantity = cri.MaxQuantity
		}
	}

	if err := u.repo.SaveCashRegister(oldCR); err != nil {
		return utils.HandleError(utils.ErrorMessage{StatusCode: http.StatusInternalServerError,
			ErrorCode: models.ErrorCashRegisterGeneral, Error: err.Error(), Success: false})
	}

	return nil
}

func (u *usecase) GetCashRegister() (models.CashRegister, *utils.ErrorMessage) {
	cr, err := u.repo.GetCashRegister()
	if err != nil {
		return cr, utils.HandleError(utils.ErrorMessage{StatusCode: http.StatusInternalServerError,
			ErrorCode: models.ErrorCashRegisterGeneral, Error: err.Error(), Success: false})
	}
	if len(cr.BankNoteAndCoins) == 0 {
		for _, value := range models.CashRegisterBankNotes {
			cri := models.CashRegisterIndexMap[value]

			cr.BankNoteAndCoins = append(cr.BankNoteAndCoins, models.BankNoteAndCoin{
				MaxQuantity: cri.MaxQuantity,
				Value:       cri.Value,
				Index:       cri.Index,
			})
		}
	}

	return cr, nil
}

func (u *usecase) CustomerPaid(pi models.PaidInfo) (models.ChangeInfo, *utils.ErrorMessage) {
	ci := models.ChangeInfo{
		TotalChange: pi.TotalPaid - pi.TotalProductPrice,
	}

	cr, err := u.repo.GetCashRegister()
	if err != nil {
		return ci, utils.HandleError(utils.ErrorMessage{StatusCode: http.StatusInternalServerError,
			ErrorCode: models.ErrorCashRegisterGeneral, Error: err.Error(), Success: false})
	}
	ci.CashRegisterAfterPaid = cr.BankNoteAndCoins

	ci.Changes = make([]models.BankNoteAndCoin, len(models.CashRegisterIndexMap))

	remainingChange := ci.TotalChange

	for i, note := range cr.BankNoteAndCoins {
		ci.Changes[i].Value = note.Value
		ci.Changes[i].Index = note.Index
		ci.Changes[i].MaxQuantity = note.MaxQuantity
		if remainingChange > 0 {
			if remainingChange >= note.Value {
				numNotes := int(remainingChange / note.Value)
				if numNotes > note.Quantity {
					numNotes = note.Quantity
				}
				note.Quantity -= numNotes
				remainingChange -= float64(numNotes) * note.Value

				ci.Changes[i].Quantity = int(numNotes)

				ci.CashRegisterAfterPaid[i].Quantity -= numNotes
			}
		}
	}

	if remainingChange > 0 {
		return ci, utils.HandleError(utils.ErrorMessage{StatusCode: http.StatusMethodNotAllowed,
			ErrorCode: models.ErrorNotEnoughChange, Error: "เงินถอนไม่เพียงพอ", Success: false})
	}

	if err := u.repo.SaveCashRegister(models.CashRegister{
		BankNoteAndCoins: ci.CashRegisterAfterPaid,
	}); err != nil {
		return ci, utils.HandleError(utils.ErrorMessage{StatusCode: http.StatusInternalServerError,
			ErrorCode: models.ErrorCashRegisterGeneral, Error: err.Error(), Success: false})
	}

	return ci, nil
}
