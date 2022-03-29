package model

type ReviewDTO struct {
	Id            uint   `json:"Id"`
	Comment       string `json:"Comment"`
	Rating        int    `json:"Rating"`
	TimesReported int    `json:"TimesReported"`
	UserId        uint   `json:"UserId"`
	UserUsername  string `json:"UserName"`
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
