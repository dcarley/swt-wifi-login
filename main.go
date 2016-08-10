package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type LoginResponse struct {
	Errorcode int `json:"errorcode"`
}

type LoginError struct {
	StatusCode int
	Body       string
}

func (e LoginError) Error() string {
	return fmt.Sprintf("%#v", e)
}

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

func Login(req *http.Request, client *http.Client) error {
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return LoginError{
			StatusCode: resp.StatusCode,
			Body:       fmt.Sprintf("%s", raw),
		}
	}

	var message LoginResponse
	err = json.Unmarshal(raw, &message)
	if err != nil {
		return err
	}

	if message.Errorcode != 0 {
		return LoginError{
			StatusCode: resp.StatusCode,
			Body:       fmt.Sprintf("%s", raw),
		}
	}

	return nil
}
