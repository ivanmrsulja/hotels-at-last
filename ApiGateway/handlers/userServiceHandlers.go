package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ivanmrsulja/hotels-at-last/api-gateway/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8083/api/users/login", r.Body)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	response, _ := client.Do(req)

	utils.DelegateResponse(response, w)
}

func Register(w http.ResponseWriter, r *http.Request) {
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8083/api/users/register", r.Body)
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	response, _ := client.Do(req)

	utils.DelegateResponse(response, w)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, _ := strconv.ParseUint(params["id"], 10, 32)

	response, _ := http.Get("http://localhost:8083/api/users/" + strconv.FormatUint(uint64(userId), 10))
	
	utils.DelegateResponse(response, w)
}

func BanUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, _ := strconv.ParseUint(params["id"], 10, 32)

	response, _ := http.Get("http://localhost:8083/api/users/" + strconv.FormatUint(uint64(userId), 10) + "/ban")
	
	utils.DelegateResponse(response, w)
}
