package routes

import (
	"github.com/Isaiah-peter/e-commerce-backend/pkg/controllers"
	"github.com/gorilla/mux"
)

var AuthUser = func(route *mux.Router) {
    route.HandleFunc("/register", controllers.Register).Methods("POST")
	route.HandleFunc("/login", controllers.Login).Methods("POST")
	route.HandleFunc("/auth/google/login", controllers.GoogleLogin).Methods("GET")
	route.HandleFunc("/auth/google/callback", controllers.GoogleCallback).Methods("GET")
}
