package cart

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/reymandhan/online-store-api/common"
)

type CartItemHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	GetByUsername(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	cartItemService CartItemService
}

func NewCartItemHandler() CartItemHandler {
	return &handler{cartItemService: NewCartItemService()}
}

func (ci *handler) Get(w http.ResponseWriter, r *http.Request) {
	list, err := ci.cartItemService.GetAll()
	if err != nil {
		common.GenerateFailedResponse(w, "Error", err)
	} else {
		common.GenerateOKSuccessResponse(w, "Data Retrieved", list)
	}
}

func (ci *handler) GetByUsername(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	res, err := ci.cartItemService.GetCartByUserName(username)
	if err != nil {
		common.GenerateFailedResponse(w, "Error", err)
	} else {
		common.GenerateOKSuccessResponse(w, "Data Retrieved", res)
	}
}

func (ci *handler) Create(w http.ResponseWriter, r *http.Request) {
	var param AddCartItemRequest

	if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
		common.GenerateFailedResponse(w, "Add to cart failed.", err)
		return
	}

	if err := common.DoValidation(param); err != nil {
		common.GenerateFailedResponse(w, "Add to cart failed.", err)
		return
	}

	result, err := ci.cartItemService.AddToCart(param)
	if err != nil {
		common.GenerateFailedResponse(w, "Add to cart failed.", err)
	} else {
		common.GenerateOKSuccessResponse(w, "Item added to cart successfully.", result)
	}
}

func (ci *handler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	idItem, err := strconv.Atoi(id)

	if err != nil {
		common.GenerateFailedResponse(w, "Failed to delete cart item.", err)
	}

	err = ci.cartItemService.Delete(idItem)
	if err != nil {
		common.GenerateFailedResponse(w, "Failed to delete cart item.", err)
	} else {
		common.GenerateOKSuccessResponse(w, "Cart Item deleted successfully.", nil)
	}
}
