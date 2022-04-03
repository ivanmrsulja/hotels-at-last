package model

type CreateReservationRequest struct {
	UserId   int    `json:"user_id"`
	RoomId   int    `json:"room_id"`
	Start    string `json:"start"`
	End      string `json:"end"`
	Username string `json:"username"`
}
