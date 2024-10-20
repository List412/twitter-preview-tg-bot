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
	Id string `json:"id"`
}

type Variant struct {
	ContentType string `json:"content_type"`
	Url         string `json:"url"`
	Bitrate     int    `json:"bitrate,omitempty"`
}
