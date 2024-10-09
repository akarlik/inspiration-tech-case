package tests

import (
	"inspiration-tech-case/internal/entities"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccount_ValidateAdjustmentTransactionAmount(t *testing.T) {
	transactions := []entities.Transaction{}
	transactions = append(transactions, entities.Transaction{TransactionId: "1", MessageType: "PAYMENT", Origin: "VISA", Amount: 10, Commission: 0, Timestamp: time.Now()})
	transactions = append(transactions, entities.Transaction{TransactionId: "1", MessageType: "ADJUSTMENT", Origin: "VISA", Amount: 5, Commission: 0, Timestamp: time.Now()})
	account := entities.Account{AccountId: 1, Balance: 0, Transactions: transactions}
	actual := account.ValidateAdjustmentTransactionAmount("1", 4)
	assert.True(t, actual)
	actual = account.ValidateAdjustmentTransactionAmount("1", 6)
	assert.False(t, actual)
}
