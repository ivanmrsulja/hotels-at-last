package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	handler "github.com/ivanmrsulja/hotels-at-last/user-service/handlers"
)

func HandleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/api/users/login", handler.Login).Methods("POST")

	router.HandleFunc("/api/users/authorize/admin", handler.AuthoriseAdmin).Methods("GET")
	router.HandleFunc("/api/users/authorize/user", handler.AuthoriseUser).Methods("GET")
	router.HandleFunc("/api/users/{id}", handler.FindUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":8083", router))
}