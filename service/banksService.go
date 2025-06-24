package service

import (
	"github.com/textures1245/payso-check-slip-backend/handler"
	"github.com/textures1245/payso-check-slip-backend/model"
	"github.com/textures1245/payso-check-slip-backend/repository"
)

type BanksService interface {
	GetAllBank() ([]model.SureSureBank, error)
	GetBankByID(id int) ([]model.SureSureBank, error)
	CreateBank(bank model.SureSureBank) (int, error)
	UpdateBank(bank model.SureSureBank) error
	DeleteBank(id int) error
}

type banksService struct {
	banksHandler handler.BanksHandler
}

func NewBanksService(banksHandler handler.BanksHandler) BanksService {
	return &banksService{banksHandler}
}

func (s *banksService) GetAllBank() ([]model.SureSureBank, error) {
	return repository.GetAllBank()
}

func (s *banksService) GetBankByID(id int) ([]model.SureSureBank, error) {
	return repository.GetBankByID(id)
}

func (s *banksService) CreateBank(bank model.SureSureBank) (int, error) {
	return repository.CreateBank(bank)
}

func (s *banksService) UpdateBank(bank model.SureSureBank) error {
	return repository.UpdateBank(bank)
}

func (s *banksService) DeleteBank(id int) error {
	return repository.DeleteBank(id)
}
