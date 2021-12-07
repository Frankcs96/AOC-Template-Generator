package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/Frankcs96/AOC-Template-Generator"
)

func TestGetInputRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		if len(req.Cookies()) != 1 {
			t.Error("you are sending more than one cookie")
		}
		cookie, cookieNotFound := req.Cookie("session")

		if cookieNotFound != nil {
			t.Error("Your cookie should be a session cookie")
		}

		if cookie.Value != "FAKESESSION" {

			t.Error("wrong cookie value")
		}

		rw.Write([]byte(`OK`))

	}))
	defer server.Close()

	c := NewAocClient(server.URL, server.Client(), "FAKESESSION")

	input, _ := c.RequestAocInput()

	if input.Payload != "OK" {

		t.Error("Wrong body")
	}

}

func TestCheckStatusCodes(t *testing.T) {

	err := CheckStatusCode(404)

	if err == nil {

		t.Error("404 should throw an error")

	}
	err = CheckStatusCode(400)

	if err == nil {

		t.Error("400 should throw an error")

	}
	err = CheckStatusCode(200)

	if err != nil {

		t.Error("200 should not throw an error")

	}

}
