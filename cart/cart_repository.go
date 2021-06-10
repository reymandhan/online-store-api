package cart

import (
	"github.com/jmoiron/sqlx"
	db "github.com/reymandhan/online-store-api/db"
)

type CartRepository struct {
	DB *sqlx.DB `inject:""`
}

func NewCartRepository() *CartRepository {
	return &CartRepository{
		DB: db.Db,
	}
}

const (
	queryCreateCart = `INSERT INTO public.carts
		(created_at, updated_at, username, total_price)
		VALUES(now(), now(), $1, $2) RETURNING *;`

	querySelectCart = `SELECT 
		*
		FROM public.carts `

	queryUpdateCart = `UPDATE carts
		SET total_price = $1,
		updated_at = now() 
		WHERE username = $2 RETURNING *;`
)

func (c *CartRepository) BeginTx() *sqlx.Tx {
	return c.DB.MustBegin()
}

func (c *CartRepository) Rollback(tx *sqlx.Tx) error {
	return tx.Rollback()
}

func (c *CartRepository) Commit(tx *sqlx.Tx) error {
	err := tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (c *CartRepository) Insert(username string, total float32, tx *sqlx.Tx) (*Cart, error) {
	var result Cart
	err := tx.QueryRowx(queryCreateCart,
		username,
		total).StructScan(&result)

	return &result, err
}

func (c *CartRepository) GetByUsername(username string) (cart Cart, err error) {
	err = c.DB.Get(&cart, querySelectCart+" WHERE username = $1", username)
	return
}

func (c *CartRepository) Update(username string, total float32, tx *sqlx.Tx) (*Cart, error) {
	var result Cart
	err := tx.QueryRowx(queryUpdateCart,
		total,
		username).StructScan(&result)

	return &result, err
}
