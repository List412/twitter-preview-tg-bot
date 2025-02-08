package twitterapi45

type Response struct {
	Likes          int    `json:"likes"`
	CreatedAt      string `json:"created_at"`
	Text           string `json:"text"`
	Retweets       int    `json:"retweets"`
	Bookmarks      int    `json:"bookmarks"`
	Quotes         int    `json:"quotes"`
	Replies        int    `json:"replies"`
	Lang           string `json:"lang"`
	ConversationId string `json:"conversation_id"`
	Author         struct {
		RestId       string `json:"rest_id"`
		Name         string `json:"name"`
		ScreenName   string `json:"screen_name"`
		Image        string `json:"image"`
		BlueVerified bool   `json:"blue_verified"`
	} `json:"author"`
	Media struct {
		Photo []struct {
			MediaUrlHttps string `json:"media_url_https"`
			Id            string `json:"id"`
		} `json:"photo"`
		Video []struct {
			MediaUrlHttps string    `json:"media_url_https"`
			Variants      []Variant `json:"variants"`
		} `json:"video"`
	} `json:"media"`
	Id       string `json:"id"`
	Entities struct {
		Media []struct {
			DisplayUrl    string `json:"display_url"`
			ExpandedUrl   string `json:"expanded_url"`
			IdStr         string `json:"id_str"`
			Indices       []int  `json:"indices"`
			MediaKey      string `json:"media_key"`
			MediaUrlHttps string `json:"media_url_https"`
			Type          string `json:"type"`
			Url           string `json:"url"`
			VideoInfo     struct {
				AspectRatio    []int     `json:"aspect_ratio"`
				DurationMillis int       `json:"duration_millis"`
				Variants       []Variant `json:"variants"`
			} `json:"video_info"`
			MediaResults struct {
				Result struct {
					MediaKey string `json:"media_key"`
				} `json:"result"`
			} `json:"media_results"`
		} `json:"media"`
		Symbols      []interface{} `json:"symbols"`
		Urls         []interface{} `json:"urls"`
		UserMentions []interface{} `json:"user_mentions"`
	} `json:"entities"`
}

type Variant struct {
	ContentType string `json:"content_type"`
	Url         string `json:"url"`
	Bitrate     int    `json:"bitrate,omitempty"`
}
