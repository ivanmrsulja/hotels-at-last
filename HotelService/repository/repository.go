package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"

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

func GetAllHotels(r *http.Request) ([]model.Hotel, int32) {
	var hotels []model.Hotel

	name := "%" + r.URL.Query().Get("name") + "%"
	address := "%" + r.URL.Query().Get("address") + "%"
	bedsFrom, err1 := strconv.Atoi(r.URL.Query().Get("bedsFrom"))
	bedsTo, err2 := strconv.Atoi(r.URL.Query().Get("bedsTo"))
	priceFrom, err3 := strconv.ParseFloat(r.URL.Query().Get("priceFrom"), 64)
	priceTo, err4 := strconv.ParseFloat(r.URL.Query().Get("priceTo"), 64)
	airCond, err5 := strconv.ParseBool(r.URL.Query().Get("airCond"))
	parking, err6 := strconv.ParseBool(r.URL.Query().Get("parking"))
	tv, err7 := strconv.ParseBool(r.URL.Query().Get("tv"))
	starsFrom, err8 := strconv.Atoi(r.URL.Query().Get("starsFrom"))
	starsTo, err9 := strconv.Atoi(r.URL.Query().Get("starsTo"))
	
	var totalResults int32

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil || err7 != nil || err8 != nil || err9 != nil {
		utils.Db.Scopes(Paginate(r)).Find(&hotels)
		utils.Db.Table("hotels").Where("deleted_at IS NULL").Select("COUNT(*)").Row().Scan(&totalResults)
	} else {
		utils.Db.Scopes(Paginate(r)).Table("hotels").Joins("JOIN rooms ON rooms.hotel_id = hotels.id").Where("hotels.name LIKE ? AND hotels.address LIKE ? AND rooms.number_of_beds BETWEEN ? AND ? AND rooms.price BETWEEN ? AND ? AND rooms.air_conditioned = ? AND rooms.has_parking_space = ? AND rooms.has_tv = ? AND hotels.stars >= ? AND hotels.stars <= ?", name, address, bedsFrom, bedsTo, priceFrom, priceTo, airCond, parking, tv, starsFrom, starsTo).Group("hotels.id").Find(&hotels)
		utils.Db.Table("hotels").Joins("JOIN rooms ON rooms.hotel_id = hotels.id").Where("hotels.deleted_at IS NULL AND hotels.name LIKE ? AND hotels.address LIKE ? AND rooms.number_of_beds BETWEEN ? AND ? AND rooms.price BETWEEN ? AND ? AND rooms.air_conditioned = ? AND rooms.has_parking_space = ? AND rooms.has_tv = ? AND hotels.stars >= ? AND hotels.stars <= ?", name, address, bedsFrom, bedsTo, priceFrom, priceTo, airCond, parking, tv, starsFrom, starsTo).Group("hotels.id").Select("COUNT(*)").Row().Scan(&totalResults)
	}

	return hotels, totalResults
}

func FindRoomsForHotel(r *http.Request, id uint) ([]model.Room, int32) {
	var hotel model.Hotel
	var rooms []model.Room

	var totalResults int32

	utils.Db.First(&hotel, id)
	utils.Db.Model(&hotel).Scopes(Paginate(r)).Related(&rooms)
	utils.Db.Table("rooms").Where("hotel_id = ?", id).Select("COUNT(*)").Row().Scan(&totalResults)


	return rooms, totalResults
}

func FindRoom(id uint) (model.Room, error) {
	var room model.Room

	utils.Db.First(&room, id)

	if room.ID == 0 {
		return room, errors.New("Room with ID " + strconv.FormatUint(uint64(id), 10) + " does not exist")
	}

	return room, nil
}

func FindHotel(id uint) (model.Hotel, error) {
	var hotel model.Hotel

	utils.Db.First(&hotel, id)

	if hotel.ID == 0 {
		return hotel, errors.New("Hotel with ID " + strconv.FormatUint(uint64(id), 10) + " does not exist")
	}

	return hotel, nil
}

func CreateHotel(newHotel model.Hotel) (model.Hotel, error) {

	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	filename := uuid + ".jpeg"
    // utils.ToJPG(newHotel.Base64Image, "images/" + filename)

	body := model.ImageResponse{Base64Image: newHotel.Base64Image}
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(body)
	http.Post("http://localhost:8084/api/images/" + filename, "application/json", bytes.NewReader(reqBodyBytes.Bytes()))

	newHotel.Base64Image = filename

	if newHotel.Stars <= 0 || newHotel.Stars > 5 {
		return newHotel, errors.New("Stars must be an integer between 1 and 5.")
	}

	createdHotel := utils.Db.Create(&newHotel)

	if createdHotel.Error != nil {
		return newHotel, createdHotel.Error
	}

	return newHotel, nil
}

