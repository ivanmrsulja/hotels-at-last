package model

import (
	"github.com/dgrijalva/jwt-go"
)

type UserDTO struct {
	Id      uint   `json:"Id"`
	Email   string `json:"Email"`
	Username string `json:"Username"`
	Name    string `json:"Name"`
	Surname string `json:"Surname"`
}

func (user *UserDTO) ToUser() User {
	return User{Email: user.Email, Username: user.Username, Name: user.Name, Surname: user.Surname}
}

type Credentials struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type Claims struct {
	Email string `json:"email"`
	Role UserRole `json:"role"`
	Username string `json:"username"`
	jwt.StandardClaims
}

type ErrorResponse struct {
	Message    string `json:"Message"`
	StatusCode int    `json:"StatusCode"`
}

type LoginResponse struct {
	Token string `json:"Token"`
}

type BanUserRequest struct {
	EndOfBan string `json:"EndOfBan"`
}

type CreateUserRequest struct {
	Email   string `json:"Email"`
	Username string `json:"Username"`
	Name    string `json:"Name"`
	Surname string `json:"Surname"`
	Password string `json:"Password"`
}

func (user *CreateUserRequest) ToUser() User {
	return User{Email: user.Email, Username: user.Username, Name: user.Name, Surname: user.Surname, Password: user.Password}
}
