package utils

import (
	"io"
	"net/http"
)

func DelegateResponse(response *http.Response, w http.ResponseWriter) {
	w.Header().Set("Content-Type", response.Header.Get("Content-Type"))
    w.Header().Set("Content-Length", response.Header.Get("Content-Length"))
	w.WriteHeader(response.StatusCode)
    io.Copy(w, response.Body)
    response.Body.Close()
}
