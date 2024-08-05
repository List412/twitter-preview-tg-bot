package telegram

import "tweets-tg-bot/internal/events/telegram/tgTypes"

type UpdateResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	Id      int              `json:"update_id"`
	Message *IncomingMessage `json:"message"`
}

type IncomingMessage struct {
	Text string `json:"text"`
	From User   `json:"from"`
	Chat Chat   `json:"chat"`
}

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type Chat struct {
	ID int `json:"id"`
}

type MediaObject struct {
	Type      string `json:"type"`
	Media     string `json:"media"`
	Caption   string `json:"caption,omitempty"`
	ParseMode string `json:"parse_mode,omitempty"`
}

type MediaForEncoding struct {
	Media           []tgTypes.MediaObject
	MediaType       string
	ForceNeedUpload bool
}

type Button struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]Button `json:"inline_keyboard"`
}
