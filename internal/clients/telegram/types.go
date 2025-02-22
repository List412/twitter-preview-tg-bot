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
	Id           int    `json:"id"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

type Chat struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
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

type ChatFullInfo struct {
	Ok     bool `json:"ok"`
	Result struct {
		FirstName       string   `json:"first_name"`
		LastName        string   `json:"last_name"`
		Username        string   `json:"username"`
		CanSendGift     bool     `json:"can_send_gift"`
		ActiveUsernames []string `json:"active_usernames"`
		Photo           struct {
			SmallFileId       string `json:"small_file_id"`
			SmallFileUniqueId string `json:"small_file_unique_id"`
			BigFileId         string `json:"big_file_id"`
			BigFileUniqueId   string `json:"big_file_unique_id"`
		} `json:"photo"`
		Id                int64  `json:"id"`
		Title             string `json:"title"`
		Type              string `json:"type"`
		HasVisibleHistory bool   `json:"has_visible_history"`
		Permissions       struct {
			CanSendMessages       bool `json:"can_send_messages"`
			CanSendMediaMessages  bool `json:"can_send_media_messages"`
			CanSendAudios         bool `json:"can_send_audios"`
			CanSendDocuments      bool `json:"can_send_documents"`
			CanSendPhotos         bool `json:"can_send_photos"`
			CanSendVideos         bool `json:"can_send_videos"`
			CanSendVideoNotes     bool `json:"can_send_video_notes"`
			CanSendVoiceNotes     bool `json:"can_send_voice_notes"`
			CanSendPolls          bool `json:"can_send_polls"`
			CanSendOtherMessages  bool `json:"can_send_other_messages"`
			CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`
			CanChangeInfo         bool `json:"can_change_info"`
			CanInviteUsers        bool `json:"can_invite_users"`
			CanPinMessages        bool `json:"can_pin_messages"`
			CanManageTopics       bool `json:"can_manage_topics"`
		} `json:"permissions"`
		JoinToSendMessages bool `json:"join_to_send_messages"`
		MaxReactionCount   int  `json:"max_reaction_count"`
		AccentColorId      int  `json:"accent_color_id"`
	} `json:"result"`
}

type ChatAdmins struct {
	Ok     bool `json:"ok"`
	Result []struct {
		User struct {
			Id           int    `json:"id"`
			IsBot        bool   `json:"is_bot"`
			FirstName    string `json:"first_name"`
			LastName     string `json:"last_name"`
			Username     string `json:"username"`
			LanguageCode string `json:"language_code"`
		} `json:"user"`
		Status      string `json:"status"`
		IsAnonymous bool   `json:"is_anonymous"`
	} `json:"result"`
}
