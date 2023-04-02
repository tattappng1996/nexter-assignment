package repository

import (
	"nexter-assignment/usecase"
)

// NewRepository ..
func NewRepository(cashRegisterFilePath, cashStorageFilePath string) usecase.Repository {
	return &repository{
		cashRegisterFilePath: cashRegisterFilePath,
		cashStorageFilePath:  cashStorageFilePath,
	}
}

type repository struct {
	cashRegisterFilePath string
	cashStorageFilePath  string
}
