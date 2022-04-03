package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	handler "github.com/ivanmrsulja/hotels-at-last/api-gateway/handlers"
)

func HandleRequests() {
	router := mux.NewRouter()

	// Review Service
	router.HandleFunc("/api/reviews/rooms/{id}", handler.GetAllReviewsForRoom).Methods("GET")
	router.HandleFunc("/api/reviews/reported", handler.GetAllReportedReviews).Methods("GET")
	router.HandleFunc("/api/reviews/rating/{id}", handler.GetAverageRatingForRoom).Methods("GET")
	router.HandleFunc("/api/reviews", handler.CreateReview).Methods("POST")
	router.HandleFunc("/api/reviews/{id}/report", handler.ReportReview).Methods("PATCH")
	router.HandleFunc("/api/reviews/{id}/dismiss", handler.DismissReviewReports).Methods("PATCH")
	router.HandleFunc("/api/reviews/{id}", handler.DeleteReview).Methods("DELETE")

	// User Service
	router.HandleFunc("/api/users/login", handler.Login).Methods("POST")
	router.HandleFunc("/api/users/register", handler.Register).Methods("POST")
	router.HandleFunc("/api/users/{id}", handler.GetUser).Methods("GET")
	router.HandleFunc("/api/users/{id}/ban", handler.BanUser).Methods("PATCH")

	// Hotel Service
	router.HandleFunc("/api/hotels", handler.GetAllHotels).Methods("GET")
	router.HandleFunc("/api/hotels/{id}/rooms", handler.GetAllRoomsForHotel).Methods("GET")
	router.HandleFunc("/api/rooms/{id}", handler.GetRoom).Methods("GET")
	router.HandleFunc("/api/hotels", handler.CreateHotel).Methods("POST")
	router.HandleFunc("/api/hotels/{id}/rooms", handler.CreateRoomForHotel).Methods("POST")
	router.HandleFunc("/api/hotels/{id}", handler.UpdateHotel).Methods("PUT")
	router.HandleFunc("/api/rooms/{id}", handler.UpdateRoom).Methods("PUT")
	router.HandleFunc("/api/hotels/{id}", handler.DeleteHotel).Methods("DELETE")
	router.HandleFunc("/api/rooms/{id}", handler.DeleteRoom).Methods("DELETE")

	// Booking Service
	router.HandleFunc("/api/reservations", handler.CreateReservation).Methods("POST")
	router.HandleFunc("/api/reservations/{id}/cancel", handler.CancelReservation).Methods("PUT")
	router.HandleFunc("/api/reservations/{id}", handler.GetAllReservationsForRoom).Methods("GET")
	router.HandleFunc("/api/reservations/user/{id}", handler.GetAllReservationsForUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":8085", router))
}
