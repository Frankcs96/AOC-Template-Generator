package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type AocClient struct {
	baseURL string
	Client  *http.Client
	Session string
}

type InputResponse struct {
	StatusCode int
	Payload    string
}

func NewAocClient(url string, client *http.Client, session string) *AocClient {

	return &AocClient{url, client, session}
}

func (client *AocClient) RequestAocInput() (*InputResponse, error) {

	cookie := &http.Cookie{
		Name:  "session",
		Value: client.Session,
	}

	req, err := http.NewRequest("GET", client.baseURL, nil)
	req.AddCookie(cookie)

	if err != nil {
		fmt.Println(err)
	}

	res, err := client.Client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
	}

	sb := string(body)

	inputResponse := &InputResponse{
		StatusCode: res.StatusCode,
		Payload:    sb,
	}
	return inputResponse, err
}

func GetProblemInput(day string, session string) *InputResponse {

	var client = &http.Client{
		Timeout: time.Second * 10,
	}

	dayToInt, err := strconv.Atoi(day)

	if err != nil {
		fmt.Println("error parsing your day to int")
	}

	AocClient := NewAocClient("https://adventofcode.com/2021/day/"+strconv.Itoa(dayToInt)+"/input", client, session)
	input, err := AocClient.RequestAocInput()

	if err != nil {
		fmt.Println("Error fetching your input")
	}

	return input
}

func CheckStatusCode(statusCode int) error {

	if statusCode == 404 {
		return errors.New("This day is still not active please try another one")
	}

	if statusCode == 400 {
		return errors.New("Invalid session token please try again")
	}
	if statusCode == 500 {
		return errors.New("500: Something happened with the server please try again")
	}
	return nil
}
