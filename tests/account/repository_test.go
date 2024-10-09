package tests

import (
	"inspiration-tech-case/internal/entities"
	"inspiration-tech-case/internal/repositories"
	"inspiration-tech-case/internal/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func NewMockAccountRepository() repositories.AccountRepository {
	return &MockAccountRepository{}
}

type MockAccountRepository struct{}

func (mock *MockAccountRepository) GetAccountByID(id int64) entities.Account {
	return entities.Account{AccountId: 123, Balance: 100.00}
}
func (mock *MockAccountRepository) Seed() error {
	return nil
}

func TestRespositoryGetAccountByID(t *testing.T) {
	mockRepo := NewMockAccountRepository()
	service := services.NewAccountService(&mockRepo)

	account, err := service.GetAccountByID(123)
	assert.Nil(t, err)
	assert.Equal(t, int64(123), account.AccountId)
	assert.Equal(t, 100.00, account.Balance)
}
