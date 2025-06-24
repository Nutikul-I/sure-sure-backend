package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/textures1245/payso-check-slip-backend/model"
	"github.com/textures1245/payso-check-slip-backend/service"
	"github.com/textures1245/payso-check-slip-backend/util"
)

type RoomsController interface {
	GetAllRooms(c *fiber.Ctx) error
	GetRoomByID(c *fiber.Ctx) error
	CreateRoom(c *fiber.Ctx) error
	UpdateRoom(c *fiber.Ctx) error
	DeleteRoom(c *fiber.Ctx) error
	HowTo(c *fiber.Ctx) error
}

type roomsController struct {
	RoomsService service.RoomsService
}

func NewRoomsController(roomsService service.RoomsService) RoomsController {
	return &roomsController{roomsService}
}

func (ctrl *roomsController) GetAllRooms(c *fiber.Ctx) error {
	rooms, err := ctrl.RoomsService.GetAllRooms()
	if err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, rooms)
	return nil
}

func (ctrl *roomsController) GetRoomByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	room, err := ctrl.RoomsService.GetRoomByID(id)
	if err != nil {
		util.JSONResponse(c, fiber.StatusBadRequest, 4000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, room)
	return nil
}

func (ctrl *roomsController) CreateRoom(c *fiber.Ctx) error {
	var room model.SureSureRoom
	if err := c.BodyParser(&room); err != nil {
		util.JSONResponse(c, fiber.StatusBadRequest, 4000, nil)
		return nil
	}
	id, err := ctrl.RoomsService.CreateRoom(room)
	if err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, fiber.Map{"id": id})
	return nil
}

func (ctrl *roomsController) UpdateRoom(c *fiber.Ctx) error {
	var room model.SureSureRoom
	if err := c.BodyParser(&room); err != nil {
		util.JSONResponse(c, fiber.StatusBadRequest, 4000, nil)
		return nil
	}
	if err := ctrl.RoomsService.UpdateRoom(room); err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, nil)
	return nil
}

func (ctrl *roomsController) DeleteRoom(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := ctrl.RoomsService.DeleteRoom(id); err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, nil)
	return nil
}

func (ctrl *roomsController) HowTo(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user_id := c.Params("user_id")
	err := ctrl.RoomsService.HowTo(id, user_id)
	if err != nil {
		util.JSONResponse(c, fiber.StatusBadRequest, 4000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, id)
	return nil
}
