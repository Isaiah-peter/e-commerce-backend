package main

import (
	"fmt"
	"github.com/Isaiah-peter/e-commerce-backend/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	port := "Localhost:5000"
	r := mux.NewRouter()
	routes.AuthUser(r)
	routes.UserRoute(r)
	routes.Product(r)
	routes.OrderRoute(r)
	routes.CartRoute(r)
	fmt.Println("server running on Port", port)
	log.Fatal(http.ListenAndServe(port, r))
}
