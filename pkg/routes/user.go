package routes

import (
	"github.com/Isaiah-peter/e-commerce-backend/pkg/controllers"
	"github.com/gorilla/mux"
)

var UserRoute = func(route *mux.Router) {
	route.HandleFunc("/user/{id}", controllers.GetUserById).Methods("GET")
	route.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")
	route.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	route.HandleFunc("/user", controllers.GetAllUser).Methods("GET")
	route.HandleFunc("/stat", controllers.UserStat).Methods("GET")
	route.HandleFunc("/user", controllers.GetUserUsername).Methods("GET")
	route.HandleFunc("/payment", controllers.CreateCharge).Methods("POST")
}
