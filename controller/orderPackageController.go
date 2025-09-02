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

// GetOrderPackageAll godoc
// @Summary List order packages
// @Tags OrderPackage
// @Produce json
// @Success 200 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /order-package/get [get]
func (ctrl *orderPackageController) GetOrderPackageAll(c *fiber.Ctx) error {
	orderPackages, err := ctrl.OrderPackageService.GetOrderPackageAll()
	if err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, orderPackages)
	return nil
}

// GetOrderPackageByID godoc
// @Summary Get order package by user ID
// @Tags OrderPackage
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /order-package/get/{id} [get]
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

// GetOrderPackageByRefNo godoc
// @Summary Get order package by RefNo
// @Tags OrderPackage
// @Produce json
// @Param RefNo path string true "Reference No"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /order-package/get/refno/{RefNo} [get]
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

// CreateOrderPackage godoc
// @Summary Create order package
// @Tags OrderPackage
// @Accept json
// @Produce json
// @Param order body model.SureSureOrderPackage true "Order package payload"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /order-package/create [post]
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

// UpdateOrderPackage godoc
// @Summary Update order package
// @Tags OrderPackage
// @Accept json
// @Produce json
// @Param order body model.SureSureOrderPackage true "Order package payload"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /order-package/update [put]
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

// DeleteOrderPackage godoc
// @Summary Delete order package by ID
// @Tags OrderPackage
// @Produce json
// @Param id path int true "Order package ID"
// @Success 200 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /order-package/delete/{id} [delete]
func (ctrl *orderPackageController) DeleteOrderPackage(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := ctrl.OrderPackageService.DeleteOrderPackage(id); err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, nil)
	return nil
}
