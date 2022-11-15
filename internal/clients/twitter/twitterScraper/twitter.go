package twitterScraper

import (
	twitterscraper "github.com/n0madic/twitter-scraper"
	"log"
	"sync"
	"time"
)

type Scraper struct {
	tw *twitterscraper.Scraper
}

func NewScrapper() *Scraper {
	scraper := twitterscraper.New()
	scraper.WithReplies(true)

	return &Scraper{tw: scraper}
}

func (s Scraper) GetTweet(id string) (*TweetResult, error) {

	wg := sync.WaitGroup{}
	n := 2

	var replays *PageScrapperResult
	var tweet *twitterscraper.Tweet

	errChan := make(chan error, n)

	wg.Add(n)

	go func() {
		defer wg.Done()
		r, err := s.GetTweetSelfReplays(id)
		_ = err
		replays = r
	}()

	go func() {
		defer wg.Done()
		defer log.Printf("Scrapp tweet api done %s", id)
		tw, err := s.tw.GetTweet(id)
		if err != nil {
			_ = s.tw.GetGuestToken()
			tw, err = s.tw.GetTweet(id)
		}
		if err != nil {
			errChan <- err
			return
		}
		tweet = tw
	}()

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return nil, err
		}
	}

	return &TweetResult{
		Tweet:      tweet,
		SelfReplay: replays.SelfReplay,
	}, nil
}

func (s Scraper) UpdateTokenJob() {
	ticker := time.NewTicker(30 * time.Minute)
	_ = s.tw.GetGuestToken()
	select {
	case <-ticker.C:
		err := s.tw.GetGuestToken()
		if err != nil {
			println("error while retrieving guest token")
		} else {
			println("guest token updated")
		}
	}
}
