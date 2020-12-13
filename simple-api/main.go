package main

import (
	"alvinr.ca/learn-go/api/controllers"
	"alvinr.ca/learn-go/api/data"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Entry point
func main() {
	fmt.Println("Hello, world!")

	router := mux.NewRouter()
	data.PopulateDatastore()

	router.HandleFunc("/api/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/api/books", controllers.CreateBooks).Methods("POST")
	router.HandleFunc("/api/books/{id}", controllers.UpdateBooks).Methods("PUT")
	router.HandleFunc("/api/books/{id}", controllers.DeleteBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
