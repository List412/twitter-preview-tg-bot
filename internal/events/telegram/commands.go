package telegram

import (
	"fmt"
	twitterscraper "github.com/n0madic/twitter-scraper"
	"github.com/pkg/errors"
	"log"
	"math/rand"
	"net/url"
	"strings"
	"time"
)

const (
	RndCmd   = "/rnd"
	HelpCmd  = "/help"
	StartCmd = "/start"
)

func (p *processor) doCmd(text string, chatId int, username string) error {
	text = strings.TrimSpace(text)

	id, err := parseTweeterUrl(text)
	if err == nil {
		log.Printf("got new command: %s from: %s", text, username)
		return p.sendTweet(chatId, id, username)
	}

	switch text {
	case StartCmd:
		return p.sendStart(chatId, username)
	case HelpCmd:
		return p.sendHelp(chatId, username)
	case RndCmd:
		return p.sendRandom(chatId, username)
	default:
		return nil
	}
}

func generateText(tweet *twitterscraper.Tweet) string {
	result := ""

	result += fmt.Sprintf("%s tweeted:\n", tweet.Username)

	result += fmt.Sprintf("%s \n", tweet.Text)

	if tweet.InReplyToStatus != nil {
		result += addInReplayTo(tweet.InReplyToStatus)
	}

	if tweet.QuotedStatus != nil {
		result += addInReplayTo(tweet.QuotedStatus)
	}

	if tweet.RetweetedStatus != nil {
		result += addInReplayTo(tweet.RetweetedStatus)
	}

	twTime := time.Unix(tweet.Timestamp, 0).Format("15:04 2 Jan 2006")

	result += fmt.Sprintf("%s | 💙%d", twTime, tweet.Likes)

	return result
}

func addInReplayTo(tweet *twitterscraper.Tweet) string {
	result := fmt.Sprintf("\n———\n")
	result += fmt.Sprintf("in replay to: %s:\n", tweet.Username)
	result += fmt.Sprintf("%s\n\n", tweet.Text)
	return result
}

func photoFromQuoted(tweets ...*twitterscraper.Tweet) []string {
	var photos []string
	for _, tweet := range tweets {
		if tweet != nil && len(tweet.Photos) > 0 {
			photos = append(photos, tweet.Photos...)
		}
	}
	return photos
}

func (p *processor) sendTweet(chatId int, id string, username string) error {

	tweet, err := p.tw.GetTweet(id)
	if err != nil {
		return err
	}

	message := generateText(tweet)

	tweet.Photos = append(tweet.Photos, photoFromQuoted(tweet.QuotedStatus, tweet.RetweetedStatus, tweet.InReplyToStatus)...)

	if len(tweet.Photos) == 1 {
		return p.tg.SendPhoto(chatId, message, photos(tweet)[0])
	}

	if len(tweet.Photos) >= 2 {
		return p.tg.SendPhotos(chatId, message, photos(tweet))
	}

	if len(tweet.Videos) > 0 {
		return p.tg.SendVideo(chatId, message, video(tweet))
	}

	if err := p.tg.SendMessage(chatId, message); err != nil {
		return err
	}

	return nil
}

func (p *processor) sendRandom(chatId int, username string) error {
	n := rand.Intn(100)
	if err := p.tg.SendMessage(chatId, fmt.Sprintf("random %d", n)); err != nil {
		return err
	}

	return nil
}

func parseTweeterUrl(text string) (string, error) {
	u, err := url.Parse(text)

	if err != nil {
		return "", err
	}

	if u.Host != "twitter.com" {
		return "", errors.New("not a twitter url")
	}

	path := strings.Split(strings.TrimLeft(u.Path, "/"), "/")
	if len(path) != 3 {
		return "", errors.New("url don't have id")
	}

	if path[2] == "" {
		return "", errors.New("id in url empty")
	}
	return path[2], nil
}

func photos(tweet *twitterscraper.Tweet) []string {
	return tweet.Photos
}

func video(tweet *twitterscraper.Tweet) string {
	if len(tweet.Videos) > 0 {
		return tweet.Videos[0].URL
	}
	return ""
}

func (p *processor) sendStart(chatId int, username string) error {
	return p.tg.SendMessage(chatId, msgHello)
}

func (p *processor) sendHelp(chatId int, username string) error {
	return p.tg.SendMessage(chatId, msgHelp)
}
