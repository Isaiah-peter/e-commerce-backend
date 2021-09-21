package routes

import (
	"github.com/Isaiah-peter/e-commerce-backend/pkg/controllers"
	"github.com/gorilla/mux"
)

var CartRoute = func(route *mux.Router) {
	route.HandleFunc("/cart", controllers.CreateCart).Methods("POST")
	route.HandleFunc("/cart", controllers.GetCartById).Methods("GET")
	route.HandleFunc("/getallcart", controllers.GetCart).Methods("GET")
	route.HandleFunc("/cart/{id}", controllers.UpdateCart).Methods("PUT")
	route.HandleFunc("/cart/{id}", controllers.DeleteCart).Methods("DELETE")
}

