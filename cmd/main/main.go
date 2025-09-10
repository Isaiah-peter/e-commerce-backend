package main

import (
	"fmt"
	"github.com/Isaiah-peter/e-commerce-backend/pkg/routes"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	r := mux.NewRouter()
	subRouter := r.PathPrefix("/api/v1").Subrouter()
	routes.AuthUser(subRouter)
	routes.UserRoute(subRouter)
	routes.Product(subRouter)
	routes.OrderRoute(subRouter)
	routes.CartRoute(subRouter)
	address := fmt.Sprintf("%s:%s", "localhost", port)
	fmt.Println("server running on Port", address)
	log.Fatal(http.ListenAndServe(address, handlers.CORS(handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}), handlers.AllowedOrigins([]string{"*"}))(r)))

}
