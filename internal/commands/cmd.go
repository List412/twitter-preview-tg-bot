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

type Cmd string
