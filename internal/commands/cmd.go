package commands

type Type int

type Event struct {
	Type Type
	Text string
	Lang string
	Meta interface{}
}

const (
	Unknown Type = iota
	Message
)

const (
	RndCmd    Cmd = "/rnd"
	HelpCmd   Cmd = "/help"
	StartCmd  Cmd = "/start"
	StatsCmd  Cmd = "/stats"
	LeaveChat Cmd = "/leavechat"
	ChatInfo  Cmd = "/chatinfo"
)

const TweetCmd Cmd = "tweet"
const TikTokCmd Cmd = "tiktok"
const InstagramCmd Cmd = "insta"

type Cmd string

type ParsedCmdUrl struct {
	OriginalUrl string
	Key         string
	StrippedUrl string
}
