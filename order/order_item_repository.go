package order

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	db "github.com/reymandhan/online-store-api/db"
)

type OrderItemRepository struct {
	DB *sqlx.DB `inject:""`
}

func NewOrderItemRepository() *OrderItemRepository {
	return &OrderItemRepository{
		DB: db.Db,
	}
}

const (
	querySelectOrderItem = `SELECT
		* 
		FROM public.order_items`

	queryUpdateOrderItem = `UPDATE order_items
		SET qty = $1,
		updated_at = now() 
		WHERE order_id = $2 AND item_id = $3 RETURNING *`

	queryDeleteOrderItem = `DELETE FROM order_items
		WHERE id = $1`

	queryExistsOrderItem = `SELECT
		COUNT(id) > 0
		FROM public.order_items`
)

func (ci *OrderItemRepository) BeginTx() *sqlx.Tx {
	return ci.DB.MustBegin()
}

func (ci *OrderItemRepository) Rollback(tx *sqlx.Tx) error {
	return tx.Rollback()
}

func (ci *OrderItemRepository) Commit(tx *sqlx.Tx) error {
	err := tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (ci *OrderItemRepository) GetAll() (orderItems []OrderItem, err error) {
	err = ci.DB.Select(&orderItems, querySelectOrderItem)
	return
}

func (ci *OrderItemRepository) GetByOrderID(orderId int) (orderItems []OrderItem, err error) {
	err = ci.DB.Select(&orderItems, querySelectOrderItem+" WHERE order_id = $1 ", orderId)
	return
}

func (ci *OrderItemRepository) Insert(request []AddOrderItemRequest, tx *sqlx.Tx) error {
	queryCreateOrderItem := `INSERT INTO public.order_items
		(order_id, item_id, created_at, updated_at, qty, price)
		VALUES `
	insertparams := []interface{}{}

	for i, data := range request {
		p := i * 4

		queryCreateOrderItem += fmt.Sprintf("($%d,$%d,now(),now(),$%d,$%d),", p+1, p+2, p+3, p+4)

		insertparams = append(insertparams, data.OrderID, data.ItemID, data.Qty, data.Price)
	}

	queryCreateOrderItem = queryCreateOrderItem[:len(queryCreateOrderItem)-1]

	_, err := tx.Exec(queryCreateOrderItem,
		insertparams...)

	return err
}

func (ci *OrderItemRepository) Update(request AddOrderItemRequest, tx *sqlx.Tx) (*OrderItem, error) {
	var result OrderItem

	err := tx.QueryRowx(queryUpdateOrderItem,
		request.Qty,
		request.OrderID,
		request.ItemID).StructScan(&result)

	return &result, err
}

func (ci *OrderItemRepository) Delete(id int) error {
	tx := ci.DB.MustBegin()

	_, err := tx.Exec(queryDeleteOrderItem,
		id)

	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (ci *OrderItemRepository) GetByOrderIDAndItemID(orderID int, itemID int) (orderItem OrderItem, err error) {
	err = ci.DB.Get(&orderItem, querySelectOrderItem+" WHERE order_id = $1 and item_id = $2", orderID, itemID)
	return
}

func (ci *OrderItemRepository) ExistsByID(id int) (exists bool, err error) {
	err = ci.DB.Get(&exists, queryExistsOrderItem+" WHERE id = $1", id)
	return
}
