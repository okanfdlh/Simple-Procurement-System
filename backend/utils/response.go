package utils

import "github.com/gofiber/fiber/v2"

type Response struct {
	Success bool        `json:"success" example:"true"`
	Message string      `json:"message" example:"request success"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

//
// ===== SUCCESS RESPONSES =====
//

// 200 OK
func OK(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// 201 Created
func Created(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// 204 No Content
func NoContent(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNoContent)
}

//
// ===== CLIENT ERROR RESPONSES =====
//

// 400 Bad Request
func BadRequest(c *fiber.Ctx, message string, err interface{}) error {
	return c.Status(fiber.StatusBadRequest).JSON(Response{
		Success: false,
		Message: message,
		Error:   err,
	})
}

// 401 Unauthorized
func Unauthorized(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusUnauthorized).JSON(Response{
		Success: false,
		Message: message,
	})
}

// 403 Forbidden
func Forbidden(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusForbidden).JSON(Response{
		Success: false,
		Message: message,
	})
}

// 404 Not Found
func NotFound(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusNotFound).JSON(Response{
		Success: false,
		Message: message,
	})
}

// 409 Conflict
func Conflict(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusConflict).JSON(Response{
		Success: false,
		Message: message,
	})
}

//
// ===== SERVER ERROR RESPONSES =====
//

// 500 Internal Server Error
func InternalError(c *fiber.Ctx, message string, err interface{}) error {
	return c.Status(fiber.StatusInternalServerError).JSON(Response{
		Success: false,
		Message: message,
		Error:   err, // bisa nil di production
	})
}
