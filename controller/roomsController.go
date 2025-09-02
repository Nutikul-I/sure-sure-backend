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

// GetAllRooms godoc
// @Summary List rooms
// @Tags Room
// @Produce json
// @Success 200 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /room2/get [get]
func (ctrl *roomsController) GetAllRooms(c *fiber.Ctx) error {
	rooms, err := ctrl.RoomsService.GetAllRooms()
	if err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, rooms)
	return nil
}

// GetRoomByID godoc
// @Summary Get room by ID
// @Tags Room
// @Produce json
// @Param id path int true "Room ID"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /room2/get/{id} [get]
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

// CreateRoom godoc
// @Summary Create room
// @Tags Room
// @Accept json
// @Produce json
// @Param room body model.SureSureRoom true "Room payload"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /room2/create [post]
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

// UpdateRoom godoc
// @Summary Update room
// @Tags Room
// @Accept json
// @Produce json
// @Param room body model.SureSureRoom true "Room payload"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /room2/update [put]
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

// DeleteRoom godoc
// @Summary Delete room by ID
// @Tags Room
// @Produce json
// @Param id path int true "Room ID"
// @Success 200 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /room2/delete/{id} [delete]
func (ctrl *roomsController) DeleteRoom(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := ctrl.RoomsService.DeleteRoom(id); err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, nil)
	return nil
}

// HowTo godoc
// @Summary HowTo action for room
// @Tags Room
// @Produce json
// @Param id path int true "Room ID"
// @Param user_id path string true "User ID"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Router /room2/howto/{id}/{user_id} [get]
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
