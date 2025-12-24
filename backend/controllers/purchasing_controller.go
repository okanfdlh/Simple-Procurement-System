package controllers

import (
	"backend/config"
	"backend/models"
	"backend/utils"

	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type PurchaseRequest struct {
	SupplierID uint `json:"supplier_id"`
	Items      []struct {
		ItemID uint `json:"item_id"`
		Qty    int  `json:"qty"`
	} `json:"items"`
}

// CreatePurchasing godoc
// @Summary Create purchasing transaction
// @Tags Purchasing
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body PurchaseRequest true "Purchase payload"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /purchasing [post]
func CreatePurchasing(c *fiber.Ctx) error {
	var req PurchaseRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(c, "Invalid payload", err.Error())
	}

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["id"].(float64))

	err := config.DB.Transaction(func(tx *gorm.DB) error {
		p := models.Purchasing{
			Date:       time.Now(), // ✅ FIX UTAMA
			SupplierID: req.SupplierID,
			UserID:     userID,
		}
		if err := tx.Create(&p).Error; err != nil {
			return err
		}

		var grandTotal float64

		for _, i := range req.Items {
			var item models.Item
			if err := tx.First(&item, i.ItemID).Error; err != nil {
				return err
			}

			sub := float64(i.Qty) * item.Price
			grandTotal += sub

			detail := models.PurchasingDetail{
				PurchasingID: p.ID,
				ItemID:       item.ID,
				Qty:          i.Qty,
				SubTotal:     sub,
			}

			if err := tx.Create(&detail).Error; err != nil {
				return err
			}

			tx.Model(&item).Update("stock", item.Stock+i.Qty)
		}

		tx.Model(&p).Update("grand_total", grandTotal)

		go utils.SendWebhook(p.ID, grandTotal)
		return nil
	})

	if err != nil {
		return utils.InternalError(c, "Transaction failed", err.Error())
	}

	return utils.Created(c, "Purchasing created successfully", nil)
}
