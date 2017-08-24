package tgbot

import (
	"testing"
)

func TestGETNoParamsOk(t *testing.T) {
	APIURL, err := LoadBotAPIURL(tokenFname)
	if err != nil {
		t.Fatal("Can't load API URL: " + err.Error())
		return
	}

	response, status, err := Get(APIURL, "getMe", Params{})
	if err != nil {
		t.Fatal("GET request to telegramBotAPI failed: " + err.Error())
		return
	}

	if status != 200 {
		t.Fatalf("GET getMe httpStatus is %v (not 200 OK)\n", status)
		return
	}

	if !response.Ok {
		t.Fatal("GET getMe response.Ok == false")
		return
	}
}

func TestGetWithParamsOk(t *testing.T) {
	APIURL, err := LoadBotAPIURL(tokenFname)
	if err != nil {
		t.Fatal("Can't load API URL: " + err.Error())
		return
	}

	response, status, err := Get(APIURL, "getMe", Params{"limit": 5})
	if err != nil {
		t.Fatal("GET request to telegramBotAPI failed: " + err.Error())
		return
	}

	if status != 200 {
		t.Fatalf("GET getMe httpStatus is %v (not 200 OK)\n", status)
		return
	}

	if !response.Ok {
		t.Fatal("GET getMe response.Ok == false")
		return
	}
}

func TestGetNotFound(t *testing.T) {
	APIURL := GenBotAPIURL("Invalid123ID:Invalid567Token")
	response, status, err := Get(APIURL, "getMe", Params{})
	if err != nil {
		t.Fatal("GET request to telegramBotAPI failed: " + err.Error())
		return
	}

	if status != 404 {
		t.Fatalf("GET getMe httpStatus is %v (not 404 NOT FOUND)\n", status)
		return
	}

	if response.Ok {
		t.Fatal("GET get response.Ok == true")
		return
	}
}

func TestPostUrlEncodedParamsOk(t *testing.T) {
	APIURL, err := LoadBotAPIURL(tokenFname)
	if err != nil {
		t.Fatal("Can't load API URL: " + err.Error())
		return
	}

	response, status, err := PostURLEncoded(APIURL, "getMe", Params{"limit": 5})
	if err != nil {
		t.Fatal("GET request to telegramBotAPI failed: " + err.Error())
		return
	}

	if status != 200 {
		t.Fatalf("GET getMe httpStatus is %v (not 200 OK)\n", status)
		return
	}

	if !response.Ok {
		t.Fatal("GET getMe response.Ok == false")
		return
	}
}

func TestPostJSONParamsOk(t *testing.T) {
	APIURL, err := LoadBotAPIURL(tokenFname)
	if err != nil {
		t.Fatal("Can't load API URL: " + err.Error())
		return
	}

	response, status, err := PostJSON(APIURL, "getMe", Params{"limit": 5})
	if err != nil {
		t.Fatal("GET request to telegramBotAPI failed: " + err.Error())
		return
	}

	if status != 200 {
		t.Fatalf("GET getMe httpStatus is %v (not 200 OK)\n", status)
		return
	}

	if !response.Ok {
		t.Fatal("GET getMe response.Ok == false")
		return
	}
}

// TODO: add TestMultipartFormParamsOk when Params.MultipartFormEncode implemented
