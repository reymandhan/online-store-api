package order

import (
	"github.com/jmoiron/sqlx"
	db "github.com/reymandhan/online-store-api/db"
)

type OrderRepository struct {
	DB *sqlx.DB `inject:""`
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		DB: db.Db,
	}
}

const (
	queryCreateOrder = `INSERT INTO public.orders
		(created_at, updated_at, username, total_price, address, status)
		VALUES(now(), now(), $1, $2, $3, $4) RETURNING *;`

	querySelectOrder = `SELECT 
		*
		FROM public.orders `

	queryUpdateOrder = `UPDATE orders
		SET status = $1,
		updated_at = now() 
		WHERE id = $2 RETURNING *;`
)

func (c *OrderRepository) BeginTx() *sqlx.Tx {
	return c.DB.MustBegin()
}

func (c *OrderRepository) Rollback(tx *sqlx.Tx) error {
	return tx.Rollback()
}

func (c *OrderRepository) Commit(tx *sqlx.Tx) error {
	err := tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (c *OrderRepository) Insert(request CreateOrderRequest, tx *sqlx.Tx) (*Order, error) {
	var result Order
	err := tx.QueryRowx(queryCreateOrder,
		request.Username,
		request.TotalPrice,
		request.Address,
		request.Status).StructScan(&result)

	return &result, err
}

func (c *OrderRepository) GetByID(id int) (order Order, err error) {
	err = c.DB.Get(&order, querySelectOrder+" WHERE id = $1", id)
	return
}

func (c *OrderRepository) GetByUsername(username string) (order Order, err error) {
	err = c.DB.Get(&order, querySelectOrder+" WHERE username = $1", username)
	return
}

func (c *OrderRepository) Update(id int, status string, tx *sqlx.Tx) (*Order, error) {
	var result Order
	err := tx.QueryRowx(queryUpdateOrder,
		status,
		id).StructScan(&result)

	return &result, err
}
