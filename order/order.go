package order

import "time"

type Order struct {
	ID         int       `json:"id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
	Username   string    `json:"username"`
	Address    string    `json:"address"`
	TotalPrice float32   `json:"total_price" db:"total_price"`
	Status     string    `json:"status"`
}

type CreateOrderRequest struct {
	Username   string `json:"username"`
	Address    string `json:"address"`
	TotalPrice float32
	Status     string
}

type UserOrder struct {
	ID         int         `json:"id"`
	Username   string      `json:"username"`
	Address    string      `json:"address"`
	TotalPrice float32     `json:"total_price" db:"total_price"`
	Status     string      `json:"status"`
	OrderItems []OrderItem `json:"order_items"`
}
