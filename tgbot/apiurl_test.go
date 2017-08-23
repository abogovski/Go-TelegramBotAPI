package tgbot

import (
	"regexp"
	"testing"
)

const tokenFname string = "../tgbot.token"

func TestLoadBotTokenBadFilename(t *testing.T) {
	_, err := LoadBotAPIURL("../#5$%^&.token")
	if err == nil {
		t.Fatal("LoadBotAPIURL(<invalid_fname>) should've failed")
		return
	}
}

func TestLoadBotAPIURLOk(t *testing.T) {
	APIURL, err := LoadBotAPIURL(tokenFname)
	if err != nil {
		t.Fatal("LoadBotAPIURL(tokenFname) failed")
		return
	}

	re := regexp.MustCompile(`https://api.telegram.org/bot[\w]+[:][\w]+/`)
	if !re.MatchString(APIURL) {
		t.Fatal("Loaded bot API URL doesn't match pattern")
	}
}
