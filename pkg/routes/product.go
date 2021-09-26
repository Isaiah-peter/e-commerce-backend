package routes

import (
	"github.com/Isaiah-peter/e-commerce-backend/pkg/controllers"
	"github.com/gorilla/mux"
)

var Product = func(route *mux.Router) {
	route.HandleFunc("/product", controllers.CreateProduct).Methods("POST")
	route.HandleFunc("/category", controllers.CreateCategory).Methods("POST")
	route.HandleFunc("/product", controllers.GetProduct).Methods("GET")
	route.HandleFunc("/product/{id}", controllers.GetProductById).Methods("GET")
	route.HandleFunc("/product/{id}", controllers.UpdateProduct).Methods("PUT")
	route.HandleFunc("/product/{id}", controllers.DeleteProduct).Methods("DELETE")

}
