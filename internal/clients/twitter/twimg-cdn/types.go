package twimg_cdn

import "time"

type Tweet struct {
	Lang              string    `json:"lang"`
	FavoriteCount     int       `json:"favorite_count"` // likes
	CreatedAt         time.Time `json:"created_at"`
	Entities          Entities  `json:"entities"`
	IdStr             string    `json:"id_str"`
	Text              string    `json:"text"`
	User              User      `json:"user"`
	Photos            []Photos  `json:"photos"`
	Video             *Video    `json:"video"`
	ConversationCount int       `json:"conversation_count"`
	NewsActionType    string    `json:"news_action_type"`
	Parent            *Parent   `json:"parent"`
}

type User struct {
	IdStr                string `json:"id_str"`
	Name                 string `json:"name"`
	ProfileImageUrlHttps string `json:"profile_image_url_https"`
	ScreenName           string `json:"screen_name"`
	Verified             bool   `json:"verified"`
}

type Parent struct {
	Lang          string    `json:"lang"`
	ReplyCount    int       `json:"reply_count"`
	RetweetCount  int       `json:"retweet_count"`
	FavoriteCount int       `json:"favorite_count"`
	CreatedAt     time.Time `json:"created_at"`
	Entities      Entities  `json:"entities"`
	IdStr         string    `json:"id_str"`
	Text          string    `json:"text"`
	User          User      `json:"user"`
}

type Entities struct {
	Urls  []Urls  `json:"urls"`
	Media []Media `json:"media"`
}

type Urls struct {
	DisplayUrl  string `json:"display_url"`
	ExpandedUrl string `json:"expanded_url"`
	Url         string `json:"url"`
}

type Media struct {
	DisplayUrl  string `json:"display_url"`
	ExpandedUrl string `json:"expanded_url"`
	Url         string `json:"url"`
}

type Photos struct {
	ExpandedUrl string `json:"expandedUrl"` // twitter url
	Url         string `json:"url"`         // cdn url
	Width       int    `json:"width"`
	Height      int    `json:"height"`
}

type Video struct {
	DurationMs int        `json:"durationMs"`
	Poster     string     `json:"poster"`
	Variants   []Variants `json:"variants"`
	VideoId    VideoId    `json:"videoId"`
	ViewCount  int        `json:"viewCount"`
}

type Variants struct {
	Type string `json:"type"` // take first video/mp4 Src
	Src  string `json:"src"`
}

type VideoId struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}
