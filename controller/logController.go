package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/textures1245/payso-check-slip-backend/model"
	"github.com/textures1245/payso-check-slip-backend/service"
	"github.com/textures1245/payso-check-slip-backend/util"
)

type LogController interface {
	GetLogAll(c *fiber.Ctx) error
	GetLogByID(c *fiber.Ctx) error
	CreateLog(c *fiber.Ctx) error
	UpdateLog(c *fiber.Ctx) error
	DeleteLog(c *fiber.Ctx) error
}

type logController struct {
	LogService service.LogService
}

func NewLogController(logService service.LogService) LogController {
	return &logController{logService}
}

func (ctrl *logController) GetLogAll(c *fiber.Ctx) error {
	logs, err := ctrl.LogService.GetLogAll()
	if err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, logs)
	return nil
}

func (ctrl *logController) GetLogByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	logs, err := ctrl.LogService.GetLogByID(id)
	if err != nil {
		util.JSONResponse(c, fiber.StatusBadRequest, 4000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, logs)
	return nil
}

func (ctrl *logController) CreateLog(c *fiber.Ctx) error {
	var logs model.SureSureLog
	if err := c.BodyParser(&logs); err != nil {
		util.JSONResponse(c, fiber.StatusBadRequest, 4000, nil)
		return nil
	}
	id, err := ctrl.LogService.CreateLog(logs)
	log.Infof("id: %d", id)
	if err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, nil)
	return nil
}

func (ctrl *logController) UpdateLog(c *fiber.Ctx) error {
	var logs model.SureSureLog
	if err := c.BodyParser(&logs); err != nil {
		util.JSONResponse(c, fiber.StatusBadRequest, 4000, nil)
		return nil
	}
	if err := ctrl.LogService.UpdateLog(logs); err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, nil)
	return nil
}

func (ctrl *logController) DeleteLog(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	if err := ctrl.LogService.DeleteLog(id); err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, nil)
	return nil
}
