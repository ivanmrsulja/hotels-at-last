package handlers

import (
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ivanmrsulja/hotels-at-last/api-gateway/utils"
)

func GetAllReviewsForRoom(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	roomId, _ := strconv.ParseUint(params["id"], 10, 32)
	page := r.URL.Query().Get("page")
	size := r.URL.Query().Get("size")
	
	response, _ := http.Get("http://localhost:8082/api/reviews/rooms/" + strconv.FormatUint(uint64(roomId), 10) + "?page=" + page + "&size=" + size)
	
	w.Header().Set("Content-Type", response.Header.Get("Content-Type"))
    w.Header().Set("Content-Length", response.Header.Get("Content-Length"))
	w.WriteHeader(response.StatusCode)
    io.Copy(w, response.Body)
    response.Body.Close()
}

func GetAllReportedReviews(w http.ResponseWriter, r *http.Request) {
	if utils.AuthorizeRole(r, "admin") != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	page := r.URL.Query().Get("page")
	size := r.URL.Query().Get("size")
	
	response, _ := http.Get("http://localhost:8082/api/reviews/reported?page=" + page + "&size=" + size)
	
	w.Header().Set("Content-Type", response.Header.Get("Content-Type"))
    w.Header().Set("Content-Length", response.Header.Get("Content-Length"))
	w.WriteHeader(response.StatusCode)
    io.Copy(w, response.Body)
    response.Body.Close()
}

func GetAverageRatingForRoom(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	roomId, _ := strconv.ParseUint(params["id"], 10, 32)
	response, _ := http.Get("http://localhost:8082/api/reviews/rating/" + strconv.FormatUint(uint64(roomId), 10))
	
	w.Header().Set("Content-Type", response.Header.Get("Content-Type"))
    w.Header().Set("Content-Length", response.Header.Get("Content-Length"))
	w.WriteHeader(response.StatusCode)
    io.Copy(w, response.Body)
    response.Body.Close()
}

func CreateReview(w http.ResponseWriter, r *http.Request) {
	if utils.AuthorizeRole(r, "user") != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8082/api/reviews", r.Body)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	response, _ := client.Do(req)

	w.Header().Set("Content-Type", response.Header.Get("Content-Type"))
    w.Header().Set("Content-Length", response.Header.Get("Content-Length"))
	w.WriteHeader(response.StatusCode)
    io.Copy(w, response.Body)
    response.Body.Close()
}

func ReportReview(w http.ResponseWriter, r *http.Request) {
	if utils.AuthorizeRole(r, "user") != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	reviewId, _ := strconv.ParseUint(params["id"], 10, 32)

	req, _ := http.NewRequest(http.MethodPatch, "http://localhost:8082/api/reviews/" + strconv.FormatUint(uint64(reviewId), 10) + "/report", r.Body)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	response, _ := client.Do(req)

	w.Header().Set("Content-Type", response.Header.Get("Content-Type"))
    w.Header().Set("Content-Length", response.Header.Get("Content-Length"))
	w.WriteHeader(response.StatusCode)
    io.Copy(w, response.Body)
    response.Body.Close()
}

func DismissReviewReports(w http.ResponseWriter, r *http.Request) {
	if utils.AuthorizeRole(r, "admin") != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	reviewId, _ := strconv.ParseUint(params["id"], 10, 32)

	req, _ := http.NewRequest(http.MethodPatch, "http://localhost:8082/api/reviews/" + strconv.FormatUint(uint64(reviewId), 10) + "/dismiss", r.Body)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	response, _ := client.Do(req)

	w.Header().Set("Content-Type", response.Header.Get("Content-Type"))
    w.Header().Set("Content-Length", response.Header.Get("Content-Length"))
	w.WriteHeader(response.StatusCode)
    io.Copy(w, response.Body)
    response.Body.Close()
}

func DeleteReview(w http.ResponseWriter, r *http.Request) {
	if utils.AuthorizeRole(r, "admin") != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	reviewId, _ := strconv.ParseUint(params["id"], 10, 32)

	req, _ := http.NewRequest(http.MethodDelete, "http://localhost:8082/api/reviews/" + strconv.FormatUint(uint64(reviewId), 10), r.Body)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	response, _ := client.Do(req)

	w.Header().Set("Content-Type", response.Header.Get("Content-Type"))
    w.Header().Set("Content-Length", response.Header.Get("Content-Length"))
	w.WriteHeader(response.StatusCode)
    io.Copy(w, response.Body)
    response.Body.Close()
}
