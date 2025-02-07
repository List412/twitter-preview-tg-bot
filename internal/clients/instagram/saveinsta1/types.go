package saveinsta1

type ParsedPost struct {
	Result []struct {
		Urls []struct {
			Url       string `json:"url"`
			Name      string `json:"name"`
			Extension string `json:"extension"`
		} `json:"urls"`
		Meta struct {
			Title        string `json:"title"`
			SourceUrl    string `json:"sourceUrl"`
			Shortcode    string `json:"shortcode"`
			CommentCount int    `json:"commentCount"`
			LikeCount    int    `json:"likeCount"`
			TakenAt      int    `json:"takenAt"`
			Comments     []struct {
				Text     string `json:"text"`
				Username string `json:"username"`
			} `json:"comments"`
		} `json:"meta"`
		PictureUrl        string `json:"pictureUrl"`
		PictureUrlWrapped string `json:"pictureUrlWrapped"`
		Service           string `json:"service"`
	} `json:"result"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}
