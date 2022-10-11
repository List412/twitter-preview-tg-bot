package twitter

type Tweet struct {
	Data Data `json:"data"`
}

type Data struct {
	Text                string   `json:"text"`
	EditHistoryTweetIds []string `json:"edit_history_tweet_ids"`
	AuthorId            string   `json:"author_id"`
	Attachments         struct {
		MediaKeys []string `json:"media_keys"`
	} `json:"attachments"`
	Id string `json:"id"`
}

type Includes struct {
	Media []struct {
		MediaKey string `json:"media_key"`
		Type     string `json:"type"`
	} `json:"media"`
}
