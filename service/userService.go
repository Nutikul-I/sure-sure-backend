package service

import (
	"github.com/textures1245/payso-check-slip-backend/handler"
	"github.com/textures1245/payso-check-slip-backend/model"
	"github.com/textures1245/payso-check-slip-backend/repository"
)

type UserService interface {
	GetUserAll() ([]model.SureSureUser, error)
	GetOrCreateUser(user model.SureSureUser) (model.SureSureUser, error)
	RegisterUser(user model.SureSureUser) (model.SureSureUser, error)
	GetUserByID(id string) (model.SureSureUser, error)
	CreateUser(user model.SureSureUser) (string, error)
	UpdateUser(user model.SureSureUser) error
	DeleteUser(id string) error
	GetCategoryAll() ([]model.MerchantCategory, error)
}

type userService struct {
	userHandler handler.UserHandler
}

func NewUserService(userHandler handler.UserHandler) UserService {
	return &userService{userHandler}
}

func (s *userService) GetUserAll() ([]model.SureSureUser, error) {
	return repository.GetUserAll()
}

func (s *userService) GetOrCreateUser(user model.SureSureUser) (model.SureSureUser, error) {
	return repository.GetOrCreateUser(user)
}

func (s *userService) RegisterUser(user model.SureSureUser) (model.SureSureUser, error) {
	return repository.RegisterUser(user)
}

func (s *userService) GetUserByID(id string) (model.SureSureUser, error) {
	return repository.GetUserByID(id)
}

func (s *userService) CreateUser(user model.SureSureUser) (string, error) {
	return repository.CreateUser(user)
}

func (s *userService) UpdateUser(user model.SureSureUser) error {
	return repository.UpdateUser(user)
}

func (s *userService) DeleteUser(id string) error {
	return repository.DeleteUser(id)
}

func (s *userService) GetCategoryAll() ([]model.MerchantCategory, error) {
	return repository.GetCategoryAll()
}
