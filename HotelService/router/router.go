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
	router.HandleFunc("/rooms/{id}", handler.GetRoom).Methods("GET")
	
	router.HandleFunc("/hotels", handler.CreateHotel).Methods("POST")
	router.HandleFunc("/hotels/{id}/rooms", handler.CreateRoomForHotel).Methods("POST")

	router.HandleFunc("/hotels/{id}", handler.UpdateHotel).Methods("PUT")
	router.HandleFunc("/rooms/{id}", handler.UpdateRoom).Methods("PUT")

	router.HandleFunc("/hotels/{id}", handler.DeleteHotel).Methods("DELETE")
	router.HandleFunc("/rooms/{id}", handler.DeleteRoom).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", router))
}