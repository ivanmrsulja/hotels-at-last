package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	model "github.com/ivanmrsulja/hotels-at-last/user-service/model"
	repository "github.com/ivanmrsulja/hotels-at-last/user-service/repository"
)

// This should be stored as an envinroinment variable
var jwtKey = []byte("z7031Q8Qy9zVO-T2o7lsFIZSrd05hH0PaeaWIBvLh9s")

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials model.Credentials
	json.NewDecoder(r.Body).Decode(&credentials)

	user, err := repository.CheckCredentials(credentials.Email, credentials.Password)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusUnauthorized})
		return
	}

	expirationTime := time.Now().Add(time.Hour * 24)
	claims := model.Claims{Email: user.Email, Role: user.Role, Username: user.Username, StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime.Unix()}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	tokenString, _ := token.SignedString(jwtKey)

	json.NewEncoder(w).Encode(model.LoginResponse{Token: tokenString})
}

func Register(w http.ResponseWriter, r *http.Request) {
	var userDTO model.CreateUserRequest
	json.NewDecoder(r.Body).Decode(&userDTO)

	createdUser, err := repository.CreateUser(userDTO.ToUser())

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusBadRequest})
	} else {
		json.NewEncoder(w).Encode(createdUser.ToDTO())
	}
}

func AuthoriseAdmin(w http.ResponseWriter, r *http.Request) {
	cookie := r.Header.Values("Authorization")
	tokenString := strings.Split(cookie[0], " ")[1]
	
	claims := model.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, 
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
	
	if err != nil || !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if token.Claims.(*model.Claims).Role != model.ADMIN {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func AuthoriseUser(w http.ResponseWriter, r *http.Request) {
	cookie := r.Header.Values("Authorization")
	tokenString := strings.Split(cookie[0], " ")[1]
	
	claims := model.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, 
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
	
	if err != nil || !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if token.Claims.(*model.Claims).Role != model.USER {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, _ := strconv.ParseUint(params["id"], 10, 32)

	user, err := repository.FindUserById(uint(userId))

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	} else {
		json.NewEncoder(w).Encode(user.ToDTO())
	}
}

func BanUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, _ := strconv.ParseUint(params["id"], 10, 32)
	var banRequest model.BanUserRequest
	json.NewDecoder(r.Body).Decode(&banRequest)

	endOfBan, parseErr := time.Parse("2006-01-02", banRequest.EndOfBan)

	if parseErr != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{Message: parseErr.Error(), StatusCode: http.StatusBadRequest})
		return
	}

	err := repository.BanUser(uint(userId), endOfBan)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusNotFound})
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
