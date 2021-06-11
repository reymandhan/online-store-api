package order

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/reymandhan/online-store-api/common"
)

type OrderHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Pay(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	orderService OrderService
}

func NewOrderHandler() OrderHandler {
	return &handler{orderService: NewOrderService()}
}

func (o *handler) Create(w http.ResponseWriter, r *http.Request) {
	var param CreateOrderRequest

	if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
		common.GenerateFailedResponse(w, "Checkout failed.", err)
		return
	}

	if err := common.DoValidation(param); err != nil {
		common.GenerateFailedResponse(w, "Checkout failed.", err)
		return
	}

	err := o.orderService.Checkout(param)
	if err != nil {
		common.GenerateFailedResponse(w, "Checkout failed.", err)
	} else {
		common.GenerateOKSuccessResponse(w, "Checkout success.", nil)
	}
}

func (o *handler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	idOrder, err := strconv.Atoi(id)

	res, err := o.orderService.GetByID(idOrder)
	if err != nil {
		common.GenerateFailedResponse(w, "Error", err)
	} else {
		common.GenerateOKSuccessResponse(w, "Data Retrieved", res)
	}
}

func (o *handler) Pay(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	idOrder, err := strconv.Atoi(id)

	err = o.orderService.Pay(idOrder)
	if err != nil {
		common.GenerateFailedResponse(w, "Payment failed.", err)
	} else {
		common.GenerateOKSuccessResponse(w, "Payment success.", nil)
	}
}
