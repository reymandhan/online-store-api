package cart

import (
	"github.com/jmoiron/sqlx"
	db "github.com/reymandhan/online-store-api/db"
)

type CartItemRepository struct {
	DB *sqlx.DB `inject:""`
}

func NewCartItemRepository() *CartItemRepository {
	return &CartItemRepository{
		DB: db.Db,
	}
}

const (
	queryCreateCartItem = `INSERT INTO public.cart_items
		(cart_id, item_id, created_at, updated_at, qty, price)
		VALUES($1, $2, now(), now(), $3, $4) RETURNING *`

	querySelectCartItem = `SELECT
		* 
		FROM public.cart_items`

	queryUpdateCartItem = `UPDATE cart_items
		SET qty = $1,
		updated_at = now() 
		WHERE cart_id = $2 AND item_id = $3 RETURNING *`

	queryDeleteCartItem = `DELETE FROM cart_items `

	queryExistsCartItem = `SELECT
		COUNT(id) > 0
		FROM public.cart_items`
)

func (ci *CartItemRepository) BeginTx() *sqlx.Tx {
	return ci.DB.MustBegin()
}

func (ci *CartItemRepository) Rollback(tx *sqlx.Tx) error {
	return tx.Rollback()
}

func (ci *CartItemRepository) Commit(tx *sqlx.Tx) error {
	err := tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (ci *CartItemRepository) GetAll() (cartItems []CartItem, err error) {
	err = ci.DB.Select(&cartItems, querySelectCartItem)
	return
}

func (ci *CartItemRepository) Insert(request AddCartItemRequest, tx *sqlx.Tx) (*CartItem, error) {
	var result CartItem

	err := tx.QueryRowx(queryCreateCartItem,
		request.CartID,
		request.ItemID,
		request.Qty,
		request.Price).StructScan(&result)

	return &result, err
}

func (ci *CartItemRepository) Update(request AddCartItemRequest, tx *sqlx.Tx) (*CartItem, error) {
	var result CartItem

	err := tx.QueryRowx(queryUpdateCartItem,
		request.Qty,
		request.CartID,
		request.ItemID).StructScan(&result)

	return &result, err
}

func (ci *CartItemRepository) Delete(id int) error {
	tx := ci.DB.MustBegin()

	_, err := tx.Exec(queryDeleteCartItem+" WHERE id = $1",
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

func (ci *CartItemRepository) DeleteByCartId(id int, tx *sqlx.Tx) error {
	_, err := tx.Exec(queryDeleteCartItem+" WHERE cart_id = $1",
		id)
	return err
}

func (ci *CartItemRepository) GetByCartIDAndItemID(cartID int, itemID int) (cartItem CartItem, err error) {
	err = ci.DB.Get(&cartItem, querySelectCartItem+" WHERE cart_id = $1 and item_id = $2", cartID, itemID)
	return
}

func (ci *CartItemRepository) GetByCartID(cartID int) (cartItem []CartItem, err error) {
	err = ci.DB.Select(&cartItem, querySelectCartItem+" WHERE cart_id = $1", cartID)
	return
}

func (ci *CartItemRepository) ExistsByID(id int) (exists bool, err error) {
	err = ci.DB.Get(&exists, queryExistsCartItem+" WHERE id = $1", id)
	return
}
