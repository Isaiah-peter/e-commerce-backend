package routes

import (
	"github.com/Isaiah-peter/e-commerce-backend/pkg/controllers"
	"github.com/gorilla/mux"
)

var AuthUser = func(route *mux.Router) {
    route.HandleFunc("/register", controllers.Register).Methods("POST")
	route.HandleFunc("/login", controllers.Login).Methods("POST")
}
