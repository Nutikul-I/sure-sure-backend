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

// GetPackageAll godoc
// @Summary List packages
// @Tags Package
// @Produce json
// @Success 200 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /package/get [get]
func (ctrl *packageController) GetPackageAll(c *fiber.Ctx) error {
	packages, err := ctrl.PackageService.GetPackageAll()
	if err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, err.Error())
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, packages)
	return nil
}

// GetPackageByID godoc
// @Summary Get package by ID
// @Tags Package
// @Produce json
// @Param id path int true "Package ID"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /package/get/{id} [get]
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

// CreatePackage godoc
// @Summary Create package
// @Tags Package
// @Accept json
// @Produce json
// @Param package body model.SureSurePackage true "Package payload"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /package/create [post]
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

// UpdatePackage godoc
// @Summary Update package
// @Tags Package
// @Accept json
// @Produce json
// @Param package body model.SureSurePackage true "Package payload"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /package/update [put]
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

// DeletePackage godoc
// @Summary Delete package by ID
// @Tags Package
// @Produce json
// @Param id path int true "Package ID"
// @Success 200 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /package/delete/{id} [delete]
func (ctrl *packageController) DeletePackage(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := ctrl.PackageService.DeletePackage(id); err != nil {
		util.JSONResponse(c, fiber.StatusInternalServerError, 5000, nil)
		return nil
	}
	util.JSONResponse(c, fiber.StatusOK, 2006, nil)
	return nil
}
