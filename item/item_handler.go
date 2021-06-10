package item

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/reymandhan/online-store-api/common"
)

type ItemHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	itemService ItemService
}

func NewItemHandler() ItemHandler {
	return &handler{itemService: NewItemService()}
}

func (i *handler) Get(w http.ResponseWriter, r *http.Request) {
	list, err := i.itemService.GetAll()
	if err != nil {
		common.GenerateFailedResponse(w, "Error", err)
	} else {
		common.GenerateOKSuccessResponse(w, "Data Retrieved", list)
	}
}

func (i *handler) Create(w http.ResponseWriter, r *http.Request) {
	var param NewItemRequest

	if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
		common.GenerateFailedResponse(w, "Failed to create item.", err)
		return
	}

	if err := common.DoValidation(param); err != nil {
		common.GenerateFailedResponse(w, "Failed to create item.", err)
		return
	}

	result, err := i.itemService.Create(param)
	if err != nil {
		common.GenerateFailedResponse(w, "Failed to create item.", err)
	} else {
		common.GenerateOKSuccessResponse(w, "Item created successfully.", result)
	}
}

func (i *handler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	idItem, err := strconv.Atoi(id)

	var param UpdateItemRequest

	if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
		common.GenerateFailedResponse(w, "Failed to update item.", err)
		return
	}

	if err := common.DoValidation(param); err != nil {
		common.GenerateFailedResponse(w, "Failed to update item.", err)
		return
	}

	param.ID = idItem
	result, err := i.itemService.Update(param)
	if err != nil {
		common.GenerateFailedResponse(w, "Failed to update item.", err)
	} else {
		common.GenerateOKSuccessResponse(w, "Item Updated successfully.", result)
	}
}

func (i *handler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	idItem, err := strconv.Atoi(id)

	if err != nil {
		common.GenerateFailedResponse(w, "Failed to delete item.", err)
	}

	err = i.itemService.Delete(idItem)
	if err != nil {
		common.GenerateFailedResponse(w, "Failed to delete item.", err)
	} else {
		common.GenerateOKSuccessResponse(w, "Item deleted successfully.", nil)
	}
}
