package telegram

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

type Photo struct {
	Type    string `json:"type"`
	Media   string `json:"media"`
	Caption string `json:"caption,omitempty"`
}

type Button struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]Button `json:"inline_keyboard"`
}
