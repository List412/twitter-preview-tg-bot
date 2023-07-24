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
