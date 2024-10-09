package repositories

import (
	"errors"
	"inspiration-tech-case/internal/entities"
	"inspiration-tech-case/internal/models"
	"inspiration-tech-case/internal/utils"
	"strconv"
	"time"

	"github.com/patrickmn/go-cache"
)

func NewTransactionRepository(DB *cache.Cache) TransactionRepository {
	return &TransactionRepositoryImpl{DB: DB}
}

type TransactionRepositoryImpl struct {
	DB *cache.Cache
}

func (repository *TransactionRepositoryImpl) HandleTransaction(request models.TransactionRequest) error {
	accountData, ok := repository.DB.Get(strconv.FormatInt(request.AccountId, 10))
	if !ok {
		return errors.New("account not found")
	}

	account, ok := accountData.(entities.Account)
	if !ok {
		return errors.New("account not found")
	}

	commissionAmount := calculateCommission(request.Amount, request.Origin)
	totalAmount := request.Amount + commissionAmount
	if request.MessageType == "PAYMENT" {
		isValid := account.ValidateDuplicateTransaction(request.TransactionId)
		if !isValid {
			return errors.New("transaction denied(duplicate)")
		}
		if totalAmount > account.Balance {
			return errors.New("insufficient balance")
		}
		account.Balance -= totalAmount
	} else if request.MessageType == "ADJUSTMENT" {
		isValid := account.ValidateAdjustmentTransactionAmount(request.TransactionId, totalAmount)
		if !isValid {
			return errors.New("transaction denied(amount)")
		}
		account.Balance += totalAmount
	}

	newTransaction := entities.Transaction{
		TransactionId: request.TransactionId,
		MessageType:   request.MessageType,
		Origin:        request.Origin,
		Amount:        request.Amount,
		Commission:    commissionAmount,
		Timestamp:     time.Now(),
	}

	account.Transactions = append(account.Transactions, newTransaction)
	repository.DB.Set(strconv.FormatInt(account.AccountId, 10), account, cache.NoExpiration)
	return nil
}

func calculateCommission(amount float64, origin string) float64 {
	switch origin {
	case "VISA":
		return amount * float64(utils.VisaCommissionRate) / 100
	case "MASTER":
		return amount * float64(utils.MasterCommissionRate) / 100
	default:
		return 0
	}
}
