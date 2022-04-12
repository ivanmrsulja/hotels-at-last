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
	router.HandleFunc("/api/reviews/rooms/{id}", handler.GetAllReviewsForRoom).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/reviews/reported", handler.GetAllReportedReviews).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/reviews/rating/{id}", handler.GetAverageRatingForRoom).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/reviews", handler.CreateReview).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/reviews/{id}/report", handler.ReportReview).Methods("PATCH", "OPTIONS")
	router.HandleFunc("/api/reviews/{id}/dismiss", handler.DismissReviewReports).Methods("PATCH", "OPTIONS")
	router.HandleFunc("/api/reviews/{id}", handler.DeleteReview).Methods("DELETE", "OPTIONS")

	// User Service
	router.HandleFunc("/api/users/login", handler.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/users/register", handler.Register).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/users/{id}", handler.GetUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/users/{id}/ban", handler.BanUser).Methods("PATCH", "OPTIONS")

	// Hotel Service
	router.HandleFunc("/api/hotels", handler.GetAllHotels).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/hotels/{id}/rooms", handler.GetAllRoomsForHotel).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/rooms/{id}", handler.GetRoom).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/hotels/{id}", handler.GetHotel).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/hotels", handler.CreateHotel).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/hotels/{id}/rooms", handler.CreateRoomForHotel).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/hotels/{id}", handler.UpdateHotel).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/rooms/{id}", handler.UpdateRoom).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/hotels/{id}", handler.DeleteHotel).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/rooms/{id}", handler.DeleteRoom).Methods("DELETE", "OPTIONS")

	router.HandleFunc("/api/images/{path}", handler.GetImage).Methods("GET", "OPTIONS")

	// Booking Service
	router.HandleFunc("/api/reservations", handler.CreateReservation).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/reservations/{id}/cancel", handler.CancelReservation).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/reservations/{id}", handler.GetAllReservationsForRoom).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/reservations/user/{id}", handler.GetAllReservationsForUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/reservations/user/{userId}/room/{roomId}",handler.GetCountForUserAndRoom).Methods("GET", "OPTIONS")

	log.Fatal(http.ListenAndServe(":8085", router))
}
