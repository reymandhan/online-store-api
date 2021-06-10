package item

import "time"

type Item struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Name      string    `json:"name"`
	SKU       string    `json:"sku"`
	Price     float32   `json:"price"`
	Qty       int       `json:"qty"`
}

type NewItemRequest struct {
	Name  string  `json:"name" validate:"required"`
	SKU   string  `json:"sku" validate:"required"`
	Price float32 `json:"price" validate:"required,numeric"`
	Qty   int     `json:"qty" validate:"required,numeric"`
}

type UpdateItemRequest struct {
	ID    int
	Name  string  `json:"name" validate:"required"`
	SKU   string  `json:"sku" validate:"required"`
	Price float32 `json:"price" validate:"required,numeric"`
	Qty   int     `json:"qty" validate:"required,numeric"`
}
