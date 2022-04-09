package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ivanmrsulja/hotels-at-last/api-gateway/model"
	"github.com/ivanmrsulja/hotels-at-last/api-gateway/utils"
)

func GetAllReviewsForRoom(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	roomId, _ := strconv.ParseUint(params["id"], 10, 32)
	page := r.URL.Query().Get("page")
	size := r.URL.Query().Get("size")
	
	response, err := http.Get(utils.BaseReviewServicePathRoundRobin.Next().Host + "/api/reviews/rooms/" + strconv.FormatUint(uint64(roomId), 10) + "?page=" + page + "&size=" + size)
	
	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func GetAllReportedReviews(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	if utils.AuthorizeRole(r, "admin") != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	page := r.URL.Query().Get("page")
	size := r.URL.Query().Get("size")
	
	response, err := http.Get(utils.BaseReviewServicePathRoundRobin.Next().Host + "/api/reviews/reported?page=" + page + "&size=" + size)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}
	
	utils.DelegateResponse(response, w)
}

func GetAverageRatingForRoom(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	roomId, _ := strconv.ParseUint(params["id"], 10, 32)
	response, err := http.Get(utils.BaseReviewServicePathRoundRobin.Next().Host + "/api/reviews/rating/" + strconv.FormatUint(uint64(roomId), 10))

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}
	
	utils.DelegateResponse(response, w)
}

func CreateReview(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	if utils.AuthorizeRole(r, "user") != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	data, _ := ioutil.ReadAll(r.Body)
	var reviewRequest model.ReviewCreateRequestDTO
	json.NewDecoder(bytes.NewReader(data)).Decode(&reviewRequest)

	res, err1 := http.Get(utils.BaseBookingServicePathRoundRobin.Next().Host + "/api/reservations/user/" + strconv.FormatUint(uint64(reviewRequest.UserId), 10) + "/room/" + strconv.FormatUint(uint64(reviewRequest.RoomId), 10))
	if err1 != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}
	var num int32
	json.NewDecoder(res.Body).Decode(&num)
	if num == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	req, _ := http.NewRequest(http.MethodPost, utils.BaseReviewServicePathRoundRobin.Next().Host + "/api/reviews", bytes.NewReader(data))
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func ReportReview(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	if utils.AuthorizeRole(r, "user") != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	reviewId, _ := strconv.ParseUint(params["id"], 10, 32)

	req, _ := http.NewRequest(http.MethodPatch, utils.BaseReviewServicePathRoundRobin.Next().Host + "/api/reviews/" + strconv.FormatUint(uint64(reviewId), 10) + "/report", r.Body)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func DismissReviewReports(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	if utils.AuthorizeRole(r, "admin") != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	reviewId, _ := strconv.ParseUint(params["id"], 10, 32)

	req, _ := http.NewRequest(http.MethodPatch, utils.BaseReviewServicePathRoundRobin.Next().Host + "/api/reviews/" + strconv.FormatUint(uint64(reviewId), 10) + "/dismiss", r.Body)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func DeleteReview(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	if utils.AuthorizeRole(r, "admin") != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	params := mux.Vars(r)
	reviewId, _ := strconv.ParseUint(params["id"], 10, 32)

	req, _ := http.NewRequest(http.MethodDelete, utils.BaseReviewServicePathRoundRobin.Next().Host + "/api/reviews/" + strconv.FormatUint(uint64(reviewId), 10), r.Body)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}
