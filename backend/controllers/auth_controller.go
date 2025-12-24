package controllers

import (
	"time"

	"backend/config"
	"backend/dto"
	"backend/models"
	"backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// Register godoc
// @Summary Register new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body dto.RegisterRequest true "Register payload"
// @Success 201 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 409 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /register [post]
func Register(c *fiber.Ctx) error {
	var req dto.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(c, "invalid payload", err.Error())
	}

	// check duplicate username
	var existing models.User
	if err := config.DB.Where("username = ?", req.Username).First(&existing).Error; err == nil {
		return utils.Conflict(c, "username already exists")
	}

	user := models.User{
		Username: req.Username,
		Password: utils.HashPassword(req.Password),
		Role:     req.Role,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return utils.InternalError(c, "failed to register user", err.Error())
	}

	return utils.Created(c, "Register success", nil)
}

// Login godoc
// @Summary Login user
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body dto.LoginRequest true "Login payload"
// @Success 200 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /login [post]
func Login(c *fiber.Ctx) error {
	var req dto.LoginRequest
	var user models.User

	if err := c.BodyParser(&req); err != nil {
		return utils.BadRequest(c, "invalid payload", err.Error())
	}

	if err := config.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return utils.Unauthorized(c, "invalid credentials")
	}

	if !utils.CheckPassword(req.Password, user.Password) {
		return utils.Unauthorized(c, "invalid credentials")
	}

	claims := jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(config.JwtSecret()))

	return utils.OK(c, "Login success", fiber.Map{
		"token": signedToken,
	})
}
