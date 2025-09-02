package controller

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/textures1245/payso-check-slip-backend/model"
	"github.com/textures1245/payso-check-slip-backend/service"
	"github.com/textures1245/payso-check-slip-backend/util"
)

type UserController interface {
	GetUserAll(c *fiber.Ctx) error
	GetUserByID(c *fiber.Ctx) error
	GetOrCreateUser(c *fiber.Ctx) error
	CreateUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
	GetCategoryAll(c *fiber.Ctx) error
}

type userController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{userService}
}

// GetUserAll godoc
// @Summary List users
// @Tags User
// @Produce json
// @Success 200 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /user/get [get]
func (ctrl *userController) GetUserAll(c *fiber.Ctx) error {
	users, err := ctrl.UserService.GetUserAll()
	if err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, users)
	return nil
}

// GetUserByID godoc
// @Summary Get user by UID
// @Description Retrieve a user record by UID
// @Tags User
// @Produce json
// @Param id path string true "User UID"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /user/get/{id} [get]
func (ctrl *userController) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := ctrl.UserService.GetUserByID(id)
	if err != nil {
		util.JSONResponse(c, fiber.StatusBadRequest, 4000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, user)
	return nil
}

// GetOrCreateUser godoc
// @Summary Login or create user
// @Description Create user if not exists, otherwise return existing user
// @Tags User
// @Accept json
// @Produce json
// @Param user body model.SureSureUser true "User payload"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /login [post]
func (ctrl *userController) GetOrCreateUser(c *fiber.Ctx) error {
	var user model.SureSureUser

	// ล็อกการเริ่มต้นฟังก์ชัน
	log.Println("Received request to get or create user ")

	// อ่านข้อมูลจาก body ของ request
	if err := c.BodyParser(&user); err != nil {
		// ล็อกเมื่อเกิดข้อผิดพลาด
		log.Printf("Error parsing body: %v", err)
		util.JSONResponse(c, fiber.StatusBadRequest, 4000, nil)
		return nil
	}

	// ล็อกข้อมูล user ที่ได้รับจาก request
	log.Printf("Received user: %+v", user)

	// เรียกใช้ service เพื่อตรวจสอบหรือสร้าง user
	user, err := ctrl.UserService.GetOrCreateUser(user)
	if err != nil {
		// ล็อกเมื่อเกิดข้อผิดพลาดจาก service
		log.Printf("Error in GetOrCreateUser service: %v", err)
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, nil)
		return nil
	}

	// ล็อกข้อมูล user ที่ได้จาก service
	log.Printf("User processed: %+v", user)

	// ส่ง response กลับไป
	util.JSONResponse(c, fiber.StatusOK, 2006, user)

	// ล็อกการเสร็จสิ้นการทำงาน
	log.Println("User successfully created or fetched")

	return nil
}

// CreateUser godoc
// @Summary Create user
// @Tags User
// @Accept json
// @Produce json
// @Param user body model.SureSureUser true "User payload"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /user/create [post]
func (ctrl *userController) CreateUser(c *fiber.Ctx) error {
	var user model.SureSureUser
	if err := c.BodyParser(&user); err != nil {
		util.JSONResponse(c, fiber.StatusBadRequest, 4000, nil)
		return nil
	}
	id, err := ctrl.UserService.CreateUser(user)
	log.Infof("id: %s", id)
	if err != nil {
		if err.Error() == "duplicate store" {
			util.JSONResponse(c, fiber.StatusBadRequest, 4004, nil)
			return nil
		}
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, nil)
	return nil
}

// UpdateUser godoc
// @Summary Update user
// @Tags User
// @Accept json
// @Produce json
// @Param user body model.SureSureUser true "User payload"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /user/update [put]
func (ctrl *userController) UpdateUser(c *fiber.Ctx) error {
	var user model.SureSureUser
	if err := c.BodyParser(&user); err != nil {
		util.JSONResponse(c, fiber.StatusBadRequest, 4000, nil)
		return nil
	}
	if err := ctrl.UserService.UpdateUser(user); err != nil {
		if err.Error() == "duplicate store" {
			util.JSONResponse(c, fiber.StatusBadRequest, 4004, nil)
			return nil
		}
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, nil)
	return nil
}

// DeleteUser godoc
// @Summary Delete user by UID
// @Tags User
// @Produce json
// @Param id path string true "User UID"
// @Success 200 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /user/delete/{id} [delete]
func (ctrl *userController) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := ctrl.UserService.DeleteUser(id); err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, nil)
	return nil
}

// GetCategoryAll godoc
// @Summary List categories
// @Description Retrieve merchant categories
// @Tags User
// @Produce json
// @Success 200 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /user/category/get [get]
func (ctrl *userController) GetCategoryAll(c *fiber.Ctx) error {
	users, err := ctrl.UserService.GetCategoryAll()
	if err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, users)
	return nil
}
