package telegram

import "strconv"

// Recipient is basically any possible endpoint you can send
// messages to. It's usually a distinct user or a chat.
type Recipient interface {
	// ID of user or group chat, @Username for channel
	Destination() string
}

// User object represents a Telegram user, bot
//
// object represents a group chat if Title is empty.
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`

	LastName string `json:"last_name"`
	Username string `json:"username"`
}

// Destination is internal user ID.
func (u User) Destination() string {
	return strconv.Itoa(u.ID)
}

// Chat object represents a Telegram user, bot or group chat.
//
// Target of chat, can be either “private”, “group”, "supergroup" or “channel”
type Chat struct {
	ID int64 `json:"id"`

	// See telebot.ChatType and consts.
	Type ChatType `json:"type"`

	// Won't be there for ChatPrivate.
	Title string `json:"title"`

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

// Destination is internal chat ID.
func (c Chat) Destination() string {
	ret := "@" + c.Username
	if c.Type != "channel" {
		ret = strconv.FormatInt(c.ID, 10)
	}
	return ret
}

// IsGroupChat returns true if chat object represents a group chat.
func (c Chat) IsGroupChat() bool {
	return c.Type != "private"
}

// Update object represents an incoming update.
type Update struct {
	ID      int64    `json:"update_id"`
	Payload *Message `json:"message"`

	// optional
	Callback *Callback `json:"callback_query"`
	Query    *Query    `json:"inline_query"`
}

// KeyboardButton represents a button displayed on in a message.
type KeyboardButton struct {
	Text        string `json:"text"`
	URL         string `json:"url,omitempty"`
	Data        string `json:"callback_data,omitempty"`
	InlineQuery string `json:"switch_inline_query,omitempty"`
}

// Photo object represents a photo with caption.
type Photo struct {
	File
	Data    string
	Caption string
}

// Thumbnail object represents an image/sticker of a particular size.
type Thumbnail struct {
	File

	Width  int `json:"width"`
	Height int `json:"height"`
}

// Location object represents geographic position.
type Location struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}
