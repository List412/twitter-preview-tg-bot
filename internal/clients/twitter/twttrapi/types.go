package twttrapi

type Variant struct {
	ContentType string `json:"content_type"`
	Url         string `json:"url"`
}

type TweetData struct {
	ConversationIdStr string `json:"conversation_id_str"`
	CreatedAt         string `json:"created_at"`
	DisplayTextRange  []int  `json:"display_text_range"`

	ExtendedEntities struct {
		Media []struct {
			DisplayUrl    string `json:"display_url"`
			ExpandedUrl   string `json:"expanded_url"`
			MediaUrlHttps string `json:"media_url_https"`
			Type          string `json:"type"`
			MediaKey      string `json:"media_key"`
			VideoInfo     struct {
				Variants []Variant `json:"variants"`
			} `json:"video_info"`
		} `json:"media"`
	} `json:"extended_entities"`
	FavoriteCount int    `json:"favorite_count"`
	Favorited     bool   `json:"favorited"`
	FullText      string `json:"full_text"`
	IsQuoteStatus bool   `json:"is_quote_status"`
	Lang          string `json:"lang"`

	QuoteCount   int    `json:"quote_count"`
	ReplyCount   int    `json:"reply_count"`
	RetweetCount int    `json:"retweet_count"`
	Retweeted    bool   `json:"retweeted"`
	UserIdStr    string `json:"user_id_str"`
}

type UserData struct {
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
}

type NoteResult struct {
}

type Tweet struct {
	Tweet *Tweet `json:"tweet"`
	Core  *struct {
		UserResult struct {
			Result struct {
				Legacy UserData `json:"legacy"`
			} `json:"result"`
		} `json:"user_result"`
	} `json:"core"`
	Legacy *TweetData `json:"legacy"`

	NoteTweet *struct {
		NoteTweetResults struct {
			Result struct {
				Text string `json:"text"`
			} `json:"result"`
		} `json:"note_tweet_results"`
	} `json:"note_tweet"`

	ViewCountInfo *struct {
		Count string `json:"count"`
		State string `json:"state"`
	} `json:"view_count_info"`
}

type ParsedTweet struct {
	Data struct {
		TweetResult struct {
			Result *Tweet `json:"result"`
		} `json:"tweet_result"`
	} `json:"data"`
	Errors []struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
	Error *string `json:"error"`
}
