package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	port := "Localhost:5000"
	r := mux.NewRouter()
	fmt.Println("server running on Port", port)
	log.Fatal(http.ListenAndServe(port, r))
}
