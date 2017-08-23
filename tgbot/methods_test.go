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

	user, err := GetMe(APIURL)
	if err != nil {
		t.Fatal("getMe failed: " + err.Error())
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

	updates, err := GetUpdates(APIURL, Params{"limit": 5, "timeout": 0})
	if err != nil {
		t.Fatal("getUpdates failed: " + err.Error())
		return
	}

	log.Printf("Updates: %v", updates)
}
