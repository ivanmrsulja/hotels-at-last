package model

type HotelDTO struct {
	Id      uint   `json:"Id"`
	Name    string `json:"Name"`
	Address string `json:"Address"`
}

func (h *HotelDTO) ToHotel() Hotel {
	return Hotel{Name: h.Name, Address: h.Address}
}

type RoomDTO struct {
	Id              uint    `json:"Id"`
	NumberOfBeds    int     `json:"NumberOfBeds"`
	Price           float64 `json:"Price"`
	AirConditioned  bool    `json:"AirConditioned"`
	HasParkingSpace bool    `json:"HasParkingSpace"`
	HasTV           bool    `json:"HasTV"`
}

func (r *RoomDTO) ToRoom() Room {
	return Room{NumberOfBeds: r.NumberOfBeds, Price: r.Price, AirConditioned: r.AirConditioned, HasParkingSpace: r.HasParkingSpace, HasTV: r.HasTV}
}