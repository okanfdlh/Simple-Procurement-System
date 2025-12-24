package controllers

import (
	"backend/config"
	"backend/models"
	"backend/utils"

	"github.com/gofiber/fiber/v2"
)

// Item godoc
// @Summary Get all items
// @Tags Item
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} utils.Response
// @Router /items [get]
func GetItems(c *fiber.Ctx) error {
	var items []models.Item
	config.DB.Find(&items)
	return utils.OK(c, "Get items success", items)
}

// CreateItem godoc
// @Summary Create item
// @Tags Items
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body models.Item true "Item payload"
// @Success 201 {object} models.Item
// @Router /items [post]
// POST /items
func CreateItem(c *fiber.Ctx) error {
	var item models.Item
	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).JSON(err)
	}

	config.DB.Create(&item)
	return utils.Created(c, "Create item success", item)
}

// UpdateItem godoc
// @Summary Update item
// @Tags Items
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Item ID"
// @Param body body models.Item true "Item payload"
// @Success 200 {object} utils.Response
// @Router /items/{id} [put]
// PUT /items/:id
func UpdateItem(c *fiber.Ctx) error {
	id := c.Params("id")
	var item models.Item

	if err := config.DB.First(&item, id).Error; err != nil {
		return utils.NotFound(c, "Item not found")
	}

	c.BodyParser(&item)
	config.DB.Save(&item)

	return utils.OK(c, "Update item success", item)
}

// DeleteItem godoc
// @Summary Delete item
// @Tags Items
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Item ID"
// @Success 200 {object} utils.Response
// @Router /items/{id} [delete]
// DELETE /items/:id
func DeleteItem(c *fiber.Ctx) error {
	id := c.Params("id")
	config.DB.Delete(&models.Item{}, id)
	return utils.OK(c, "Delete item success", fiber.Map{"message": "deleted"})
}
