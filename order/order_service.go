package order

import (
	"errors"
	"fmt"
	"log"

	"github.com/reymandhan/online-store-api/cart"
	"github.com/reymandhan/online-store-api/item"
)

type OrderService interface {
	Checkout(request CreateOrderRequest) (err error)
	GetByID(id int) (userOrder UserOrder, err error)
	Pay(id int) error
}

type orderService struct {
	orderRepository     *OrderRepository
	orderItemRepository *OrderItemRepository
	cartRepository      *cart.CartRepository
	cartItemRepository  *cart.CartItemRepository
	itemRepository      *item.ItemRepository
}

func NewOrderService() OrderService {
	return &orderService{
		orderRepository:     NewOrderRepository(),
		orderItemRepository: NewOrderItemRepository(),
		cartRepository:      cart.NewCartRepository(),
		cartItemRepository:  cart.NewCartItemRepository(),
		itemRepository:      item.NewItemRepository(),
	}
}

func (o *orderService) Checkout(request CreateOrderRequest) (err error) {
	cart, err := o.cartRepository.GetByUsername(request.Username)
	if err != nil {
		return errors.New(fmt.Sprintf("Cart not exists for user: %v", request.Username))
	}

	cartItems, err := o.cartItemRepository.GetByCartID(cart.ID)

	request.Status = "CHECKOUT"
	request.TotalPrice = cart.TotalPrice

	tx := o.orderRepository.BeginTx()
	order, err := o.orderRepository.Insert(request, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	var orderItems []AddOrderItemRequest
	for _, val := range cartItems {
		var oi AddOrderItemRequest
		oi.OrderID = order.ID
		oi.ItemID = val.ItemID
		oi.Qty = val.Qty
		oi.Price = val.Price

		orderItems = append(orderItems, oi)
	}
	err = o.orderItemRepository.Insert(orderItems, tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return
}

func (o *orderService) GetByID(id int) (UserOrder, error) {
	var result UserOrder
	order, err := o.orderRepository.GetByID(id)

	if err != nil {
		return result, errors.New(fmt.Sprintf("Order by id: %v is not exists", id))
	}

	orderItems, err := o.orderItemRepository.GetByOrderID(order.ID)

	result.ID = order.ID
	result.Username = order.Username
	result.Address = order.Address
	result.TotalPrice = order.TotalPrice
	result.Status = order.Status
	result.OrderItems = orderItems
	return result, err
}

func (o *orderService) Pay(id int) error {
	order, err := o.orderRepository.GetByID(id)

	if err != nil {
		return errors.New(fmt.Sprintf("Order by id: %v is not exists", id))
	}

	if order.Status == "PAID" {
		return errors.New(fmt.Sprintf("Order by id: %v has been paid", id))
	}

	orderItems, err := o.orderItemRepository.GetByOrderID(order.ID)

	// start transaction and lock to be updated row
	tx := o.orderRepository.BeginTx()

	for _, orderItem := range orderItems {
		item, err := o.itemRepository.GetByIDWithLock(orderItem.ItemID, tx)
		log.Print(item)
		if err != nil {
			log.Print(err)
			tx.Rollback()
			return err
		}
		if item.Qty < orderItem.Qty {
			tx.Rollback()
			return errors.New(fmt.Sprintf("Payment failed, Stock insufficient for item %v.", item.Name))
		} else {
			o.itemRepository.UpdateWithLock(item.ID, (item.Qty - orderItem.Qty), tx)
		}
	}
	// Update order status
	log.Print("Update Order status")
	_, err = o.orderRepository.Update(id, "PAID", tx)
	if err != nil {
		log.Print(err)
		tx.Rollback()
		return err
	}

	// Clean up cart and cart item for the user
	log.Print("Clean cart")
	cart, err := o.cartRepository.GetByUsername(order.Username)
	err = o.cartItemRepository.DeleteByCartId(cart.ID, tx)
	if err != nil {
		log.Print(err)
		tx.Rollback()
		return err
	}
	err = o.cartRepository.DeleteById(cart.ID, tx)
	if err != nil {
		log.Print(err)
		tx.Rollback()
		return err
	}

	tx.Commit()
	log.Print("Success")

	return err
}
