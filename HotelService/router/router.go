package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	handler "github.com/ivanmrsulja/hotels-at-last/handlers"
)

func HandleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/test", handler.Test).Methods("GET")
	router.HandleFunc("/hotels", handler.GetAllHotels).Methods("GET")
	router.HandleFunc("/hotels/{id}/rooms", handler.GetAllRoomsForHotel).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", router))
}