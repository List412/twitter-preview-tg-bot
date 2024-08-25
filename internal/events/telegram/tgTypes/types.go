package tgTypes

import "time"

type Media struct {
	Photos []MediaObject
	Videos []MediaObject
}

type MediaObject struct {
	Name       string
	Url        string
	Data       []byte
	NeedUpload bool
}

type Tweet struct {
	Media    Media
	Text     string
	UserName string
	UserId   string
	Time     time.Time
	Likes    int
	Retweets int
	Quotes   int
	Views    string
	Replies  int
}

type TweetContent struct {
	Text  string
	Media Media
}

type UserNote struct {
	Text  string
	Title string
}

type TweetThread struct {
	UserName string
	UserId   string
	Time     time.Time
	Likes    int
	Retweets int
	Quotes   int
	Views    string
	Replies  int
	Tweets   []TweetContent
	UserNote UserNote
}
