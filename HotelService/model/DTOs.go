package model

type HotelDTO struct {
	Id          uint   `json:"Id"`
	Name        string `json:"Name"`
	Address     string `json:"Address"`
	Description string `json:"Description"`
	Base64Image string `json:"Base64Image"`
	Stars       int32  `json:"Stars"`
}

func (h *HotelDTO) ToHotel() Hotel {
	return Hotel{Name: h.Name, Address: h.Address, Base64Image: h.Base64Image, Description: h.Description, Stars: h.Stars}
}

type RoomDTO struct {
	Id              uint    `json:"Id"`
	RoomNumber      string  `json:"RoomNumber"`
	NumberOfBeds    int     `json:"NumberOfBeds"`
	Price           float64 `json:"Price"`
	AirConditioned  bool    `json:"AirConditioned"`
	HasParkingSpace bool    `json:"HasParkingSpace"`
	HasTV           bool    `json:"HasTV"`
}

func (r *RoomDTO) ToRoom() Room {
	return Room{RoomNumber: r.RoomNumber, NumberOfBeds: r.NumberOfBeds, Price: r.Price, AirConditioned: r.AirConditioned, HasParkingSpace: r.HasParkingSpace, HasTV: r.HasTV}
}

type ErrorResponse struct {
	Message    string `json:"Message"`
	StatusCode int    `json:"StatusCode"`
}

type HotelsPageable struct {
	Results      []HotelDTO `json:"Results"`
	TotalResults int32      `json:"TotalResults"`
}

type RoomsPageable struct {
	Results      []RoomDTO `json:"Results"`
	TotalResults int32     `json:"TotalResults"`
}

type ImageResponse struct {
	Base64Image string `json:"Base64Image"`
}
