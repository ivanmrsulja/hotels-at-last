package utils

import (
	"bytes"
	"errors"
	"net/http"
)

// roles: user, admin (all lowercase letters)
func AuthorizeRole(r *http.Request, role string) error {
	authRequest, _ := http.NewRequest(http.MethodGet, "http://localhost:8083/api/users/authorize/" + role, bytes.NewBufferString(""))
	authRequest.Header.Set("Accept", "application/json")
	values := r.Header.Values("Authorization")

	if len(values) == 0 {
		return errors.New("Unauthorised")
	}

	authRequest.Header.Set("Authorization", values[0])
	authClient := &http.Client{}
	authResponse, _ := authClient.Do(authRequest)

	if authResponse.StatusCode != 200 {
		return errors.New("Unauthorised")
	}

	return nil
}

