package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	model "github.com/ivanmrsulja/hotels-at-last/image-service/model"
	utils "github.com/ivanmrsulja/hotels-at-last/image-service/utils"
)

func GetImage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	imagePath, _ := params["path"]

	base64Image := utils.GetB64Image("images/" + imagePath)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.ImageResponse{Base64Image: base64Image})
}

func SaveImage(w http.ResponseWriter, r *http.Request) {
	var image model.ImageResponse
	json.NewDecoder(r.Body).Decode(&image)

	params := mux.Vars(r)
	imagePath, _ := params["path"]

	_ = os.Remove("images/" + imagePath)

	utils.ToJPG(image.Base64Image, "images/" + imagePath)
}
