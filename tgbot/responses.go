package tgbot

import "encoding/json"

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
