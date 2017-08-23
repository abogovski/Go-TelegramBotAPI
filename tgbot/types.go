package tgbot

import "encoding/json"

///////////////////////////////////////////////////////////////////////////////
// Custom Primitive Types
///////////////////////////////////////////////////////////////////////////////

// Dummy type for json-structs that are not implemented yet
type Dummy *json.RawMessage

// Placeholder Some fields are placeholders
type Placeholder Dummy

// Integer represents Telegram's Integer type
type Integer int64

///////////////////////////////////////////////////////////////////////////////
// General Types
///////////////////////////////////////////////////////////////////////////////

// User https://core.telegram.org/bots/api#user
type User struct {
	ID        Integer `json:"id"`         // Unique identifier for this user or bot
	IsBot     bool    `json:"is_bot"`     // True, if this user is a bot
	FirstName string  `json:"first_name"` // User's or bot's first name

	// Optional
	LastName     *string `json:"last_name,omitempty"`     // Optional. User's or bot's last name
	Username     *string `json:"username,omitempty"`      // Optional. User's or bot's username
	LanguageCode *string `json:"language_code,omitempty"` // Optional. IETF language tag of user's language
}

// Chat https://core.telegram.org/bots/api#chat
type Chat struct {
	ID   Integer `json:"id"`   // Unique identifier for this chat. (< 52 bits)
	Type string  `json:"type"` // Type of chat: "private", "group", "supergroup", "channel"

	// Optional, if available
	Title                       *string `json:"title,omitempty"`                          // Optional. Title, for supergroups, channels and group chats
	Username                    *string `json:"username,omitempty"`                       // Optional. Username, for private chats, supergroups and channels if available
	FirstName                   *string `json:"first_name,omitempty"`                     // Optional. First name of the other party in a private chat
	LastName                    *string `json:"last_name,omitempty"`                      // Optional. Last name of the other party in a private chat
	AllMembersAreAdministrators *bool   `json:"all_members_are_administrators,omitempty"` // Optional. True if a group has ‘All Members Are Admins’ enabled

	// Optional, returned only in getChat
	Photo       *ChatPhoto `json:"photo,omitempty"`       // Optional. Chat photo. Returned only in getChat
	Description *string    `json:"description,omitempty"` // Optional. Description, for supergroups and channel chats. Returned only in getChat
	InviteLink  *string    `json:"invite_link,omitempty"` // Optional. Chat invite link, for supergroups and channel chats. Returned only in getChat
}

