package twitterScraper

import (
	twitterscraper "github.com/n0madic/twitter-scraper"
)

type Scraper struct {
	tw *twitterscraper.Scraper
}

func NewScrapper() *Scraper {
	scraper := twitterscraper.New()
	scraper.WithReplies(true)

	return &Scraper{tw: scraper}
}

func (s Scraper) GetTweet(id string) (*twitterscraper.Tweet, error) {

	return s.tw.GetTweet(id)
}
