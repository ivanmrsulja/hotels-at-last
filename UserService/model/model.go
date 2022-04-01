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
	Email string `gorm:"not null;default:null;unique_index"`
	Username string `gorm:"not null;default:null"`
	Password string `gorm:"not null;default:null"`
	Name string `gorm:"not null;default:null"`
	Surname string `gorm:"not null;default:null"`
	Role UserRole
	BannedUntil time.Time
}

func (user *User) ToDTO() UserDTO {
	return UserDTO{Id: user.ID, Email: user.Email, Username: user.Username, Name: user.Name, Surname: user.Surname}
}
