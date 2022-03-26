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

type Hotel struct {
	gorm.Model

	Name string
	Rooms []Room
}

func (h *Hotel) toDTO() HotelDTO{
	return HotelDTO{h.ID, h.Name}
}