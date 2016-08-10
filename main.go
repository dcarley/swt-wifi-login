package main

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"net/url"
)

func main() {}

func LoginRequest(username, password string, baseURL string) (*http.Request, error) {
	loginURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	passwordHash := fmt.Sprintf("%x", sha1.Sum([]byte(password)))
	params := url.Values{}
	params.Set("rq", "login")
	params.Set("username", username)
	params.Set("password", passwordHash)
	loginURL.RawQuery = params.Encode()

	return http.NewRequest("GET", loginURL.String(), nil)
}
