package controllers

import (
	"backend/config"
	"backend/models"
	"backend/utils"

	"github.com/gofiber/fiber/v2"
)

// Supplier godoc
// @Summary Get all suppliers
// @Tags Supplier
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} utils.Response
// @Router /suppliers [get]
func GetSuppliers(c *fiber.Ctx) error {
	var suppliers []models.Supplier
	config.DB.Find(&suppliers)
	return utils.OK(c, "Get suppliers success", suppliers)
}

// CreateSupplier godoc
// @Summary Create supplier
// @Tags Supplier
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body models.Supplier true "Supplier payload"
// @Success 201 {object} models.Supplier
// @Router /suppliers [post]
func CreateSupplier(c *fiber.Ctx) error {
	var supplier models.Supplier
	if err := c.BodyParser(&supplier); err != nil {
		return c.Status(400).JSON(err)
	}

	config.DB.Create(&supplier)
	return utils.Created(c, "Create supplier success", supplier)
}

// UpdateSupplier godoc
// @Summary Update supplier
// @Tags Supplier
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Supplier ID"
// @Param body body models.Supplier true "Supplier payload"
// @Success 200 {object} models.Supplier
// @Router /suppliers/{id} [put]
// PUT /suppliers/:id
func UpdateSupplier(c *fiber.Ctx) error {
	id := c.Params("id")
	var supplier models.Supplier

	if err := config.DB.First(&supplier, id).Error; err != nil {
		return utils.NotFound(c, "Supplier not found")
	}

	c.BodyParser(&supplier)
	config.DB.Save(&supplier)

	return utils.OK(c, "Update supplier success", supplier)
}

// DeleteSupplier godoc
// @Summary Delete supplier
// @Tags Supplier
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Supplier ID"
// @Success 200 {object} utils.Response
// @Router /suppliers/{id} [delete]
// DELETE /suppliers/:id
func DeleteSupplier(c *fiber.Ctx) error {
	id := c.Params("id")
	config.DB.Delete(&models.Supplier{}, id)
	return utils.OK(c, "Delete supplier success", fiber.Map{"message": "deleted"})
}
