package services

import (
	"inspiration-tech-case/internal/models"
)

type AccountService interface {
	GetAccountByID(accountID int64) (models.AccountResponse, error)
}

type TransactionService interface {
	HandleTransaction(models.TransactionRequest) error
}