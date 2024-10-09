package services

import (
	"errors"
	"inspiration-tech-case/internal/models"
	"inspiration-tech-case/internal/repositories"
)

func NewAccountService(repository *repositories.AccountRepository) AccountService {
	return &AccountServiceImpl{AccountRepository: *repository}
}

type AccountServiceImpl struct {
	repositories.AccountRepository
}

func (service *AccountServiceImpl) GetAccountByID(accountId int64) (models.AccountResponse, error) {
	account := service.AccountRepository.GetAccountByID(accountId)
	if account.AccountId > 0 {
		response := models.AccountResponse{AccountId: account.AccountId, Balance: account.Balance}
		return response, nil
	} else {
		return models.AccountResponse{}, errors.New("account not found")
	}

}
