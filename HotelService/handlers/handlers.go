package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	model "github.com/ivanmrsulja/hotels-at-last/model"
	repository "github.com/ivanmrsulja/hotels-at-last/repository"
)

func Test(w http.ResponseWriter, r *http.Request) {
	hotel := model.Hotel{Name: "Hotel Lux"}
	var hotelList []model.HotelDTO
	hotelList = append(hotelList, model.HotelDTO{Id: hotel.ID, Name: hotel.Name})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hotelList)
}

func GetAllHotels(w http.ResponseWriter, r *http.Request) {
	hotels := repository.GetAllHotels(r)
	var hotelList []model.HotelDTO
	for _, hotel := range hotels {
		hotelList = append(hotelList, model.HotelDTO{Id: hotel.ID, Name: hotel.Name})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hotelList)
}

func GetAllRoomsForHotel(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)
	rooms := repository.FindRoomsForHotel(r, uint(id))
	var roomList []model.RoomDTO
	for _, room := range rooms {
		roomList = append(roomList, model.RoomDTO{Id: room.ID, NumberOfBeds: room.NumberOfBeds, Price: room.Price, AirConditioned: room.AirConditioned, HasParkingSpace: room.HasParkingSpace})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(roomList)
}

func GetRoom(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	room := repository.FindRoom(uint(id))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.RoomDTO{Id: room.ID, NumberOfBeds: room.NumberOfBeds, Price: room.Price, AirConditioned: room.AirConditioned, HasParkingSpace: room.HasParkingSpace})
}

func CreateRoomForHotel(w http.ResponseWriter, r *http.Request) {

}

func CreateHotel(w http.ResponseWriter, r *http.Request) {

}