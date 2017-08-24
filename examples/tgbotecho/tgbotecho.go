package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/abogovski/Go-TelegramBotAPI/tgbot"
)

func getTokenFname() string {
	fname := "../../tgbot.token"
	if len(os.Args) == 1 {
		log.Print("Token was not provided in os.Args. Using default token filename.")
	} else if len(os.Args) == 2 {
		fname = os.Args[1]
		log.Print("")
	} else if len(os.Args) > 2 {
		log.Print("Too many args provided. Using default token filename.")
	}
	log.Print("Token filename: \"" + fname + "\"")
	return fname
}

func main() {
	// Acquire botAPIURL
	APIURL, err := tgbot.LoadBotAPIURL(getTokenFname())
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	log.Print("TelegramBot API endpoint: " + APIURL)

	// GetMe
	user, err := tgbot.GetMe(APIURL)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	marshaledUser, _ := json.Marshal(user)
	log.Print("BotInfo: " + string(marshaledUser))

	// Start Updates polling
	lastUpdateID := tgbot.Integer(-1)
	for true {
		// poll Messages
		updates, err := tgbot.GetUpdates(APIURL, tgbot.Params{
			"offset":          lastUpdateID + 1,
			"timeout":         15,
			"allowed_updates": []string{"messages"}})
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		log.Printf("received %v updates", len(updates))
		for i := range updates {
			lastUpdateID = updates[i].UpdateID
			if updates[i].Message != nil && updates[i].Message.Text != nil {
				receivedMessage := updates[i].Message
				_, err := tgbot.SendMessage(APIURL, tgbot.Params{
					"chat_id":             receivedMessage.Chat.ID,
					"text":                "Echo " + *receivedMessage.Text,
					"reply_to_message_id": receivedMessage.ID})
				if err != nil {
					log.Fatalf("failed echo message: %v", receivedMessage.ID)
					os.Exit(1)
				}
			}
		}
	}
}
