package item

import (
	"log"

	"github.com/jmoiron/sqlx"
	db "github.com/reymandhan/online-store-api/db"
)

type ItemRepository struct {
	DB *sqlx.DB `inject:""`
}

func NewItemRepository() *ItemRepository {
	return &ItemRepository{
		DB: db.Db,
	}
}

const (
	querySelectItem = `SELECT 
		id,
		name,
		sku,
		price,
		qty,
		created_at,
		updated_at
	FROM public.items `

	queryInsertItem = `INSERT INTO public.items (
		created_at, 
		updated_at, 
		name, 
		sku, 
		price, 
		qty) 
	VALUES(now(), now(), $1, $2, $3, $4) RETURNING *`

	queryExistItem = `SELECT
			COUNT(id) > 0
		FROM public.items`

	queryUpdateItem = `UPDATE public.items 
		SET
			updated_at = now(),
			name = $1, 
			sku = $2,
			price = $3,
			qty= $4 
		WHERE id = $5 RETURNING *`

	queryDeleteItem = `UPDATE public.items 
		SET
			deleted_at = now()
		WHERE id = $1`

	queryUpdateQtyItem = `UPDATE public.items 
		SET
			updated_at = now(),
			qty= $1 
		WHERE id = $2 `
)

func (r *ItemRepository) GetAll() (items []Item, err error) {
	err = r.DB.Select(&items, querySelectItem+" WHERE deleted_at is null")
	return
}

func (r *ItemRepository) Insert(item NewItemRequest) (*Item, error) {
	tx := r.DB.MustBegin()

	var result Item
	err := tx.QueryRowx(queryInsertItem,
		item.Name,
		item.SKU,
		item.Price,
		item.Qty).StructScan(&result)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	log.Print(&result)
	return &result, nil
}

func (r *ItemRepository) Update(item UpdateItemRequest) (*Item, error) {
	tx := r.DB.MustBegin()

	var result Item
	err := tx.QueryRowx(queryUpdateItem,
		item.Name,
		item.SKU,
		item.Price,
		item.Qty,
		item.ID).StructScan(&result)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *ItemRepository) Delete(id int) error {
	tx := r.DB.MustBegin()

	_, err := tx.Exec(queryDeleteItem, id)

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

func (r *ItemRepository) GetByID(id int) (item Item, err error) {
	err = r.DB.Get(&item, querySelectItem+" WHERE id = $1 and deleted_at is null", id)
	return
}

func (r *ItemRepository) ExistsBySKU(sku string) (exists bool, err error) {
	err = r.DB.Get(&exists, queryExistItem+" WHERE sku = $1 and deleted_at is null", sku)
	return
}

func (r *ItemRepository) ExistsDuplicateSKUByID(sku string, id int) (exists bool, err error) {
	err = r.DB.Get(&exists, queryExistItem+" WHERE sku = $1 AND id <> $2 and deleted_at is null", sku, id)
	return
}

func (r *ItemRepository) ExistsByID(id int) (exists bool, err error) {
	err = r.DB.Get(&exists, queryExistItem+" WHERE id = $1 and deleted_at is null", id)
	return
}

func (r *ItemRepository) GetByIDWithLock(id int, tx *sqlx.Tx) (item Item, err error) {
	err = tx.Get(&item, querySelectItem+" WHERE id = $1 and deleted_at is null FOR UPDATE", id)
	return
}

func (r *ItemRepository) UpdateWithLock(id int, qty int, tx *sqlx.Tx) error {

	_, err := tx.Exec(queryUpdateQtyItem,
		qty,
		id)

	return err
}
