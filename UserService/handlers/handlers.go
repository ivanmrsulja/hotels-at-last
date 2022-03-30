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
	claims := model.Claims{Email: user.Email, Role: user.Role, StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime.Unix()}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	tokenString, _ := token.SignedString(jwtKey)

	json.NewEncoder(w).Encode(model.LoginResponse{Token: tokenString})
}

func Register(w http.ResponseWriter, r *http.Request) {
	
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
	
}
