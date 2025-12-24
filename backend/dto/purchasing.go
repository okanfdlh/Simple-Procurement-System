package dto

type PurchaseItemRequest struct {
	ItemID uint `json:"item_id" example:"1"`
	Qty    int  `json:"qty" example:"2"`
}

type PurchaseRequest struct {
	SupplierID uint                  `json:"supplier_id" example:"1"`
	Items      []PurchaseItemRequest `json:"items"`
}
