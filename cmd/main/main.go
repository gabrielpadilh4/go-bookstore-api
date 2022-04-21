package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gabrielpadilha/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/", router)

	fmt.Printf("Server started on port 8000")

	log.Fatal(http.ListenAndServe(":8000", router))
}
