package service

import (
	"github.com/textures1245/payso-check-slip-backend/handler"
	"github.com/textures1245/payso-check-slip-backend/model"
	"github.com/textures1245/payso-check-slip-backend/repository"
)

type PackageService interface {
	GetPackageAll() ([]model.SureSurePackage, error)
	GetPackageByID(id int) (model.SureSurePackage, error)
	CreatePackage(pkg model.SureSurePackage) (int, error)
	UpdatePackage(pkg model.SureSurePackage) error
	DeletePackage(id int) error
}

type packageService struct {
	packageHandler handler.PackageHandler
}

func NewPackageService(packageHandler handler.PackageHandler) PackageService {
	return &packageService{packageHandler}
}

func (s *packageService) GetPackageAll() ([]model.SureSurePackage, error) {
	return repository.GetPackageAll()
}

func (s *packageService) GetPackageByID(id int) (model.SureSurePackage, error) {
	return repository.GetPackageByID(id)
}

func (s *packageService) CreatePackage(pkg model.SureSurePackage) (int, error) {
	return repository.CreatePackage(pkg)
}

func (s *packageService) UpdatePackage(pkg model.SureSurePackage) error {
	return repository.UpdatePackage(pkg)
}

func (s *packageService) DeletePackage(id int) error {
	return repository.DeletePackage(id)
}
