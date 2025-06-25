package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/textures1245/payso-check-slip-backend/model"
	"github.com/textures1245/payso-check-slip-backend/service"
	"github.com/textures1245/payso-check-slip-backend/util"
)

type PackageController interface {
	GetPackageAll(c *fiber.Ctx) error
	GetPackageByID(c *fiber.Ctx) error
	CreatePackage(c *fiber.Ctx) error
	UpdatePackage(c *fiber.Ctx) error
	DeletePackage(c *fiber.Ctx) error
}

type packageController struct {
	PackageService service.PackageService
}

func NewPackageController(packageService service.PackageService) PackageController {
	return &packageController{packageService}
}

func (ctrl *packageController) GetPackageAll(c *fiber.Ctx) error {
	packages, err := ctrl.PackageService.GetPackageAll()
	if err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, err.Error())
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, packages)
	return nil
}

func (ctrl *packageController) GetPackageByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	pkg, err := ctrl.PackageService.GetPackageByID(id)
	if err != nil {
		util.JSONResponse(c, fiber.StatusBadRequest, 4000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, pkg)
	return nil
}

func (ctrl *packageController) CreatePackage(c *fiber.Ctx) error {
	var pkg model.SureSurePackage
	if err := c.BodyParser(&pkg); err != nil {
		util.JSONResponse(c, fiber.StatusBadRequest, 4000, err.Error())
		return nil
	}
	id, err := ctrl.PackageService.CreatePackage(pkg)
	if err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, err.Error())
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, fiber.Map{"id": id})
	return nil
}

func (ctrl *packageController) UpdatePackage(c *fiber.Ctx) error {
	var pkg model.SureSurePackage
	if err := c.BodyParser(&pkg); err != nil {
		util.JSONResponse(c, fiber.StatusBadRequest, 4000, err.Error())
		return nil
	}
	if err := ctrl.PackageService.UpdatePackage(pkg); err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, err.Error())
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, nil)
	return nil
}

func (ctrl *packageController) DeletePackage(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := ctrl.PackageService.DeletePackage(id); err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, nil)
	return nil
}
