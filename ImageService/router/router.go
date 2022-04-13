package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	handler "github.com/ivanmrsulja/hotels-at-last/image-service/handlers"
)

func HandleRequests() {
	router := mux.NewRouter()

	router.HandleFunc("/api/images/{path}", handler.GetImage).Methods("GET")
	router.HandleFunc("/api/images/{path}", handler.SaveImage).Methods("POST")

	log.Fatal(http.ListenAndServe(":8084", router))
}