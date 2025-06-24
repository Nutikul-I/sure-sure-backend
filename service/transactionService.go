package service

import (
	"github.com/textures1245/payso-check-slip-backend/handler"
	"github.com/textures1245/payso-check-slip-backend/model"
	"github.com/textures1245/payso-check-slip-backend/repository"
)

type TransactionService interface {
	GetTransactionAll() ([]model.SureSureTransaction, error)
	GetTransactionByID(id int) ([]model.SureSureTransaction, error)
	CreateTransaction(transaction model.SureSureTransaction) (int, error)
	UpdateTransaction(transaction model.SureSureTransaction) error
	DeleteTransaction(id int) error
}

type transactionService struct {
	transactionHandler handler.TransactionHandler
}

func NewTransactionService(transactionHandler handler.TransactionHandler) TransactionService {
	return &transactionService{transactionHandler}
}

func (s *transactionService) GetTransactionAll() ([]model.SureSureTransaction, error) {
	return repository.GetTransactionAll()
}

func (s *transactionService) GetTransactionByID(id int) ([]model.SureSureTransaction, error) {
	return repository.GetTransactionByID(id)
}

func (s *transactionService) CreateTransaction(transaction model.SureSureTransaction) (int, error) {
	return repository.CreateTransaction(transaction)
}

func (s *transactionService) UpdateTransaction(transaction model.SureSureTransaction) error {
	return repository.UpdateTransaction(transaction)
}

func (s *transactionService) DeleteTransaction(id int) error {
	return repository.DeleteTransaction(id)
}
