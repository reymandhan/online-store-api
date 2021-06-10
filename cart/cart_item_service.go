package cart

import (
	"errors"
	"fmt"

	"github.com/reymandhan/online-store-api/item"
)

type CartItemService interface {
	GetAll() ([]CartItem, error)
	GetCartByUserName(username string) (UserCart, error)
	AddToCart(request AddCartItemRequest) (*CartItem, error)
	Delete(id int) (err error)
}

type cartItemService struct {
	cartRepository     *CartRepository
	cartItemRepository *CartItemRepository
	itemRepository     *item.ItemRepository
}

func NewCartItemService() CartItemService {
	return &cartItemService{
		cartRepository:     NewCartRepository(),
		cartItemRepository: NewCartItemRepository(),
		itemRepository:     item.NewItemRepository(),
	}
}

func (ci *cartItemService) GetAll() ([]CartItem, error) {
	items, err := ci.cartItemRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (ci *cartItemService) GetCartByUserName(username string) (UserCart, error) {
	var result UserCart
	cart, err := ci.cartRepository.GetByUsername(username)

	if err != nil {
		return result, err
	}

	cartItems, err := ci.cartItemRepository.GetByCartID(cart.ID)
	if err != nil {
		return result, err
	}

	result.CartItems = cartItems
	result.ID = cart.ID
	result.Username = cart.Username
	result.TotalPrice = cart.TotalPrice

	return result, nil
}

func (ci *cartItemService) AddToCart(request AddCartItemRequest) (result *CartItem, err error) {
	item, err := ci.itemRepository.GetByID(request.ItemID)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Item with id: %v is not exists", request.ItemID))
	}

	// check if user has active cart
	cart, err := ci.cartRepository.GetByUsername(request.Username)
	tx := ci.cartRepository.BeginTx()

	if cart.ID == 0 {
		// Create new cart
		newCart, err := ci.cartRepository.Insert(request.Username, item.Price*float32(request.Qty), tx)
		if err != nil {
			return nil, err
		}
		request.CartID = newCart.ID
	} else {
		request.CartID = cart.ID
		// update total price of existing cart
		_, err = ci.cartRepository.Update(request.Username, cart.TotalPrice+(item.Price*float32(request.Qty)), tx)
		if err != nil {
			return nil, err
		}
	}

	var cartItem *CartItem
	// check if selected item is already exists on active cart
	cItem, err := ci.cartItemRepository.GetByCartIDAndItemID(cart.ID, request.ItemID)

	//validate stock availability
	if item.Qty < (request.Qty + cItem.Qty) {
		// rollback transaction and return error message
		ci.cartRepository.Rollback(tx)
		return nil, errors.New(fmt.Sprintf("Cannot add %v item (total: %v) to cart, current stock: %v", request.Qty, request.Qty+cItem.Qty, item.Qty))
	}

	if cItem.ID > 0 {
		// update qty on cart
		request.Qty += cItem.Qty
		cartItem, err = ci.cartItemRepository.Update(request, tx)
	} else {
		// insert new cart item
		cartItem, err = ci.cartItemRepository.Insert(request, tx)
	}

	if err != nil {
		ci.cartRepository.Rollback(tx)
		return nil, err
	}
	err = ci.cartRepository.Commit(tx)
	if err != nil {
		return nil, err
	}
	return cartItem, err
}

func (ci *cartItemService) Delete(id int) (err error) {
	// check if data exists
	exists, err := ci.cartItemRepository.ExistsByID(id)
	if !exists {
		return errors.New(fmt.Sprintf("CartItem with id: %v is not exists", id))
	}

	err = ci.cartItemRepository.Delete(id)
	return err
}
