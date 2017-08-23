package tgbot

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// GenBotAPIURL Generate Telegram API URL from Bot token
func GenBotAPIURL(token string) string {
	return fmt.Sprintf("https://api.telegram.org/bot%v/", token)
}

// LoadBotAPIURL Load Telegram Bot token and generate API URL
func LoadBotAPIURL(fname string) (string, error) {
	f, err := os.Open(fname)
	if err != nil {
		return "", errors.New("tgbotapi.LoadBotAPIURL: " + err.Error())
	}
	defer f.Close()

	token, _ := bufio.NewReader(f).ReadString('\n')
	return GenBotAPIURL(token[:len(token)-1]), nil
}
