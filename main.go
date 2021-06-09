package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/reymandhan/online-store-api/app/db"
	"github.com/reymandhan/online-store-api/configs"
)

func main() {
	configs.Init()
	db.Init(configs.Global.Database.Host,
		configs.Global.Database.Port,
		configs.Global.Database.Username,
		configs.Global.Database.Password,
		configs.Global.Database.Name,
		configs.Global.Database.SSLMode)

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

	// HealthCheck endpoint
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}).Methods(http.MethodGet)

	http.ListenAndServe(configs.Global.Port, router)
}
