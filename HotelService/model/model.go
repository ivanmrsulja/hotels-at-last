package model

import (
	"github.com/jinzhu/gorm"
)


type Room struct {
	gorm.Model

	NumberOfBeds int `gorm:"minimum(1)"`
	Price float64 `gorm:"min:0.0"`
	AirConditioned bool
	HasParkingSpace bool
	HasTV bool
	HotelID uint
}

func (room *Room) ToDTO() RoomDTO {
	return RoomDTO{Id: room.ID, NumberOfBeds: room.NumberOfBeds, Price: room.Price, AirConditioned: room.AirConditioned, HasParkingSpace: room.HasParkingSpace, HasTV: room.HasTV}
}


type Hotel struct {
	gorm.Model

	Name string `gorm:"not null;default:null"`
	Address string `gorm:"not null;default:null"`
	Base64Image string `gorm:"not null;default:null"`
	Rooms []Room
}

func (hotel *Hotel) ToDTO() HotelDTO{
	return HotelDTO{hotel.ID, hotel.Name, hotel.Address, hotel.Base64Image}
}
