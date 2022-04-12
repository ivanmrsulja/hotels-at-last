package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ivanmrsulja/hotels-at-last/api-gateway/utils"
)

func GetAllHotels(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	page := r.URL.Query().Get("page")
	size := r.URL.Query().Get("size")
	bedsFrom := r.URL.Query().Get("bedsFrom")
	bedsTo := r.URL.Query().Get("bedsTo")
	priceFrom := r.URL.Query().Get("priceFrom")
	priceTo := r.URL.Query().Get("priceTo")
	airCond := r.URL.Query().Get("airCond")
	parking := r.URL.Query().Get("parking")
	tv := r.URL.Query().Get("tv")
	name := r.URL.Query().Get("name")
	address := r.URL.Query().Get("address")
	starsFrom := r.URL.Query().Get("starsFrom")
	starsTo := r.URL.Query().Get("starsTo")

	response, err := http.Get(utils.BaseHotelServicePathRoundRobin.Next().Host + "/api/hotels?page=" + page + "&size=" + size + "&bedsFrom=" + bedsFrom + "&bedsTo=" + bedsTo + "&priceFrom=" + priceFrom + "&priceTo=" + priceTo + "&airCond=" + airCond + "&parking=" + parking + "&tv=" + tv + "&name=" + name + "&address=" + address + "&starsFrom=" + starsFrom + "&starsTo=" + starsTo)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}
	
	utils.DelegateResponse(response, w)
}

func GetAllRoomsForHotel(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	hotelId, _ := strconv.ParseUint(params["id"], 10, 32)
	page := r.URL.Query().Get("page")
	size := r.URL.Query().Get("size")

	response, err := http.Get(utils.BaseHotelServicePathRoundRobin.Next().Host + "/api/hotels/" + strconv.FormatUint(uint64(hotelId), 10) + "/rooms?page=" + page + "&size=" + size)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}
	
	utils.DelegateResponse(response, w)
}

func GetRoom(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	roomId, _ := strconv.ParseUint(params["id"], 10, 32)

	response, err := http.Get(utils.BaseHotelServicePathRoundRobin.Next().Host + "/api/rooms/" + strconv.FormatUint(uint64(roomId), 10))

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}
	
	utils.DelegateResponse(response, w)
}

func GetImage(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	image, _ := params["path"]

	response, err := http.Get(utils.BaseHotelServicePathRoundRobin.Next().Host + "/api/images/" + image)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}
	
	utils.DelegateResponse(response, w)
}

func GetHotel(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	hotelId, _ := strconv.ParseUint(params["id"], 10, 32)

	response, err := http.Get(utils.BaseHotelServicePathRoundRobin.Next().Host + "/api/hotels/" + strconv.FormatUint(uint64(hotelId), 10))

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}
	
	utils.DelegateResponse(response, w)
}

func CreateHotel(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	req, _ := http.NewRequest(http.MethodPost, utils.BaseHotelServicePathRoundRobin.Next().Host + "/api/hotels", r.Body)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func CreateRoomForHotel(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	hotelId, _ := strconv.ParseUint(params["id"], 10, 32)

	req, _ := http.NewRequest(http.MethodPost, utils.BaseHotelServicePathRoundRobin.Next().Host + "/api/hotels/" + strconv.FormatUint(uint64(hotelId), 10) + "/rooms", r.Body)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func UpdateHotel(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	hotelId, _ := strconv.ParseUint(params["id"], 10, 32)

	req, _ := http.NewRequest(http.MethodPut, utils.BaseHotelServicePathRoundRobin.Next().Host + "/api/hotels/" + strconv.FormatUint(uint64(hotelId), 10) , r.Body)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func UpdateRoom(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	roomId, _ := strconv.ParseUint(params["id"], 10, 32)

	req, _ := http.NewRequest(http.MethodPut, utils.BaseHotelServicePathRoundRobin.Next().Host + "/api/rooms/" + strconv.FormatUint(uint64(roomId), 10) , r.Body)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func DeleteHotel(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	hotelId, _ := strconv.ParseUint(params["id"], 10, 32)

	req, _ := http.NewRequest(http.MethodDelete, utils.BaseHotelServicePathRoundRobin.Next().Host + "/api/hotels/" + strconv.FormatUint(uint64(hotelId), 10), r.Body)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}

func DeleteRoom(w http.ResponseWriter, r *http.Request) {
	utils.SetupResponse(&w, r)
	if r.Method == "OPTIONS" {
		return
	}

	params := mux.Vars(r)
	roomId, _ := strconv.ParseUint(params["id"], 10, 32)

	req, _ := http.NewRequest(http.MethodDelete, utils.BaseHotelServicePathRoundRobin.Next().Host + "/api/rooms/" + strconv.FormatUint(uint64(roomId), 10), r.Body)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		w.WriteHeader(http.StatusGatewayTimeout)
		return
	}

	utils.DelegateResponse(response, w)
}