// Message https://core.telegram.org/bots/api#message
type Message struct {
	ID Integer `json:"id"` // Unique message identifier inside this chat

	// Optional
	From *User `json:"from,omitempty"` // Optional. Sender, can be empty for messages sent to channels

	Date Integer `json:"date"` // Date the message was sent in Unix time
	Chat Chat    `json:"chat"` // Conversation the message belongs to

	// Optional, for forwarded messages
	ForwardFrom          *User    `json:"forward_from,omitempty"`            // Optional. For forwarded messages, sender of the original message
	ForwardFromChat      *Chat    `json:"forward_from_chat,omitempty"`       // Optional. For messages forwarded from a channel, information about the original channel
	ForwardFromMessageID *Integer `json:"forward_from_message_id,omitempty"` // Optional. For forwarded channel posts, identifier of the original message in the channel
	ForwardDate          *Integer `json:"forward_date,omitempty"`            // Optional. For forwarded messages, date the original message was sent in Unix time

	// Optional, for replies
	ReplyToMessage *Message `json:"reply_to_message,omitempty"` // Optional. For replies, the original message. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.

	// Optional, if edited
	EditDate *Integer `json:"edit_date,omitempty"` // Optional. Date the message was last edited in Unix time

	// Optional, for text messages
	Text     *string         `json:"text,omitempty"`     // Optional. For text messages, the actual UTF-8 text of the message, 0-4096 characters.
	Entities []MessageEntity `json:"entities,omitempty"` // Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text

	// Optional, other types of messages
	Audio                 *Audio             `json:"audio,omitempty"`                   // Optional. Message is an audio file, information about the file
	Document              *Document          `json:"document,omitempty"`                // Optional. Message is a general file, information about the file
	Game                  *Game              `json:"game,omitempty"`                    // Optional. Message is a game, information about the game
	Photo                 []PhotoSize        `json:"photo,omitempty"`                   // Optional. Message is a photo, available sizes of the photo
	Sticker               *Sticker           `json:"sticker,omitempty"`                 // Optional. Message is a sticker, information about the sticker
	Video                 *Video             `json:"video,omitempty"`                   // Optional. Message is a video, information about the video
	Voice                 *Voice             `json:"voice,omitempty"`                   // Optional. Message is a voice message, information about the file
	VideoNote             *VideoNote         `json:"video_note,omitempty"`              // Optional. Message is a video note, information about the video message
	NewChatMembers        []User             `json:"new_chat_members,omitempty"`        // Optional. New members that were added to the group or supergroup and information about them (the bot itself may be one of these members)
	Caption               *string            `json:"caption,omitempty"`                 // Optional. Caption for the document, photo or video, 0-200 characters
	Contact               *Contact           `json:"contact,omitempty"`                 // Optional. Message is a shared contact, information about the contact
	Location              *Location          `json:"location,omitempty"`                // Optional. Message is a shared location, information about the location
	Venue                 *Venue             `json:"venue,omitempty"`                   // Optional. Message is a venue, information about the venue
	NewChatMember         *User              `json:"new_chat_member,omitempty"`         // Optional. A new member was added to the group, information about them (this member may be the bot itself)
	LeftChatMember        *User              `json:"left_chat_member,omitempty"`        // Optional. A member was removed from the group, information about them (this member may be the bot itself)
	NewChatTitle          *string            `json:"new_chat_title,omitempty"`          // Optional. A chat title was changed to this value
	NewChatPhoto          []PhotoSize        `json:"new_chat_photo,omitempty"`          // Optional. A chat photo was change to this value
	DeleteChatPhoto       bool               `json:"delete_chat_photo,omitempty"`       // *True* Optional. Service message: the chat photo was deleted
	GroupChatCreated      bool               `json:"group_chat_created,omitempty"`      // *True* Optional. Service message: the group has been created
	SupergroupChatCreated bool               `json:"supergroup_chat_created,omitempty"` // *True* Optional. Service message: the supergroup has been created. This field can‘t be received in a message coming through updates, because bot can’t be a member of a supergroup when it is created. It can only be found in reply_to_message if someone replies to a very first message in a directly created supergroup
	ChannelChatCreated    bool               `json:"channel_chat_created,omitempty"`    // *True* Optional. Optional. Service message: the channel has been created. This field can‘t be received in a message coming through updates, because bot can’t be a member of a channel when it is created. It can only be found in reply_to_message if someone replies to a very first message in a channel
	MigrateToChatID       *Integer           `json:"migrate_to_chat_id,omitempty"`      // Optional. The group has been migrated to a supergroup with the specified identifier. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier
	MigrateFromChatID     *Integer           `json:"migrate_from_chat_id,omitempty"`    // Optional. The supergroup has been migrated from a group with the specified identifier. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier
	PinnedMessage         *Message           `json:"pinned_message,omitempty"`          // Optional. Specified message was pinned. Note that the Message object in this field will not contain further reply_to_message fields even if it is itself a reply
	Invoice               *Invoice           `json:"invoice,omitempty"`                 // Optional. Message is an invoice for a payment, information about the invoice
	SuccessfulePayment    *SuccessfulPayment `json:"successful_payment"`                // Optional. Message is a service message about a successful payment, information about the payment
}

