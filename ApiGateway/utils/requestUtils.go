package utils

import (
	"io"
	"net/http"
	"net/url"

	roundrobin "github.com/hlts2/round-robin"
)

// var BaseHotelServicePath string = "http://localhost:8081"
// var BaseReviewServicePath string = "http://localhost:8082"
// var BaseUserServicePath string = "http://localhost:8083"

var BaseHotelServicePathRoundRobin, _ = roundrobin.New(
    &url.URL{Host: "http://localhost:8081"},
)

var BaseReviewServicePathRoundRobin, _ = roundrobin.New(
    &url.URL{Host: "http://localhost:8082"},
)

var BaseUserServicePathRoundRobin, _ = roundrobin.New(
    &url.URL{Host: "http://localhost:8083"},
)

var BaseBookingServicePathRoundRobin, _ = roundrobin.New(
    &url.URL{Host: "http://localhost:8000"},
)

func DelegateResponse(response *http.Response, w http.ResponseWriter) {
	w.Header().Set("Content-Type", response.Header.Get("Content-Type"))
    w.Header().Set("Content-Length", response.Header.Get("Content-Length"))
    w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(response.StatusCode)
    io.Copy(w, response.Body)
    response.Body.Close()
}
