package tgbot

import (
	"log"
	"testing"
)

func TestGetMe(t *testing.T) {
	APIURL, err := LoadBotAPIURL(tokenFname)
	if err != nil {
		t.Fatal("Can't load API URL: " + err.Error())
		return
	}

	user, status, err := GetMe(APIURL)
	if err != nil {
		t.Fatal("getMe failed: " + err.Error())
		return
	}

	if status != 200 {
		t.Fatalf("getMe status is %v (not 200 OK)\n", status)
		return
	}

	if !user.IsBot {
		t.Fatal("user.IsBot from getMe should be true")
		return
	}
}

func TestGetUpdates(t *testing.T) {
	APIURL, err := LoadBotAPIURL(tokenFname)
	if err != nil {
		t.Fatal("Can't load API URL: " + err.Error())
		return
	}

	updates, status, err := GetUpdates(APIURL, Params{"limit": 5, "timeout": 0})
	if err != nil {
		t.Fatal("getUpdates failed: " + err.Error())
		return
	}

	if status != 200 {
		t.Fatalf("GET getUpdates httpStatus is %v (not 200 OK)\n", status)
		return
	}

	log.Printf("Updates: %v", updates)
}
