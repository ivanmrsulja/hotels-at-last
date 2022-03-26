package model

import (
	"github.com/jinzhu/gorm"
)


type Room struct {
	gorm.Model

	NumberOfBeds int
	Price float64
	AirConditioned bool
	HasParkingSpace bool
	HotelID int
}

func (room *Room) ToDTO() RoomDTO {
	return RoomDTO{Id: room.ID, NumberOfBeds: room.NumberOfBeds, Price: room.Price, AirConditioned: room.AirConditioned, HasParkingSpace: room.HasParkingSpace}
}


type Hotel struct {
	gorm.Model

	Name string
	Rooms []Room
}

func (hotel *Hotel) ToDTO() HotelDTO{
	return HotelDTO{hotel.ID, hotel.Name}
}
