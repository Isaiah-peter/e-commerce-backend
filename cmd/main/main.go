package main

import (
	"fmt"
	"github.com/Isaiah-peter/e-commerce-backend/pkg/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	os.Setenv("PORT", "5000")
	port := os.Getenv("PORT")
	r := mux.NewRouter()
	routes.AuthUser(r)
	routes.UserRoute(r)
	routes.Product(r)
	routes.OrderRoute(r)
	routes.CartRoute(r)
	address := fmt.Sprintf("%s:%s", "localhost", port)
	fmt.Println("server running on Port", address)
	log.Fatal(http.ListenAndServe(address, r))
}
