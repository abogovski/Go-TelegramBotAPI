package tgbot

import (
	"encoding/json"
	"errors"
)

// ResponseParameters https://core.telegram.org/bots/api/#responseparameters
type ResponseParameters struct {
	MigrateToChatID *Integer `json:"migrate_to_chat_id,omitempty"` // Optional. The group has been migrated to a supergroup with the specified identifier. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
	RetryAfter      *Integer `json:"retry_after,omitempty"`        // Optional. In case of exceeding flood control, the number of seconds left to wait before the request can be repeated
}

// Response https://core.telegram.org/bots/api#making-requests
type Response struct {
	Ok bool `json:"ok"`

	// optional
	Description *string             `json:"description,omitempty"`
	Parameters  *ResponseParameters `json:"parameters,omitempty"`
	Result      *json.RawMessage    `json:"result,omitempty"`
}

// GetRawResult safely gets Result from Ok==true response
func (response Response) GetRawResult() (*json.RawMessage, error) {
	if response.Ok == false {
		return nil, errors.New("tgbot.Response.GetRawResult: Ok is false")
	}

	if response.Result == nil {
		return nil, errors.New("tgbot.Response.GetRawResult: Result is nil, although Ok is true")
	}

	if len(*response.Result) == 0 {
		return nil, errors.New("tgbot.Response.GetRawResult: zero len Result")
	}

	return response.Result, nil
}

// GetResultUser safely gets Result from Ok==true response as User
func (response Response) GetResultUser() (*User, error) {
	result, err := response.GetRawResult()
	if err != nil {
		return nil, errors.New("tgbot.Response.GetResultUser: " + err.Error())
	}

	var user User
	err = json.Unmarshal(*result, &user)
	if err != nil {
		return nil, errors.New("tgbot.Response.GetResultUser unmarshal result as User:" + err.Error())
	}

	return &user, nil
}

// GetResultUpdates safely gets Result from Ok==true response as []Update
func (response Response) GetResultUpdates() ([]Update, error) {
	result, err := response.GetRawResult()
	if err != nil {
		return nil, errors.New("tgbot.Response.GetResultUpdates: " + err.Error())
	}

	var updates []Update
	err = json.Unmarshal(*result, &updates)
	if err != nil {
		return nil, errors.New("tgbot.Response.GetResultUpdates unmarshal result as []Update:" + err.Error())
	}

	return updates, nil
}

// GetResultMessage safely gets Result from Ok==true response as Message
func (response Response) GetResultMessage() (*Message, error) {
	result, err := response.GetRawResult()
	if err != nil {
		return nil, errors.New("tgbot.Response.GetResultMessage: " + err.Error())
	}

	var message Message
	err = json.Unmarshal(*result, &message)
	if err != nil {
		return nil, errors.New("tgbot.Response.GetResultMessage unmarshal result as Message:" + err.Error())
	}

	return &message, nil
}
