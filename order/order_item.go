package order

import "time"

type OrderItem struct {
	ID        int       `json:"id"`
	OrderID   int       `json:"order_id" db:"order_id"`
	ItemID    int       `json:"item_id" db:"item_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Qty       int       `json:"qty"`
	Price     float32   `json:"price"`
}

type AddOrderItemRequest struct {
	ItemID  int     `json:"item_id" validated:"required"`
	Qty     int     `json:"qty" validated:"required"`
	Price   float32 `json:"price" validated:"required"`
	OrderID int
}
