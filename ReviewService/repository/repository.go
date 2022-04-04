package repository

import (
	"encoding/json"
	"errors"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	roundrobin "github.com/hlts2/round-robin"
	model "github.com/ivanmrsulja/hotels-at-last/review-service/model"
	utils "github.com/ivanmrsulja/hotels-at-last/review-service/utils"
	"github.com/jinzhu/gorm"
)

var baseHotelServicePathRoundRobin, _ = roundrobin.New(
    &url.URL{Host: "http://localhost:8081"},
)

var baseUserServicePathRoundRobin, _ = roundrobin.New(
    &url.URL{Host: "http://localhost:8083"},
)

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
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

func CreateReview(review model.Review) (model.Review, error) {
	trimmedComment := strings.Trim(review.Comment, " ")
	if trimmedComment == "" {
		return review, errors.New("You need to insert a comment.")
	}

	if review.Rating < 1 || review.Rating > 5 {
		return review, errors.New("Rating must be a number between 1 and 5.")
	}

	response, errRoom := http.Get(baseHotelServicePathRoundRobin.Next().Host + "/api/rooms/" + strconv.FormatUint(uint64(review.RoomId), 10))
	if response.StatusCode != 200 {
		var err model.ErrorResponse
		json.NewDecoder(response.Body).Decode(&err)
		return review, errors.New(err.Message)
	}

	if errRoom != nil {
		return review, errors.New("Hotel service not avaliable.")
	}

	responseUser, errUser := http.Get(baseUserServicePathRoundRobin.Next().Host + "/api/users/" + strconv.FormatUint(uint64(review.UserId), 10))
	if responseUser.StatusCode != 200 {
		var err model.ErrorResponse
		json.NewDecoder(responseUser.Body).Decode(&err)
		return review, errors.New(err.Message)
	}

	if errUser != nil {
		return review, errors.New("User service not avaliable.")
	}

	var user model.UserDTO
	json.NewDecoder(responseUser.Body).Decode(&user)
	review.UserUsername = user.Username
	createdReview := utils.Db.Create(&review)

	if createdReview.Error != nil {
		return review, createdReview.Error
	}

	return *createdReview.Value.(*model.Review), nil
}

func ReportReview(id uint) error {
	var review model.Review

	utils.Db.First(&review, id)

	if review.ID == 0 {
		return errors.New("Review with ID " + strconv.FormatUint(uint64(id), 10) + " does not exist.")
	}

	review.TimesReported++
	result := utils.Db.Save(&review)

	return result.Error
}

func DismissReviewReports(id uint) error {
	var review model.Review

	utils.Db.First(&review, id)

	if review.ID == 0 {
		return errors.New("Review with ID " + strconv.FormatUint(uint64(id), 10) + " does not exist.")
	}

	review.TimesReported = 0
	result := utils.Db.Save(&review)

	return result.Error
}

func GetAllReportedReviews(r *http.Request) ([]model.Review, int32) {
	var reviews []model.Review

	pageSize, _ := strconv.Atoi(r.URL.Query().Get("size"))
	switch {
    case pageSize > 100:
      pageSize = 100
    case pageSize <= 0:
      pageSize = 10
    }
	var totalResults float64

	utils.Db.Scopes(Paginate(r)).Where("times_reported > 0").Find(&reviews)
	utils.Db.Table("reviews").Where("times_reported > 0").Select("COUNT(*)").Row().Scan(&totalResults)

	totalPages := int32(math.Ceil(totalResults/float64(pageSize)))

	return reviews, totalPages
}

func GetAllReviewsForRoom(r *http.Request, id uint) ([]model.Review, int32) {
	var reviews []model.Review

	pageSize, _ := strconv.Atoi(r.URL.Query().Get("size"))
	switch {
    case pageSize > 100:
      pageSize = 100
    case pageSize <= 0:
      pageSize = 10
    }
	var totalResults float64

	utils.Db.Scopes(Paginate(r)).Where("room_id = ?", id).Find(&reviews)
	utils.Db.Table("reviews").Where("room_id = ?", id).Select("COUNT(*)").Row().Scan(&totalResults)

	totalPages := int32(math.Ceil(totalResults/float64(pageSize)))

	return reviews, totalPages
}

func DeleteReview(id uint) error {
	var review model.Review

	utils.Db.First(&review, id)

	if review.ID == 0 {
		return errors.New("Review with ID " + strconv.FormatUint(uint64(id), 10) + " does not exist.")
	}

	utils.Db.Delete(&review)

	return nil
}

func GetAverageGradeForRoom(id uint) float64 {
	var n float64
	utils.Db.Table("reviews").Where("room_id = ?", id).Select("AVG(rating)").Row().Scan(&n)

	return n
}
