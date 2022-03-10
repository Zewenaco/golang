package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"test-api/pkg/db"
	"test-api/pkg/handlers"
)

const ROUTE_BOOKS = "/book"
const ROUTE_BOOKS_BY_ID = ROUTE_BOOKS + "/{id}"

func main() {
	DB := db.Init()
	handlers := handlers.New(DB)
	router := mux.NewRouter()

	// Here we'll define our api endpoints
	router.HandleFunc(ROUTE_BOOKS, handlers.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc(ROUTE_BOOKS, handlers.AddBook).Methods(http.MethodPost)
	router.HandleFunc(ROUTE_BOOKS_BY_ID, handlers.GetBook).Methods(http.MethodGet)
	router.HandleFunc(ROUTE_BOOKS_BY_ID, handlers.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc(ROUTE_BOOKS_BY_ID, handlers.DeleteBook).Methods(http.MethodDelete)

	log.Println("API is running under port :8080")
	http.ListenAndServe(":8080", router)
}
