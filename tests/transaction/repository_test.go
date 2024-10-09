package tests

import (
	"errors"
	"inspiration-tech-case/internal/models"
	"inspiration-tech-case/internal/repositories"
	"inspiration-tech-case/internal/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func NewMockTransactionRepository() repositories.TransactionRepository {
	return &MockTransactionRepository{}
}

type MockTransactionRepository struct{}

func (mock *MockTransactionRepository) HandleTransaction(request models.TransactionRequest) error {
	return nil
}

func TestRespositoryHandleTransaction(t *testing.T) {
	mockRepo := NewMockTransactionRepository()
	service := services.NewTransactionService(&mockRepo)

	err := service.HandleTransaction(models.TransactionRequest{})
	assert.Equal(t, err, errors.New("invalid message type"))

	err = service.HandleTransaction(models.TransactionRequest{MessageType: "PAYMENT"})
	assert.Equal(t, err, errors.New("invalid origin"))

	err = service.HandleTransaction(models.TransactionRequest{MessageType: "PAYMENT",Origin: "VISA"})
	assert.Nil(t, err)
}
