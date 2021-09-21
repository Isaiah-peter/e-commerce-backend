package routes

import (
	"github.com/Isaiah-peter/e-commerce-backend/pkg/controllers"
	"github.com/gorilla/mux"
)

var OrderRoute = func(route *mux.Router) {
	route.HandleFunc("/order", controllers.CreateOrder).Methods("POST")
	route.HandleFunc("/order", controllers.GetOrderById).Methods("GET")
	route.HandleFunc("/getallorder", controllers.GetOrder).Methods("GET")
	route.HandleFunc("/order/{id}", controllers.UpdateOrder).Methods("PUT")
	route.HandleFunc("/order/{id}", controllers.DeleteOrder).Methods("DELETE")
}
