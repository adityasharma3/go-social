package router

import (
	"github.com/adityasharma3/go-social/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/stock/{id}", middleware.GetStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/stocks", middleware.GetAllStocks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/new-stock", middleware.CreateStock).Methods("POST", "OPTIONS")
	// router.HandleFunc("/api/stock", middleware.UpdateStock).Methods("PUT", "OPTIONS")
	// router.HandleFunc("/api/stock/{id}", middleware.DeleteStock).Methods("DELETE", "OPTIONS")

	return router
}
