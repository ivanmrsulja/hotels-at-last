package model

import (
	"time"
)

type ReviewDTO struct {
	Id            uint   `json:"Id"`
	Comment       string `json:"Comment"`
	Rating        int    `json:"Rating"`
	TimesReported int    `json:"TimesReported"`
	UserId        uint   `json:"UserId"`
	UserUsername  string `json:"UserName"`
	CreatedAt     time.Time `json:"CreatedAt"`
}

type ReviewCreateRequestDTO struct {
	Comment string `json:"Comment"`
	Rating  int    `json:"Rating"`
	UserId  uint   `json:"UserId"`
	RoomId  uint   `json:"RoomId"`
}

func (review *ReviewCreateRequestDTO) ToReview() Review {
	return Review{Comment: review.Comment, Rating: review.Rating, TimesReported: 0, UserId: review.UserId, UserUsername: "", RoomId: review.RoomId}
}

type ErrorResponse struct {
	Message    string `json:"Message"`
	StatusCode int    `json:"StatusCode"`
}

type UserDTO struct {
	Id       uint   `json:"Id"`
	Email    string `json:"Email"`
	Username string `json:"Username"`
	Name     string `json:"Name"`
	Surname  string `json:"Surname"`
}

type ReviewsPageable struct {
	Results      []ReviewDTO `json:"Results"`
	TotalResults int32       `json:"TotalResults"`
}
