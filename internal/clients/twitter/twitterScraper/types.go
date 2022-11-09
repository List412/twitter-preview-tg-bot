package twitterScraper

import twitterscraper "github.com/n0madic/twitter-scraper"

type SelfReplay struct {
	Text string
	Id   string
}

type Collab struct {
	Name       string
	ScreenName string
}

type PageScrapperResult struct {
	SelfReplay  []SelfReplay
	CollabUsers []Collab
}

type TweetResult struct {
	Tweet      *twitterscraper.Tweet
	SelfReplay []SelfReplay
}
