package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/textures1245/payso-check-slip-backend/model"
	"github.com/textures1245/payso-check-slip-backend/service"
	"github.com/textures1245/payso-check-slip-backend/util"
)

type BanksController interface {
	GetAllBanks(c *fiber.Ctx) error
	GetBankByID(c *fiber.Ctx) error
	CreateBank(c *fiber.Ctx) error
	UpdateBank(c *fiber.Ctx) error
	DeleteBank(c *fiber.Ctx) error
}

type banksController struct {
	BanksService service.BanksService
}

func NewBanksController(banksService service.BanksService) BanksController {
	return &banksController{banksService}
}

// GetAllBanks godoc
// @Summary List banks
// @Tags Bank
// @Produce json
// @Success 200 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /bank2/get [get]
func (ctrl *banksController) GetAllBanks(c *fiber.Ctx) error {
	banks, err := ctrl.BanksService.GetAllBank()
	if err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, banks)
	return nil
}

// GetBankByID godoc
// @Summary Get bank by ID
// @Tags Bank
// @Produce json
// @Param id path int true "Bank ID"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /bank2/get/{id} [get]
func (ctrl *banksController) GetBankByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	bank, err := ctrl.BanksService.GetBankByID(id)
	if err != nil {
		util.JSONResponse(c, fiber.StatusBadRequest, 4000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, bank)
	return nil
}

// CreateBank godoc
// @Summary Create bank
// @Tags Bank
// @Accept json
// @Produce json
// @Param bank body model.SureSureBank true "Bank payload"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /bank2/create [post]
func (ctrl *banksController) CreateBank(c *fiber.Ctx) error {
	var bank model.SureSureBank
	if err := c.BodyParser(&bank); err != nil {
		util.JSONResponse(c, fiber.StatusBadRequest, 4000, nil)
		return nil
	}
	id, err := ctrl.BanksService.CreateBank(bank)
	if err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, fiber.Map{"id": id})
	return nil
}

// UpdateBank godoc
// @Summary Update bank
// @Tags Bank
// @Accept json
// @Produce json
// @Param bank body model.SureSureBank true "Bank payload"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /bank2/update [put]
func (ctrl *banksController) UpdateBank(c *fiber.Ctx) error {
	var bank model.SureSureBank
	if err := c.BodyParser(&bank); err != nil {
		util.JSONResponse(c, fiber.StatusBadRequest, 4000, nil)
		return nil
	}
	if err := ctrl.BanksService.UpdateBank(bank); err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, nil)
	return nil
}

// DeleteBank godoc
// @Summary Delete bank by ID
// @Tags Bank
// @Produce json
// @Param id path int true "Bank ID"
// @Success 200 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /bank2/delete/{id} [delete]
func (ctrl *banksController) DeleteBank(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := ctrl.BanksService.DeleteBank(id); err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, nil)
	return nil
}
