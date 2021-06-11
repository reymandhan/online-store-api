package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/reymandhan/online-store-api/cart"
	validator "github.com/reymandhan/online-store-api/common"
	"github.com/reymandhan/online-store-api/configs"
	"github.com/reymandhan/online-store-api/db"
	"github.com/reymandhan/online-store-api/item"
	"github.com/reymandhan/online-store-api/order"
)

func main() {
	configs.Init()
	db.Init(configs.Global.Database.Host,
		configs.Global.Database.Port,
		configs.Global.Database.Username,
		configs.Global.Database.Password,
		configs.Global.Database.Name,
		configs.Global.Database.SSLMode)

	db.Migrate(configs.Global.Database.Host,
		configs.Global.Database.Port,
		configs.Global.Database.Username,
		configs.Global.Database.Password,
		configs.Global.Database.Name,
		configs.Global.Database.SSLMode)

	validator.Init()

	router := mux.NewRouter().
		PathPrefix(configs.Global.APIPrefix).
		Subrouter()

	// Set json contentType
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	// Health check endpoint
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"app": true, "db": db.Ping()})
	}).Methods(http.MethodGet)

	itemHandler := item.NewItemHandler()
	router.HandleFunc("/item", itemHandler.Get).Methods(http.MethodGet)
	router.HandleFunc("/item", itemHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/item/{id}", itemHandler.Update).Methods(http.MethodPut)
	router.HandleFunc("/item/{id}", itemHandler.Delete).Methods(http.MethodDelete)

	cartItemHandler := cart.NewCartItemHandler()
	router.HandleFunc("/cart/item", cartItemHandler.Get).Methods(http.MethodGet)
	router.HandleFunc("/cart/add", cartItemHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/cart/item/{id}", cartItemHandler.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/cart", cartItemHandler.GetByUsername).Methods(http.MethodGet)

	orderHandler := order.NewOrderHandler()
	router.HandleFunc("/order/checkout", orderHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/order/detail/{id}", orderHandler.GetByID).Methods(http.MethodGet)
	router.HandleFunc("/order/pay/{id}", orderHandler.Pay).Methods(http.MethodPut)

	http.ListenAndServe(configs.Global.Port, router)
}
