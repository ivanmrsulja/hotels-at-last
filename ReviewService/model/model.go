package model

import (
	"github.com/jinzhu/gorm"
)

type Review struct {
	gorm.Model

	Comment string
	Rating int
	TimesReported int
	UserId uint
	RoomId uint
	UserUsername string
}

func (review *Review) ToDTO() ReviewDTO {
	return ReviewDTO{Id: review.ID, Comment: review.Comment, Rating: review.Rating, TimesReported: review.TimesReported, UserId: review.UserId, UserUsername: review.UserUsername}
}
