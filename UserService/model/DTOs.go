package model

import "github.com/dgrijalva/jwt-go"

type UserDTO struct {
	Id      uint   `json:"Id"`
	Email   string `json:"Email"`
	Name    string `json:"Name"`
	Surname string `json:"Surname"`
}

func (user *UserDTO) ToUser() User {
	return User{Email: user.Email, Name: user.Name, Surname: user.Surname}
}

type Credentials struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type Claims struct {
	Email string `json:"email"`
	Role UserRole `json:"role"`
	jwt.StandardClaims
}

type ErrorResponse struct {
	Message    string `json:"Message"`
	StatusCode int    `json:"StatusCode"`
}

type LoginResponse struct {
	Token string `json:"Token"`
}