// MessageEntity https://core.telegram.org/bots/api#messageentity
type MessageEntity struct {
	Type   string  `json:"type"`   // Type of the entity. Can be mention (@username), hashtag, bot_command, url, email, bold (bold text), italic (italic text), code (monowidth string), pre (monowidth block), text_link (for clickable text URLs), text_mention (for users without usernames)
	Offset Integer `json:"offset"` // Offset in UTF-16 code units to the start of the entity
	Length Integer `json:"length"` // Length of the entity in UTF-16 code units

	// Optional
	URL  *string // Optional. For “text_link” only, url that will be opened after user taps on the text
	User *User   // Optional. For “text_mention” only, the mentioned user
}

// PhotoSize https://core.telegram.org/bots/api/#photosize
type PhotoSize struct {
	FileID string  `json:"file_id"` // Unique identifier for this file
	Width  Integer `json:"width"`   // Photo width
	Height Integer `json:"height"`  // Photo height

	// Optional
	FileSize *Integer `json:"file_size,omitempty"` // Optional. File size
}

// Audio https://core.telegram.org/bots/api/#audio
type Audio struct {
	FileID   string  `json:"file_id"`  // Unique identifier for this file
	Duration Integer `json:"duration"` // Duration of the audio in seconds as defined by sender

	// Optional
	Performer *string  `json:"performer,omitempty"` // Optional. Performer of the audio as defined by sender or by audio tags
	Title     *string  `json:"title,omitempty"`     // Optional. Title of the audio as defined by sender or by audio tags
	MIMEType  *string  `json:"mime_type,omitempty"` // Optional. MIME type of the file as defined by sender
	FileSize  *Integer `json:"file_size"`           // Optional. File size
}

// Document https://core.telegram.org/bots/api/#document
type Document struct {
	FileID string `json:"file_id"` // Unique file identifier

	// Optional
	Thumb    *PhotoSize `json:"thumb,omitempty"`     // Optional. Document thumbnail as defined by sender
	FileName *string    `json:"file_name,omitempty"` // Optional. Original filename as defined by sender
	MIMEType *string    `json:"mime_type,omitempty"` // Optional. MIME type of the file as defined by sender
	FileSize *Integer   `json:"file_size,omitempty"` // Optional. File size
}

// Video https://core.telegram.org/bots/api/#video
type Video struct {
	FileID   string  `json:"file_id"`  // Unique identifier for this file
	Width    Integer `json:"width"`    // Video width as defined by sender
	Height   Integer `json:"height"`   // Video height as defined by sender
	Duration Integer `json:"duration"` // Duration of the video in seconds as defined by sender

	// Optional
	Thumb    *PhotoSize `json:"thumb,omitempty"`     // Optional. Video thumbnail
	MIMEType *string    `json:"mime_type,omitempty"` // Optional. Mime type of a file as defined by sender
	FileSize *Integer   `json:"file_size,omitempty"` // Optional. File size
}

// Voice https://core.telegram.org/bots/api/#voice
type Voice struct {
	FileID   string  `json:"file_id"`  // Unique identifier for this file
	Duration Integer `json:"duration"` // Duration of the audio in seconds as defined by sender

	// Optional
	MIMEType *string  `json:"mime_type,omitempty"` // Optional. MIME type of the file as defined by sender
	FileSize *Integer `json:"file_size,omitempty"` // Optional. File size
}

// VideoNote https://core.telegram.org/bots/api/#videonote
type VideoNote struct {
	FileID   string  `json:"file_id"`  // Unique identifier for this file
	Length   Integer `json:"length"`   // Video width and height as defined by sender
	Duration Integer `json:"duration"` //	Duration of the video in seconds as defined by sender

	// Optional
	Thumb    *PhotoSize `json:"thumb,omitempty"`     // Optional. Video thumbnail
	FileSize *Integer   `json:"file_size,omitempty"` // Optional. File size
}