func GetBase64Image(image string) string {
	// return utils.GetB64Image(image)
	response, _ := http.Get("http://localhost:8084/api/images/" + image)
	var imageResponse model.ImageResponse
	json.NewDecoder(response.Body).Decode(&imageResponse)
	return imageResponse.Base64Image
}

func CreateRoom(newRoom model.Room) (model.Room, error) {
	var hotel model.Hotel

	utils.Db.First(&hotel, newRoom.HotelID)

	if hotel.ID == 0 {
		return newRoom, errors.New("Hotel with ID " + strconv.FormatUint(uint64(newRoom.HotelID), 10) + " does not exist.")
	}

	if newRoom.Price <= 0 {
		return newRoom, errors.New("Price must be greater than 0.")
	}

	if newRoom.NumberOfBeds <= 0 {
		return newRoom, errors.New("Number of beds in a room must be greater than 0.")
	}

	createdRoom := utils.Db.Create(&newRoom)

	if createdRoom.Error != nil {
		return newRoom, createdRoom.Error
	}

	return newRoom, nil
}

func UpdateHotel(hotel model.Hotel, id uint) error {
	var hotelToUpdate model.Hotel

	utils.Db.First(&hotelToUpdate, id)
	
	if hotelToUpdate.ID == 0 {
		return errors.New("Hotel with ID " + strconv.FormatUint(uint64(id), 10) + " does not exist.")
	}

	if hotel.Stars <= 0 || hotel.Stars > 5 {
		return errors.New("Stars must be an integer between 1 and 5.")
	}

	if strings.Trim(hotel.Name, " ") == "" || strings.Trim(hotel.Description, " ") == "" || strings.Trim(hotel.Address, " ") == "" {
		return errors.New("There should be no empty fields, you know better :)")
	}

	hotelToUpdate.Name = hotel.Name
	hotelToUpdate.Address = hotel.Address
	hotelToUpdate.Description = hotel.Description
	hotelToUpdate.Stars = hotel.Stars
	if hotel.Base64Image != "" {
		// _ = os.Remove("images/" + hotelToUpdate.Base64Image)
		// utils.ToJPG(hotel.Base64Image, "images/" + hotelToUpdate.Base64Image)
		body := model.ImageResponse{Base64Image: hotel.Base64Image}
		reqBodyBytes := new(bytes.Buffer)
		json.NewEncoder(reqBodyBytes).Encode(body)
		http.Post("http://localhost:8084/api/images/" + hotelToUpdate.Base64Image, "application/json", bytes.NewReader(reqBodyBytes.Bytes()))
	}
	result := utils.Db.Save(&hotelToUpdate)

	return result.Error
}

func UpdateRoom(room model.Room, id uint) error {
	var roomToUpdate model.Room

	utils.Db.First(&roomToUpdate, id)

	if roomToUpdate.ID == 0 {
		return errors.New("Room with ID " + strconv.FormatUint(uint64(id), 10) + " does not exist.")
	}

	if strings.Trim(room.RoomNumber, " ") == "" {
		return errors.New("Room number must not be empty.")
	}

	if room.Price <= 0 {
		return errors.New("Price must be greater than 0.")
	}

	if room.NumberOfBeds <= 0 {
		return errors.New("Number of beds in a room must be greater than 0.")
	}

	roomToUpdate.RoomNumber = room.RoomNumber
	roomToUpdate.Price = room.Price
	roomToUpdate.AirConditioned = room.AirConditioned
	roomToUpdate.HasParkingSpace = room.HasParkingSpace
	roomToUpdate.NumberOfBeds = room.NumberOfBeds
	roomToUpdate.HasTV = room.HasTV
	result := utils.Db.Save(&roomToUpdate)

	return result.Error
}

func DeleteHotel(id uint) error {
	var hotel model.Hotel

	utils.Db.First(&hotel, id)

	if(hotel.ID == 0) {
		return errors.New("Hotel with ID " + strconv.FormatUint(uint64(id), 10) + " does not exist")
	}

	utils.Db.Delete(&hotel)

	return nil
}

func DeleteRoom(id uint) error {
	var room model.Room

	utils.Db.First(&room, id)

	if(room.ID == 0) {
		return errors.New("Room with ID " + strconv.FormatUint(uint64(id), 10) + " does not exist")
	}

	utils.Db.Delete(&room)

	return nil
}