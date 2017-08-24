package tgbot

import "errors"

// GetUpdates https://core.telegram.org/bots/api#getupdates
func GetUpdates(botAPIURL string, params Params) ([]Update, error) {
	response, err := Get(botAPIURL, "getUpdates", params)
	if err != nil {
		return nil, errors.New("tgbot.GetUpdates: " + err.Error())
	}

	updates, err := response.GetResultUpdates()
	if err != nil {
		return nil, errors.New("tgbot.GetUpdates" + err.Error())
	}

	return updates, nil
}

// GetMe https://core.telegram.org/bots/api#getupdates
func GetMe(botAPIURL string) (*User, error) {
	response, err := Get(botAPIURL, "getMe", Params{})
	if err != nil {
		return nil, errors.New("tgbot.GetMe: " + err.Error())
	}

	user, err := response.GetResultUser()
	if err != nil {
		return nil, errors.New("tgbot.GetMe" + err.Error())
	}

	return user, nil
}

// SendMessage https://core.telegram.org/bots/api#sendmessage
func SendMessage(botAPIURL string, params Params) (*Message, error) {
	response, err := Get(botAPIURL, "sendMessage", Params{})
	if err != nil {
		return nil, errors.New("tgbot.sendMessage: " + err.Error())
	}

	message, err := response.GetResultMessage()
	if err != nil {
		return nil, errors.New("tgbot.sendMessage" + err.Error())
	}

	return message, nil
}

// GetFile https://core.telegram.org/bots/api#getfile
func GetFile(botAPIURL, fileID string) (*File, error) {
	response, err := Get(botAPIURL, "getFile", Params{"file_id": fileID})
	if err != nil {
		return nil, errors.New("tgbot.getFile: " + err.Error())
	}

	file, err := response.GetResultFile()
	if err != nil {
		return nil, errors.New("tgbot.getFile" + err.Error())
	}

	return file, nil

}
