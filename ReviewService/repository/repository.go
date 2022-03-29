package repository

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	model "github.com/ivanmrsulja/hotels-at-last/review-service/model"
	utils "github.com/ivanmrsulja/hotels-at-last/review-service/utils"
	"github.com/jinzhu/gorm"
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
	// TODO: Treba dodati provjeru da li korisnik postoji i setovati username
	
	trimmedComment := strings.Trim(review.Comment, " ")
	if trimmedComment == "" {
		return review, errors.New("You need to insert a comment.")
	}

	if review.Rating < 1 || review.Rating > 5 {
		return review, errors.New("Rating must be a number between 1 and 5.")
	}

	response, _ := http.Get("http://localhost:8081/api/rooms/" + strconv.FormatUint(uint64(review.RoomId), 10))
	if response.StatusCode != 200 {
		var err model.ErrorResponse
		json.NewDecoder(response.Body).Decode(&err)
		return review, errors.New(err.Message)
	}

	review.UserUsername = "IZMENI IZMENI IZMENI"
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

func GetAllReportedReviews(r *http.Request) []model.Review {
	var reviews []model.Review

	utils.Db.Scopes(Paginate(r)).Where("times_reported > 0").Find(&reviews)

	return reviews
}

func GetAllReviewsForRoom(r *http.Request, id uint) []model.Review {
	var reviews []model.Review

	utils.Db.Scopes(Paginate(r)).Where("room_id = ?", id).Find(&reviews)

	return reviews
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