package repository

import (
	model "github.com/ivanmrsulja/hotels-at-last/model"
	utils "github.com/ivanmrsulja/hotels-at-last/utils"
)

func GetAllHotels() []model.Hotel {
	var hotels []model.Hotel

	utils.Db.Find(&hotels)

	return hotels
}

func FindRoomsForHotel(id uint) []model.Room {
	var hotel model.Hotel
	var rooms []model.Room

	utils.Db.First(&hotel, id)
	utils.Db.Model(&hotel).Related(&rooms)

	return rooms
}