package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
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

func main() {
	const (
		baseURL    = "http://swt.passengerwifi.com/cws/"
		reqTimeout = 5 * time.Second
		reqRetry   = 5 * time.Second
	)

	if len(os.Args) != 3 {
		ExitWithUsage()
	}

	username, password := os.Args[1], os.Args[2]
	if username == "" || password == "" {
		ExitWithUsage()
	}

	req, err := LoginRequest(username, password, baseURL)
	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{
		Transport: &http.Transport{
			ResponseHeaderTimeout: reqTimeout,
		},
	}

	for {
		err := Login(req, client)
		if err == nil {
			log.Println("login successful!")
			break
		}

		log.Println("login failed:", err)
		log.Println("sleeping for:", reqRetry)
		time.Sleep(reqRetry)
	}
}

func ExitWithUsage() {
	fmt.Fprintf(os.Stderr, "Usage: %s <username> <password>\n", os.Args[0])
	os.Exit(2)
}

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
