package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	handler "github.com/ivanmrsulja/hotels-at-last/api-gateway/handlers"
)

func HandleRequests() {
	router := mux.NewRouter()
	router.HandleFunc("/api/reviews/rooms/{id}", handler.GetAllReviewsForRoom).Methods("GET")
	router.HandleFunc("/api/reviews/reported", handler.GetAllReportedReviews).Methods("GET")
	router.HandleFunc("/api/reviews/rating/{id}", handler.GetAverageRatingForRoom).Methods("GET")

	router.HandleFunc("/api/reviews", handler.CreateReview).Methods("POST")

	router.HandleFunc("/api/reviews/{id}/report", handler.ReportReview).Methods("PATCH")
	router.HandleFunc("/api/reviews/{id}/dismiss", handler.DismissReviewReports).Methods("PATCH")

	router.HandleFunc("/api/reviews/{id}", handler.DeleteReview).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8085", router))
}