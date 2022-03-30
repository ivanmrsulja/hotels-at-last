package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type UserRole int

const (
	USER  UserRole = 0
	ADMIN UserRole = 1
)

type User struct {
	gorm.Model
	Email string
	Password string
	Name string
	Surname string
	Role UserRole
	BannedUntil time.Time
}

func (user *User) ToDTO() UserDTO {
	return UserDTO{Id: user.ID, Email: user.Email, Name: user.Name, Surname: user.Surname}
}