// Contact https://core.telegram.org/bots/api/#contact
type Contact struct {
	PhoneNumber string `json:"phone_number"` // Contact's phone number
	FirstName   string `json:"first_name"`   // Contact's first name

	// Optional
	LastName *string  `json:"last_name,omitempty"` // Optional. Contact's last name
	UserID   *Integer `json:"user_id,omitempty"`   // Optional. Contact's user identifier in Telegram
}

// Location https://core.telegram.org/bots/api/#location
type Location struct {
	Longitude float64 `json:"longitude"` // Longitude as defined by sender
	Latitude  float64 `json:"latitude"`  // Latitude as defined by sender
}

// Venue https://core.telegram.org/bots/api/#venue
type Venue struct {
	Location Location `json:"location"` // Venue location
	Title    string   `json:"title"`    // Name of the venue
	Address  string   `json:"address"`  // Address of the venue

	// Optional
	FoursquareID *string `json:"foursquare_id,omitempty"` // Optional. Foursquare identifier of the venue
}

// UserProfilePhotos https://core.telegram.org/bots/api/#userprofilephotos
type UserProfilePhotos struct {
	TotalCount Integer       `json:"total_count"` // Total number of profile pictures the target user has
	Photos     [][]PhotoSize `json:"photos"`      // Requested profile pictures (in up to 4 sizes each)
}

// File https://core.telegram.org/bots/api/#file
type File struct {
	FileID string `json:"file_id"` // Unique identifier for this file

	// Optional
	FileSize *Integer `json:"file_size"` // Optional. File size, if known
	FilePath *string  `json:"file_path"` // Optional. File path. Use https://api.telegram.org/file/bot<token>/<file_path> to get the file.
}

// ReplyKeyboardMarkup https://core.telegram.org/bots/api/#replykeyboardmarkup
type ReplyKeyboardMarkup struct {
	Keyboard [][]KeyboardButton `json:"keyboard"` // Array of button rows, each represented by an Array of KeyboardButton objects

	// Optional
	ResizeKeyboard  bool `json:"resize_keyboard,omitempty"`   // Optional. Requests clients to resize the keyboard vertically for optimal fit (e.g., make the keyboard smaller if there are just two rows of buttons). Defaults to false, in which case the custom keyboard is always of the same height as the app's standard keyboard.
	OneTimeKeyboard bool `json:"one_time_keyboard,omitempty"` // Optional. Requests clients to hide the keyboard as soon as it's been used. The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat – the user can press a special button in the input field to see the custom keyboard again. Defaults to false.
	Selective       bool `json:"selective,omitempty"`         // Optional. Use this parameter if you want to show the keyboard to specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
	// Selective Example: A user requests to change the bot‘s language, bot replies to the request with a keyboard to select the new language. Other users in the group don’t see the keyboard.
}

// KeyboardButton https://core.telegram.org/bots/api/#keyboardbutton
type KeyboardButton struct {
	Text string `json:"text"` // Text of the button. If none of the optional fields are used, it will be sent to the bot as a message when the button is pressed

	// Optional
	RequestContact  *bool `json:"request_contact,omitempty"`  // Optional. If True, the user's phone number will be sent as a contact when the button is pressed. Available in private chats only
	RequestLocation *bool `json:"request_location,omitempty"` // Optional. If True, the user's current location will be sent when the button is pressed. Available in private chats only
}

// ReplyKeyboardRemove https://core.telegram.org/bots/api/#replykeyboardremove
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"` // *True* Requests clients to remove the custom keyboard (user will not be able to summon this keyboard; if you want to hide the keyboard from sight but keep it accessible, use one_time_keyboard in ReplyKeyboardMarkup)

	// Optional
	Selective *bool `json:"selective,omitempty"` // Optional. Use this parameter if you want to remove the keyboard for specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
	//  Example: A user votes in a poll, bot returns confirmation message in reply to the vote and removes the keyboard for that user, while still showing the keyboard with poll options to users who haven't voted yet.
}

// InlineKeyboardMarkup https://core.telegram.org/bots/api/#inlinekeyboardmarkup
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"` // Array of button rows, each represented by an Array of InlineKeyboardButton objects
}

