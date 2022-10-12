package telegram

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"math/rand"
	"net/url"
	"strings"
	"tweets-tg-bot/internal/clients/twitter/scrapper"
	twimg_cdn "tweets-tg-bot/internal/clients/twitter/twimg-cdn"
)

const (
	RndCmd   = "/rnd"
	HelpCmd  = "/help"
	StartCmd = "/start"
)

func (p *processor) doCmd(text string, chatId int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command: %s from: %s", text, username)

	id, err := parseTweeterUrl(text)
	if err == nil {
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

func generateText(tweet *twimg_cdn.Tweet, selfReplays []scrapper.SelfReplay) string {
	result := ""

	result += fmt.Sprintf("%s (@%s) tweeted:\n", tweet.User.Name, tweet.User.ScreenName)

	result += fmt.Sprintf("%s \n", tweet.Text)

	if tweet.Parent != nil {
		parent := tweet.Parent
		result += fmt.Sprintf("\n———\n")
		result += fmt.Sprintf("in replay to: %s (@%s):\n", parent.User.Name, parent.User.ScreenName)
		result += fmt.Sprintf("%s\n\n", parent.Text)
	}

	if len(selfReplays) > 0 {
		for _, r := range selfReplays {
			result += fmt.Sprintf("%s\n", r.Text)
		}
		result += fmt.Sprintf("\n")
	}

	result += fmt.Sprintf("%s | лайков %d", tweet.CreatedAt, tweet.FavoriteCount)

	return result
}

func (p *processor) sendTweet(chatId int, id string, username string) error {

	tweet, err := p.tw.GetTweet(id)
	if err != nil {
		return err
	}

	selfReplays, err := p.twWeb.GetTweetSelfReplays(id)
	if err != nil {
		selfReplays = nil
	}

	message := generateText(tweet, selfReplays)

	if len(tweet.Photos) == 1 {
		return p.tg.SendPhoto(chatId, message, photos(tweet)[0])
	}

	if len(tweet.Photos) >= 2 {
		return p.tg.SendPhotos(chatId, message, photos(tweet))
	}

	if tweet.Video != nil {
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

func photos(tweet *twimg_cdn.Tweet) []string {
	photos := make([]string, len(tweet.Photos))

	for i, p := range tweet.Photos {
		photos[i] = p.Url
	}
	return photos
}

func video(tweet *twimg_cdn.Tweet) string {
	for _, v := range tweet.Video.Variants {
		if v.Type == "video/mp4" {
			return v.Src
		}
	}
	return ""
}

func (p *processor) sendStart(chatId int, username string) error {
	return p.tg.SendMessage(chatId, msgHello)
}

func (p *processor) sendHelp(chatId int, username string) error {
	return p.tg.SendMessage(chatId, msgHelp)
}
