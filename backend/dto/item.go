package dto

type CreateItemRequest struct {
	Name  string  `json:"name" example:"Laptop"`
	Stock int     `json:"stock" example:"10"`
	Price float64 `json:"price" example:"15000000"`
}
