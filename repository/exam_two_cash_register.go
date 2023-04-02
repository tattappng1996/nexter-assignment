package repository

import (
	"encoding/json"
	"nexter-assignment/models"
	"os"
)

func (r *repository) SaveCashRegister(cr models.CashRegister) error {
	file, err := os.Create(r.cashRegisterFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(cr); err != nil {
		return err
	}

	return nil
}

func (r *repository) GetCashRegister() (models.CashRegister, error) {
	cr := models.CashRegister{}

	crFile, err := os.Open(r.cashRegisterFilePath)
	if err != nil {
		return cr, err
	}
	defer crFile.Close()
	decoder := json.NewDecoder(crFile)
	err = decoder.Decode(&cr)
	if err != nil {
		return cr, err
	}

	return cr, nil
}
