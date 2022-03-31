package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	model "github.com/ivanmrsulja/hotels-at-last/review-service/model"
	repository "github.com/ivanmrsulja/hotels-at-last/review-service/repository"
)

func CreateReview(w http.ResponseWriter, r *http.Request) {
	var reviewDTO model.ReviewCreateRequestDTO
	json.NewDecoder(r.Body).Decode(&reviewDTO)
	
	review := reviewDTO.ToReview()

	createdReview, err := repository.CreateReview(review)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	} else {
		json.NewEncoder(w).Encode(createdReview.ToDTO())
	}
}

func ReportReview(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	reviewId, _ := strconv.ParseUint(params["id"], 10, 32)

	err := repository.ReportReview(uint(reviewId))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func DismissReviewReports(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	reviewId, _ := strconv.ParseUint(params["id"], 10, 32)

	err := repository.DismissReviewReports(uint(reviewId))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func GetAllReportedReviews(w http.ResponseWriter, r *http.Request) {
	reviews := repository.GetAllReportedReviews(r)
	reviewList := []model.ReviewDTO{}
	for _, review := range reviews {
		reviewList = append(reviewList, review.ToDTO())
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reviewList)
}

func GetAllReviewsForRoom(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	roomId, _ := strconv.ParseUint(params["id"], 10, 32)

	reviews := repository.GetAllReviewsForRoom(r, uint(roomId))
	reviewList := []model.ReviewDTO{}
	for _, review := range reviews {
		reviewList = append(reviewList, review.ToDTO())
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reviewList)
}

func DeleteReview(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	reviewId, _ := strconv.ParseUint(params["id"], 10, 32)

	err := repository.DeleteReview(uint(reviewId))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func GetAverageRatingForRoom(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	roomId, _ := strconv.ParseUint(params["id"], 10, 32)

	result := repository.GetAverageGradeForRoom(uint(roomId))

	json.NewEncoder(w).Encode(result)
}
