package item

import (
	"errors"
	"fmt"
)

type ItemService interface {
	GetAll() ([]Item, error)
	Create(NewItemRequest) (*Item, error)
	Update(UpdateItemRequest) (*Item, error)
	Delete(id int) (err error)
}

type itemService struct {
	itemRepository *ItemRepository
}

func NewItemService() ItemService {
	return &itemService{itemRepository: NewItemRepository()}
}

func (i *itemService) GetAll() ([]Item, error) {
	items, err := i.itemRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (i *itemService) Create(request NewItemRequest) (result *Item, err error) {
	// check if SKU already exists
	exists, err := i.itemRepository.ExistsBySKU(request.SKU)
	if exists {
		return nil, errors.New(fmt.Sprintf("Item with SKU: %v is already exists", request.SKU))
	}

	item, err := i.itemRepository.Insert(request)
	return item, err
}

func (i *itemService) Update(request UpdateItemRequest) (item *Item, err error) {
	// check if data exists
	exists, err := i.itemRepository.ExistsByID(request.ID)
	if !exists {
		return nil, errors.New(fmt.Sprintf("Item with id: %v is not exists", request.ID))
	}

	// Check if new SKU is exists
	exists, err = i.itemRepository.ExistsDuplicateSKUByID(request.SKU, request.ID)
	if exists {
		return nil, errors.New(fmt.Sprintf("SKU: %v is already exists", request.SKU))
	}

	item, err = i.itemRepository.Update(request)
	return item, err
}

func (i *itemService) Delete(id int) (err error) {
	// check if data exists
	exists, err := i.itemRepository.ExistsByID(id)
	if !exists {
		return errors.New(fmt.Sprintf("Item with id: %v is not exists", id))
	}

	err = i.itemRepository.Delete(id)
	return err
}
