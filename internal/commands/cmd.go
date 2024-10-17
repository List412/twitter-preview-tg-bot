package commands

type Type int

type Event struct {
	Type Type
	Text string
	Meta interface{}
}

const (
	Unknown Type = iota
	Message
)

const (
	RndCmd   Cmd = "/rnd"
	HelpCmd  Cmd = "/help"
	StartCmd Cmd = "/start"
	StatsCmd Cmd = "/stats"
)

const TweetCmd Cmd = "tweet"
const TikTokCmd Cmd = "tiktok"
const InstagramCmd Cmd = "insta"

type Cmd string
