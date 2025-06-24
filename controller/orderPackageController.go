package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/textures1245/payso-check-slip-backend/model"
	"github.com/textures1245/payso-check-slip-backend/service"
	"github.com/textures1245/payso-check-slip-backend/util"
)

type OrderPackageController interface {
	GetOrderPackageAll(c *fiber.Ctx) error
	GetOrderPackageByID(c *fiber.Ctx) error
	GetOrderPackageByRefNo(c *fiber.Ctx) error
	CreateOrderPackage(c *fiber.Ctx) error
	UpdateOrderPackage(c *fiber.Ctx) error
	DeleteOrderPackage(c *fiber.Ctx) error
	CheckPaymentTransaction()
}

type orderPackageController struct {
	OrderPackageService service.OrderPackageService
}

func NewOrderPackageController(orderPackageService service.OrderPackageService) OrderPackageController {
	return &orderPackageController{orderPackageService}
}

func (ctrl *orderPackageController) CheckPaymentTransaction() {
	log.Info("CheckPaymentTransaction")
	_ = ctrl.OrderPackageService.CheckPaymentTransaction()
}

func (ctrl *orderPackageController) GetOrderPackageAll(c *fiber.Ctx) error {
	orderPackages, err := ctrl.OrderPackageService.GetOrderPackageAll()
	if err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, orderPackages)
	return nil
}

func (ctrl *orderPackageController) GetOrderPackageByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	pkg, err := ctrl.OrderPackageService.GetOrderPackageByID(id)
	if err != nil {
		util.JSONResponse(c, fiber.StatusBadRequest, 4000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, pkg)
	return nil
}

func (ctrl *orderPackageController) GetOrderPackageByRefNo(c *fiber.Ctx) error {
	RefNo := c.Params("RefNo")
	pkg, err := ctrl.OrderPackageService.GetOrderPackageByRefNo(RefNo)
	if err != nil {
		util.JSONResponse(c, fiber.StatusBadRequest, 4000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, pkg)
	return nil
}

func (ctrl *orderPackageController) CreateOrderPackage(c *fiber.Ctx) error {
	var pkg model.SureSureOrderPackage
	if err := c.BodyParser(&pkg); err != nil {
		util.JSONResponse(c, fiber.StatusBadRequest, 4000, nil)
		return nil
	}
	log.Info("CreateOrderPackage:::")
	id, err := ctrl.OrderPackageService.CreateOrderPackage(pkg)
	if err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, fiber.Map{"id": id})
	return nil
}

func (ctrl *orderPackageController) UpdateOrderPackage(c *fiber.Ctx) error {
	var pkg model.SureSureOrderPackage
	if err := c.BodyParser(&pkg); err != nil {
		util.JSONResponse(c, fiber.StatusBadRequest, 4000, nil)
		return nil
	}
	if err := ctrl.OrderPackageService.UpdateOrderPackage(pkg); err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, nil)
	return nil
}

func (ctrl *orderPackageController) DeleteOrderPackage(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := ctrl.OrderPackageService.DeleteOrderPackage(id); err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, nil)
	return nil
}
