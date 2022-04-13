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
	hotels, totalResults := repository.GetAllHotels(r)
	hotelList := []model.HotelDTO{}
	for _, hotel := range hotels {
		hotelList = append(hotelList, hotel.ToDTO())
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.HotelsPageable{Results: hotelList, TotalResults: totalResults})
}

func GetAllRoomsForHotel(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)
	rooms, totalResults := repository.FindRoomsForHotel(r, uint(id))
	roomList := []model.RoomDTO{}
	for _, room := range rooms {
		roomList = append(roomList, room.ToDTO())
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.RoomsPageable{Results: roomList, TotalResults: totalResults})
}

func GetRoom(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)
	room, err := repository.FindRoom(uint(id))
	
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	} else {
		json.NewEncoder(w).Encode(room.ToDTO())
	}
}

func GetHotel(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)
	hotel, err := repository.FindHotel(uint(id))
	
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	} else {
		json.NewEncoder(w).Encode(hotel.ToDTO())
	}
}

func CreateRoomForHotel(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	hotelId, _ := strconv.ParseUint(params["id"], 10, 32)
	var roomDTO model.RoomDTO
	json.NewDecoder(r.Body).Decode(&roomDTO)
	
	room := roomDTO.ToRoom()
	room.HotelID = uint(hotelId)

	createdRoom, err := repository.CreateRoom(room)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	} else {
		json.NewEncoder(w).Encode(createdRoom.ToDTO())
	}
}

func CreateHotel(w http.ResponseWriter, r *http.Request) {
	var hotelDTO model.HotelDTO
	json.NewDecoder(r.Body).Decode(&hotelDTO)
	hotel := hotelDTO.ToHotel()
	createdHotel, err := repository.CreateHotel(hotel)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	} else {
		json.NewEncoder(w).Encode(createdHotel.ToDTO())
	}
}

func GetImage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	imagePath, _ := params["path"]

	base64Image := repository.GetBase64Image(imagePath)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.ImageResponse{Base64Image: base64Image})
}

func UpdateHotel(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	hotelId, _ := strconv.ParseUint(params["id"], 10, 32)
	var hotelDTO model.HotelDTO
	json.NewDecoder(r.Body).Decode(&hotelDTO)
	hotel := hotelDTO.ToHotel()

	err := repository.UpdateHotel(hotel, uint(hotelId))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func UpdateRoom(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	roomId, _ := strconv.ParseUint(params["id"], 10, 32)
	var roomDTO model.RoomDTO
	json.NewDecoder(r.Body).Decode(&roomDTO)
	room := roomDTO.ToRoom()

	err := repository.UpdateRoom(room, uint(roomId))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func DeleteHotel(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	hotelId, _ := strconv.ParseUint(params["id"], 10, 32)

	err := repository.DeleteHotel(uint(hotelId))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func DeleteRoom(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	roomId, _ := strconv.ParseUint(params["id"], 10, 32)

	err := repository.DeleteRoom(uint(roomId))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	} else {
		w.WriteHeader(http.StatusNoContent)
	}	
}