// InlineKeyboardButton https://core.telegram.org/bots/api/#inlinekeyboardbutton
type InlineKeyboardButton struct {
	Text string `json:"text"` // Label text on the button

	// Optional
	URL               *string `json:"url,omitempty"`                 // Optional. HTTP url to be opened when button is pressed
	CallbackData      *string `json:"callback_data,omitempty"`       // Optional. Data to be sent in a callback query to the bot when button is pressed, 1-64 bytes
	SwitchInlineQuery *string `json:"switch_inline_query,omitempty"` // Optional. If set, pressing the button will prompt the user to select one of their chats, open that chat and insert the bot‘s username and the specified inline query in the input field. Can be empty, in which case just the bot’s username will be inserted.
	// Note: This offers an easy way for users to start using your bot in inline mode when they are currently in a private chat with it. Especially useful when combined with switch_pm… actions – in this case the user will be automatically returned to the chat they switched from, skipping the chat selection screen.

	// Optional
	SwitchInlineQueryCurrentChat *string `json:"switch_inline_query_current_chat,omitempty"` // Optional. If set, pressing the button will insert the bot‘s username and the specified inline query in the current chat's input field. Can be empty, in which case only the bot’s username will be inserted.
	// This offers a quick way for the user to open your bot in inline mode in the same chat – good for selecting something from multiple options.

	// Optional
	CallbackGame *CallbackGame `json:"callback_game,omitempty"` // Optional. Description of the game that will be launched when the user presses the button.
	// NOTE: This type of button must always be the first button in the first row.

	// Optional
	Pay bool `json:"pay,omitempty"` // Optional. Specify True, to send a Pay button.
	// NOTE: This type of button must always be the first button in the first row.
}

// CallbackQuery https://core.telegram.org/bots/api/#callbackquery
type CallbackQuery struct {
	ID   string `json:"id"`   // Unique identifier for this query
	From User   `json:"from"` // Sender

	// Optional
	Message         *Message `json:"message,omitempty"`           // Optional. Message with the callback button that originated the query. Note that message content and message date will not be available if the message is too old
	InlineMessageID *string  `json:"inline_message_id,omitempty"` // Optional. Identifier of the message sent via the bot in inline mode, that originated the query.

	ChatInstance string `json:"chat_instance"` // Global identifier, uniquely corresponding to the chat to which the message with the callback button was sent. Useful for high scores in games.

	// Optional
	Data          *string `json:"data,omitempty"`            // Optional. Data associated with the callback button. Be aware that a bad client can send arbitrary data in this field.
	GameShortName *string `json:"game_short_name,omitempty"` // Optional. Short name of a Game to be returned, serves as the unique identifier for the game
}

// ForceReply https://core.telegram.org/bots/api/#forcereply
type ForceReply struct {
	ForceReply bool `json:"force_reply"`         // *True* Shows reply interface to the user, as if they manually selected the bot‘s message and tapped ’Reply'
	Selective  bool `json:"selective,omitempty"` // Optional. Use this parameter if you want to force reply from specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
}

// ChatPhoto https://core.telegram.org/bots/api/#chatphoto
type ChatPhoto struct {
	SmallFileID string `json:"small_file_id"` // Unique file identifier of small (160x160) chat photo. This file_id can be used only for photo download.
	BigFileID   string `json:"big_file_id"`   // Unique file identifier of big (640x640) chat photo. This file_id can be used only for photo download.
}

