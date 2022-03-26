package repository

import (
	"net/http"
	"strconv"

	model "github.com/ivanmrsulja/hotels-at-last/model"
	utils "github.com/ivanmrsulja/hotels-at-last/utils"
	"github.com/jinzhu/gorm"
)

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
  return func (db *gorm.DB) *gorm.DB {
    page, _ := strconv.Atoi(r.URL.Query().Get("page"))
    if page < 0 {
      page = 0
    }

    pageSize, _ := strconv.Atoi(r.URL.Query().Get("size"))
    switch {
    case pageSize > 100:
      pageSize = 100
    case pageSize <= 0:
      pageSize = 10
    }

    offset := page * pageSize
    return db.Offset(offset).Limit(pageSize)
  }
}

func GetAllHotels(r *http.Request) []model.Hotel {
	var hotels []model.Hotel

	utils.Db.Scopes(Paginate(r)).Find(&hotels)

	return hotels
}

func FindRoomsForHotel(r *http.Request, id uint) []model.Room {
	var hotel model.Hotel
	var rooms []model.Room

	utils.Db.First(&hotel, id)
	utils.Db.Model(&hotel).Scopes(Paginate(r)).Related(&rooms)

	return rooms
}

func FindRoom(id uint) model.Room {
	var room model.Room

	utils.Db.First(&room, id)

	return room
}