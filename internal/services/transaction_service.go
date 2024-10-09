package services

import (
	"errors"
	"inspiration-tech-case/internal/models"
	"inspiration-tech-case/internal/repositories"
)

func NewTransactionService(repository *repositories.TransactionRepository) TransactionService {
	return &TransactionServiceImpl{TransactionRepository: *repository}
}

type TransactionServiceImpl struct {
	repositories.TransactionRepository
}


func (service *TransactionServiceImpl) HandleTransaction(request models.TransactionRequest) error {
	if request.MessageType != "PAYMENT" && request.MessageType != "ADJUSTMENT" {
		return errors.New("invalid message type")
	}
	if request.Origin != "VISA" && request.Origin != "MASTER" {
		return errors.New("invalid origin")
	}
	return service.TransactionRepository.HandleTransaction(request)
}
