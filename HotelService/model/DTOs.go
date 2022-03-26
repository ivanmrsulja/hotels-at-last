package model

type HotelDTO struct {
	Id   uint
	Name string
}

type RoomDTO struct {
	Id              uint    `json:"Id"`
	NumberOfBeds    int     `json:"NumberOfBeds"`
	Price           float64 `json:"Price"`
	AirConditioned  bool    `json:"AirConditioned"`
	HasParkingSpace bool    `json:"HasParkingSpace"`
}