package cart

import "time"

type Cart struct {
	ID         int       `json:"id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	Username   string    `json:"username"`
	TotalPrice float32   `json:"total_price" db:"total_price"`
}

type UserCart struct {
	ID         int        `json:"id"`
	Username   string     `json:"username"`
	TotalPrice float32    `json:"total_price" db:"total_price"`
	CartItems  []CartItem `json:"cart_items"`
}
