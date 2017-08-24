package tgbot

import "errors"

// GetUpdates https://core.telegram.org/bots/api#getupdates
func GetUpdates(botAPIURL string, params Params) ([]Update, int, error) {
	response, status, err := Get(botAPIURL, "getUpdates", params)
	if err != nil {
		return nil, status, errors.New("tgbot.GetUpdates: " + err.Error())
	}

	updates, err := response.GetResultUpdates()
	if err != nil {
		return nil, status, errors.New("tgbot.GetUpdates" + err.Error())
	}

	return updates, status, nil
}

// GetMe https://core.telegram.org/bots/api#getupdates
func GetMe(botAPIURL string) (*User, int, error) {
	response, status, err := Get(botAPIURL, "getMe", Params{})
	if err != nil {
		return nil, status, errors.New("tgbot.GetMe: " + err.Error())
	}

	user, err := response.GetResultUser()
	if err != nil {
		return nil, status, errors.New("tgbot.GetMe" + err.Error())
	}

	return user, status, nil
}

// SendMessage https://core.telegram.org/bots/api#sendmessage
func SendMessage(botAPIURL string, params Params) (*Message, int, error) {
	response, status, err := Get(botAPIURL, "sendMessage", params)
	if err != nil {
		return nil, status, errors.New("tgbot.sendMessage: " + err.Error())
	}

	message, err := response.GetResultMessage()
	if err != nil {
		return nil, status, errors.New("tgbot.sendMessage" + err.Error())
	}

	return message, status, nil
}

// GetFile https://core.telegram.org/bots/api#getfile
func GetFile(botAPIURL, fileID string) (*File, int, error) {
	response, status, err := Get(botAPIURL, "getFile", Params{"file_id": fileID})
	if err != nil {
		return nil, status, errors.New("tgbot.getFile: " + err.Error())
	}

	file, err := response.GetResultFile()
	if err != nil {
		return nil, status, errors.New("tgbot.getFile" + err.Error())
	}

	return file, status, nil
}
