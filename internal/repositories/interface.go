package repositories

import (
	"inspiration-tech-case/internal/entities"
	"inspiration-tech-case/internal/models"
)

type AccountRepository interface {
	Seed() error
	GetAccountByID(accountID int64) entities.Account
}

type TransactionRepository interface {
	HandleTransaction(models.TransactionRequest) error
}

