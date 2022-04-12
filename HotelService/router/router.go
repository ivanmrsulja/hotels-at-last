package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	handler "github.com/ivanmrsulja/hotels-at-last/handlers"
)

func HandleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/api/test", handler.Test).Methods("GET")
	router.HandleFunc("/api/hotels", handler.GetAllHotels).Methods("GET")
	router.HandleFunc("/api/hotels/{id}/rooms", handler.GetAllRoomsForHotel).Methods("GET")
	router.HandleFunc("/api/rooms/{id}", handler.GetRoom).Methods("GET")
	router.HandleFunc("/api/hotels/{id}", handler.GetHotel).Methods("GET")

	router.HandleFunc("/api/images/{path}", handler.GetImage).Methods("GET")
	
	router.HandleFunc("/api/hotels", handler.CreateHotel).Methods("POST")
	router.HandleFunc("/api/hotels/{id}/rooms", handler.CreateRoomForHotel).Methods("POST")

	router.HandleFunc("/api/hotels/{id}", handler.UpdateHotel).Methods("PUT")
	router.HandleFunc("/api/rooms/{id}", handler.UpdateRoom).Methods("PUT")

	router.HandleFunc("/api/hotels/{id}", handler.DeleteHotel).Methods("DELETE")
	router.HandleFunc("/api/rooms/{id}", handler.DeleteRoom).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", router))
}