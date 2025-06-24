package service

import (
	"log"
	"time"

	"github.com/textures1245/payso-check-slip-backend/handler"
	"github.com/textures1245/payso-check-slip-backend/model"
	"github.com/textures1245/payso-check-slip-backend/repository"
)

type OrderPackageService interface {
	GetOrderPackageAll() ([]model.SureSureOrderPackage, error)
	GetOrderPackageByID(id int) ([]model.SureSureOrderPackage, error)
	GetOrderPackageByRefNo(RefNo string) (model.SureSureOrderPackage, error)
	CreateOrderPackage(pkg model.SureSureOrderPackage) (int, error)
	UpdateOrderPackage(pkg model.SureSureOrderPackage) error
	DeleteOrderPackage(id int) error
	CheckPaymentTransaction() error
}

type orderPackageService struct {
	orderPackageHandler handler.OrderPackageHandler
}

func NewOrderPackageService(orderPackageHandler handler.OrderPackageHandler) OrderPackageService {
	return &orderPackageService{orderPackageHandler}
}

func (s *orderPackageService) GetOrderPackageAll() ([]model.SureSureOrderPackage, error) {
	return repository.GetOrderPackageAll()
}

func (s *orderPackageService) GetOrderPackageByID(id int) ([]model.SureSureOrderPackage, error) {
	return repository.GetOrderPackageByID(id)
}

func (s *orderPackageService) GetOrderPackageByRefNo(RefNo string) (model.SureSureOrderPackage, error) {
	return repository.GetOrderPackageByRefNo(RefNo)
}

func (s *orderPackageService) CreateOrderPackage(pkg model.SureSureOrderPackage) (int, error) {
	return repository.CreateOrderPackage(pkg)
}

func (s *orderPackageService) UpdateOrderPackage(pkg model.SureSureOrderPackage) error {
	return repository.UpdateOrderPackage(pkg)
}

func (s *orderPackageService) DeleteOrderPackage(id int) error {
	return repository.DeleteOrderPackage(id)
}

func (s *orderPackageService) CheckPaymentTransaction() error {
	orderPending, _ := repository.GetOrderPackagPending()
	for _, order := range orderPending {

		log.Printf("Checking payment for order %s", order.RefNo)
		createdDate, err := time.Parse(time.RFC3339, order.CreatedDate)
		if err != nil {
			log.Printf("Error parsing CreatedDate for order %s: %v", order.RefNo, err)
			continue
		}
		// Check payment status
		statusPayment := handler.CheckPaymentTransaction(order.RefNo)
		if statusPayment {
			var orderPackage model.SureSureOrderPackage
			orderPackage.RefNo = order.RefNo
			orderPackage.Status = "SUCCESS"
			repository.UpdateOrderPackage(orderPackage)
			return nil
		}
		if time.Since(createdDate) > 10*time.Minute {
			var orderPackage model.SureSureOrderPackage
			orderPackage.RefNo = order.RefNo
			orderPackage.Status = "CANCEL"
			repository.UpdateOrderPackage(orderPackage)
			return nil
		}
	}
	return nil
}
