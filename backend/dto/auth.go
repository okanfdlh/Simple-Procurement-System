package dto

type RegisterRequest struct {
	Username string `json:"username" example:"admin"`
	Password string `json:"password" example:"password123"`
	Role     string `json:"role" example:"admin"`
}

type LoginRequest struct {
	Username string `json:"username" example:"admin"`
	Password string `json:"password" example:"password123"`
}
