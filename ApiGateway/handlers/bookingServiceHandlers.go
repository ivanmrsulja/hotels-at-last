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

func GetAllReservationsForRoom(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	roomId, _ := strconv.ParseUint(params["id"], 10, 32)
	response, err := http.Get(utils.BaseBookingServicePathRoundRobin.Next().Host + "/api/reservations/" + strconv.FormatUint(uint64(roomId), 10))

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}
	
	utils.DelegateResponse(response, w)
}

func GetAllReservationsForUser(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	userId, _ := strconv.ParseUint(params["id"], 10, 32)
	response, err := http.Get(utils.BaseBookingServicePathRoundRobin.Next().Host + "/api/reservations/user/" + strconv.FormatUint(uint64(userId), 10))

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}
	
	utils.DelegateResponse(response, w)
}

func CancelReservation(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	reservationId, _ := strconv.ParseUint(params["id"], 10, 32)

	req, _ := http.NewRequest(http.MethodPut, utils.BaseBookingServicePathRoundRobin.Next().Host + "/api/reservations/" + strconv.FormatUint(uint64(reservationId), 10) + "/cancel" , r.Body)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func CreateReservation(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	var request model.CreateReservationRequest
	data, _ := ioutil.ReadAll(r.Body)
	json.NewDecoder(bytes.NewReader(data)).Decode(&request)

	responseUser, errUser := http.Get(utils.BaseUserServicePathRoundRobin.Next().Host + "/api/users/" + strconv.FormatUint(uint64(request.UserId), 10))

	if errUser != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}
	
	if responseUser.StatusCode != 200 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	responseRoom, errRoom := http.Get(utils.BaseHotelServicePathRoundRobin.Next().Host + "/api/rooms/" + strconv.FormatUint(uint64(request.RoomId), 10))

	if responseRoom.StatusCode != 200 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if errRoom != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	req, _ := http.NewRequest(http.MethodPost, utils.BaseBookingServicePathRoundRobin.Next().Host + "/api/reservations/" , bytes.NewReader(data))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}