package cart

import "time"

type CartItem struct {
	ID        int       `json:"id"`
	CartID    int       `json:"cart_id" db:"cart_id"`
	ItemID    int       `json:"item_id" db:"item_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Qty       int       `json:"qty"`
	Price     float32   `json:"price"`
}

type AddCartItemRequest struct {
	Username string `json:"username" validated:"required"`
	ItemID   int    `json:"item_id" validated:"required"`
	Qty      int    `json:"qty" validated:"required"`
	Price    float32
	CartID   int
}
