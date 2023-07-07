package twttrapi

type ParsedTweet struct {
	Data struct {
		TweetResult struct {
			Result struct {
				Core struct {
					UserResult struct {
						Result struct {
							Legacy struct {
								Name       string `json:"name"`
								ScreenName string `json:"screen_name"`
							} `json:"legacy"`
						} `json:"result"`
					} `json:"user_result"`
				} `json:"core"`
				Legacy struct {
					ConversationIdStr string `json:"conversation_id_str"`
					CreatedAt         string `json:"created_at"`
					DisplayTextRange  []int  `json:"display_text_range"`

					ExtendedEntities struct {
						Media []struct {
							DisplayUrl  string `json:"display_url"`
							ExpandedUrl string `json:"expanded_url"`
							Type        string `json:"type"`
							VideoInfo   struct {
								Variants []struct {
									ContentType string `json:"content_type"`
									Url         string `json:"url"`
								} `json:"variants"`
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
				} `json:"legacy"`

				NoteTweet *struct {
					NoteTweetResults struct {
						Result struct {
							Text string `json:"text"`
						} `json:"result"`
					} `json:"note_tweet_results"`
				} `json:"note_tweet"`

				ViewCountInfo struct {
					Count string `json:"count"`
					State string `json:"state"`
				} `json:"view_count_info"`
			} `json:"result"`
		} `json:"tweet_result"`
	} `json:"data"`
}