// ChatMember https://core.telegram.org/bots/api/#chatmember
type ChatMember struct {
	User   User   `json:"user"`   // Information about the user
	Status string `json:"status"` // The member's status in the chat. Can be “creator”, “administrator”, “member”, “restricted”, “left” or “kicked”

	// Optional
	UntilDate             Integer `json:"until_date,omitempty"`                // Optional. Restictred and kicked only. Date when restrictions will be lifted for this user, unix time
	CanBeEdited           *bool   `json:"can_be_edited,omitempty"`             // Optional. Administrators only. True, if the bot is allowed to edit administrator privileges of that user
	CanChangeInfo         *bool   `json:"can_change_info,omitempty"`           // Optional. Administrators only. True, if the administrator can change the chat title, photo and other settings
	CanPostMessages       *bool   `json:"can_post_messages,omitempty"`         // Optional. Administrators only. True, if the administrator can post in the channel, channels only
	CanEditMessages       *bool   `json:"can_edit_messages,omitempty"`         // Optional. Administrators only. True, if the administrator can edit messages of other users, channels only
	CanDeleteMessages     *bool   `json:"can_delete_messages,omitempty"`       // Optional. Administrators only. True, if the administrator can delete messages of other users
	CanInviteUsers        *bool   `json:"can_invite_users,omitempty"`          // Optional. Administrators only. True, if the administrator can invite new users to the chat
	CanRestrictMembers    *bool   `json:"can_restrict_members,omitempty"`      // Optional. Administrators only. True, if the administrator can restrict, ban or unban chat members
	CanPinMessages        *bool   `json:"can_pin_messages,omitempty"`          //	Optional. Administrators only. True, if the administrator can pin messages, supergroups only
	CanPromoteMembers     *bool   `json:"can_promote_members,omitempty"`       // Optional. Administrators only. True, if the administrator can add new administrators with a subset of his own privileges or demote administrators that he has promoted, directly or indirectly (promoted by administrators that were appointed by the user)
	CanSendMessages       *bool   `json:"can_send_messages,omitempty"`         // Optional. Restricted only. True, if the user can send text messages, contacts, locations and venues
	CanSendMediaMessages  *bool   `json:"can_send_media_messages,omitempty"`   // Optional. Restricted only. True, if the user can send audios, documents, photos, videos, video notes and voice notes, implies can_send_messages
	CanSendOtherMessages  *bool   `json:"can_send_other_messages,omitempty"`   // Optional. Restricted only. True, if the user can send animations, games, stickers and use inline bots, implies can_send_media_messages
	CanAddWebPagePreviews *bool   `json:"can_add_web_page_previews,omitempty"` // Optional. Restricted only. True, if user may add web page previews to his messages, implies can_send_media_messages
}

///////////////////////////////////////////////////////////////////////////////
// Update Types
///////////////////////////////////////////////////////////////////////////////

// Update https://core.telegram.org/bots/api#update
type Update struct {
	UpdateID Integer `json:"update_id"` // The update‘s unique identifier. Update identifiers start from a certain positive number and increase sequentially. This ID becomes especially handy if you’re using Webhooks, since it allows you to ignore repeated updates or to restore the correct update sequence, should they get out of order.

	// Optional. At most one of the optional parameters can be present in any given update.
	Message            *Message            `json:"message,omitempty"`              // Optional. New incoming message of any kind — text, photo, sticker, etc.
	EditedMessage      *Message            `json:"edited_message,omitempty"`       // Optional. New version of a message that is known to the bot and was edited
	ChannelPost        *Message            `json:"channel_post,omitempty"`         // Optional. New incoming channel post of any kind — text, photo, sticker, etc.
	EditedChannelPost  *Message            `json:"edited_channel_post,omitempty"`  // Optional. New version of a channel post that is known to the bot and was edited
	InlineQuery        *InlineQuery        `json:"inline_query,omitempty"`         // Optional. New incoming inline query
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"` // Optional. The result of an inline query that was chosen by a user and sent to their chat partner.
	CallbackQuery      *CallbackQuery      `json:"callback_query,omitempty"`       // Optional. New incoming callback query
	ShippingQuery      *ShippingQuery      `json:"shipping_query,omitempty"`       // Optional. New incoming shipping query. Only for invoices with flexible price
	PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query,omitempty"`   // Optional. New incoming pre-checkout query. Contains full information about checkout
}

