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
func (ctrl *userController) GetUserAll(c *fiber.Ctx) error {
	users, err := ctrl.UserService.GetUserAll()
	if err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, users)
	return nil
}

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

func (ctrl *userController) CreateUser(c *fiber.Ctx) error {
	var user model.SureSureUser
	if err := c.BodyParser(&user); err != nil {
		util.JSONResponse(c, fiber.StatusBadRequest, 4000, nil)
		return nil
	}
	id, err := ctrl.UserService.CreateUser(user)
	log.Infof("id: %d", id)
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

func (ctrl *userController) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := ctrl.UserService.DeleteUser(id); err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, nil)
	return nil
}

func (ctrl *userController) GetCategoryAll(c *fiber.Ctx) error {
	users, err := ctrl.UserService.GetCategoryAll()
	if err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, users)
	return nil
}
