package model

type ReviewDTO struct {
	Id            uint   `json:"Id"`
	Comment       string `json:"Comment"`
	Rating        int    `json:"Rating"`
	TimesReported int    `json:"TimesReported"`
	UserId        uint   `json:"UserId"`
	UserUsername  string `json:"UserName"`
}

func (review *ReviewDTO) toReview() Review {
	return Review{Comment: review.Comment, Rating: review.Rating, TimesReported: review.TimesReported, UserId: review.UserId, UserUsername: review.UserUsername}
}