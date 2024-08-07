package teledau

type MessageRequest struct {
	ChatId                string      `json:"chat_id,omitempty"`
	Text                  string      `json:"text,omitempty"`
	ParseMode             string      `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool        `json:"disable_web_page_preview,omitempty"`
	ReplyMarkup           interface{} `json:"reply_markup,omitempty"`
	Photo                 []Photo     `json:"photo,omitempty"`
}
type EditMessageRequest struct {
	MessageId             int         `json:"message_id,omitempty"`
	ChatId                string      `json:"chat_id,omitempty"`
	Text                  string      `json:"text,omitempty"`
	ParseMode             string      `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool        `json:"disable_web_page_preview,omitempty"`
	ReplyMarkup           interface{} `json:"reply_markup,omitempty"`
	Photo                 []Photo     `json:"photo,omitempty"`
}
type EditCaptionRequest struct {
	MessageId             int         `json:"message_id,omitempty"`
	ChatId                string      `json:"chat_id,omitempty"`
	Text                  string      `json:"caption,omitempty"`
	ParseMode             string      `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool        `json:"disable_web_page_preview,omitempty"`
	ReplyMarkup           interface{} `json:"reply_markup,omitempty"`
	Photo                 []Photo     `json:"photo,omitempty"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text         string     `json:"text"`
	URL          string     `json:"url,omitempty"`
	CallbackData string     `json:"callback_data,omitempty"`
	WebApp       WebAppInfo `json:"web_app,omitempty"`
}

type WebAppInfo struct {
	URL string `json:"url"`
}

type KeyboardButton struct {
	Text           string `json:"text"`
	RequestContact bool   `json:"request_contact"`
}

type ReplyKeyboardMarkup struct {
	Keyboard        [][]KeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool               `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard bool               `json:"one_time_keyboard,omitempty"`
}

type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective,omitempty"`
}

type ForceReply struct {
	ForceReply bool `json:"force_reply"`
	Selective  bool `json:"selective,omitempty"`
}

type CreateChatInviteLinkRequest struct {
	ChatID string `json:"chat_id"` // Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Name   string `json:"name"`    // Title of the invite link
}
