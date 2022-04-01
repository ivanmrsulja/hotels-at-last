package utils

import (
	"io"
	"net/http"
)

var BaseHotelServicePath string = "http://localhost:8081"
var BaseReviewServicePath string = "http://localhost:8082"
var BaseUserServicePath string = "http://localhost:8083"

func DelegateResponse(response *http.Response, w http.ResponseWriter) {
	w.Header().Set("Content-Type", response.Header.Get("Content-Type"))
    w.Header().Set("Content-Length", response.Header.Get("Content-Length"))
	w.WriteHeader(response.StatusCode)
    io.Copy(w, response.Body)
    response.Body.Close()
}