///////////////////////////////////////////////////////////////////////////////
// Sticker Types
///////////////////////////////////////////////////////////////////////////////

// Sticker https://core.telegram.org/bots/api/#sticker
type Sticker Dummy

// StickerSet https://core.telegram.org/bots/api#stickerset
type StickerSet Dummy

// MaskPosition https://core.telegram.org/bots/api#maskposition
type MaskPosition Dummy

///////////////////////////////////////////////////////////////////////////////
// InlineMode Types
///////////////////////////////////////////////////////////////////////////////

// InlineQuery https://core.telegram.org/bots/api/#inlinequery
type InlineQuery Dummy

// InlineQueryResult https://core.telegram.org/bots/api/#inlinequeryresult
type InlineQueryResult Dummy

// TODO: add more InlineQueryResult types

// ChosenInlineResult https://core.telegram.org/bots/api/#choseninlineresult
type ChosenInlineResult Dummy

///////////////////////////////////////////////////////////////////////////////
// Payments Types
///////////////////////////////////////////////////////////////////////////////

// LabeledPrice https://core.telegram.org/bots/api/#labeledprice
type LabeledPrice Dummy

// Invoice https://core.telegram.org/bots/api/#invoice
type Invoice Dummy

// ShippingAddress https://core.telegram.org/bots/api/#shippingaddress
type ShippingAddress Dummy

// OrderInfo https://core.telegram.org/bots/api/#orderinfo
type OrderInfo Dummy

// ShippingOption https://core.telegram.org/bots/api/#shippingoption
type ShippingOption Dummy

// SuccessfulPayment https://core.telegram.org/bots/api/#successfulpayment
type SuccessfulPayment Dummy

// ShippingQuery https://core.telegram.org/bots/api/#shippingquery
type ShippingQuery Dummy

// PreCheckoutQuery https://core.telegram.org/bots/api/#precheckoutquery
type PreCheckoutQuery Dummy

///////////////////////////////////////////////////////////////////////////////
// Games Types
///////////////////////////////////////////////////////////////////////////////

// Game https://core.telegram.org/bots/api/#game
type Game struct {
	Title       string      `json:"title"`       // Title of the game
	Description string      `json:"description"` // Description of the game
	Photo       []PhotoSize `json:"photo"`       // Photo that will be displayed in the game message in chats.

	// Optional
	Text         *string         `json:"text,omitempty"`          // Optional. Brief description of the game or high scores included in the game message. Can be automatically edited to include current high scores for the game when the bot calls setGameScore, or manually edited using editMessageText. 0-4096 characters.
	TextEntities []MessageEntity `json:"text_entities,omitempty"` // Optional. Special entities that appear in text, such as usernames, URLs, bot commands, etc.
	Anitmation   *Animation      `json:"animation,omitempty"`     // Optional. Animation that will be displayed in the game message in chats. Upload via BotFather
}

// Animation https://core.telegram.org/bots/api/#animation
type Animation struct {
	FileID string `json:"file_id"` // 	Unique file identifier

	// Optional
	Thumb    *PhotoSize `json:"thumb,omitempty"`     // Optional. Animation thumbnail as defined by sender
	FileName *string    `json:"file_name,omitempty"` // Optional. Original animation filename as defined by sender
	MIMEType *string    `json:"mime_type,omitempty"` // Optional. MIME type of the file as defined by sender
	FileSize *Integer   `json:"file_size,omitempty"` // Optional. File size
}

// CallbackGame https://core.telegram.org/bots/api/#callbackgame
type CallbackGame Placeholder // A placeholder, currently holds no information. Use BotFather to set up your game.

// GameHighScore https://core.telegram.org/bots/api/#gamehighscore
type GameHighScore struct {
	Position Integer `json:"position"` // Position in high score table for the game
	User     User    `json:"user"`     // User
	Score    Integer `json:"score"`    // Score
}
