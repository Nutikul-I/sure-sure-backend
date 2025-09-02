package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/textures1245/payso-check-slip-backend/model"
	"github.com/textures1245/payso-check-slip-backend/service"
	"github.com/textures1245/payso-check-slip-backend/util"
)

type TransactionController interface {
	GetTransactionAll(c *fiber.Ctx) error
	GetTransactionByID(c *fiber.Ctx) error
	CreateTransaction(c *fiber.Ctx) error
	UpdateTransaction(c *fiber.Ctx) error
	DeleteTransaction(c *fiber.Ctx) error
}

type transactionController struct {
	TransactionService service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService) TransactionController {
	return &transactionController{transactionService}
}

// GetTransactionAll godoc
// @Summary List transactions
// @Tags Transaction
// @Produce json
// @Success 200 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /transaction/get [get]
func (ctrl *transactionController) GetTransactionAll(c *fiber.Ctx) error {
	transactions, err := ctrl.TransactionService.GetTransactionAll()
	if err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, transactions)
	return nil
}

// GetTransactionByID godoc
// @Summary Get transaction by user ID
// @Tags Transaction
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /transaction/get/{id} [get]
func (ctrl *transactionController) GetTransactionByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	transaction, err := ctrl.TransactionService.GetTransactionByID(id)
	if err != nil {
		util.JSONResponse(c, fiber.StatusBadRequest, 4000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, transaction)
	return nil
}

// CreateTransaction godoc
// @Summary Create transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Param transaction body model.SureSureTransaction true "Transaction payload"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /transaction/create [post]
func (ctrl *transactionController) CreateTransaction(c *fiber.Ctx) error {
	var transaction model.SureSureTransaction
	if err := c.BodyParser(&transaction); err != nil {
		util.JSONResponse(c, fiber.StatusBadRequest, 4000, nil)
		return nil
	}
	id, err := ctrl.TransactionService.CreateTransaction(transaction)
	log.Infof("id: %d", id)
	if err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, nil)
	return nil
}

// UpdateTransaction godoc
// @Summary Update transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Param transaction body model.SureSureTransaction true "Transaction payload"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /transaction/update [put]
func (ctrl *transactionController) UpdateTransaction(c *fiber.Ctx) error {
	var transaction model.SureSureTransaction
	if err := c.BodyParser(&transaction); err != nil {
		util.JSONResponse(c, fiber.StatusBadRequest, 4000, nil)
		return nil
	}
	if err := ctrl.TransactionService.UpdateTransaction(transaction); err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, nil)
	return nil
}

// DeleteTransaction godoc
// @Summary Delete transaction by ID
// @Tags Transaction
// @Produce json
// @Param id path int true "Transaction ID"
// @Success 200 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /transaction/delete/{id} [delete]
func (ctrl *transactionController) DeleteTransaction(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	if err := ctrl.TransactionService.DeleteTransaction(id); err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, nil)
	return nil
}
