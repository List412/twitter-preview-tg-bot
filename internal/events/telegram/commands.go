package telegram

import (
	"context"
	"fmt"
	twitterscraper "github.com/n0madic/twitter-scraper"
	"github.com/pkg/errors"
	"log"
	"math/rand"
	"net/url"
	"strings"
	"time"
	"tweets-tg-bot/internal/clients/twitter/twitterScraper"
	"tweets-tg-bot/internal/commands"
)

var AllCommands = []commands.Cmd{
	commands.RndCmd, commands.HelpCmd, commands.StartCmd, commands.StatsCmd,
}

var ErrorUnknownCommand = errors.New("unknown command")

func (p *processor) doCmd(text string, chatId int, username string, userId int) error {
	text = strings.TrimSpace(text)

	cmd, err := parseCmd(text)
	if errors.Is(err, ErrorUnknownCommand) {
		return nil
	}
	if err != nil {
		return err
	}

	defer func() {
		p.users.Command(cmd, username)
	}()

	switch cmd {
	case commands.TweetCmd:
		id, err := parseTweeterUrl(text)
		if err != nil {
			return nil
		}
		log.Printf("got new command: %s from: %s", text, username)
		return p.sendTweet(chatId, id, username)
	case commands.StartCmd:
		return p.sendStart(chatId, username)
	case commands.HelpCmd:
		return p.sendHelp(chatId, username)
	case commands.RndCmd:
		return p.sendRandom(chatId, username)
	case commands.StatsCmd:
		return p.sendStats(chatId, userId)
	default:
		return nil
	}
}

func generateText(tweet *twitterscraper.Tweet, replays []twitterScraper.SelfReplay) string {
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

	if len(replays) > 0 {
		for _, r := range replays {
			result += fmt.Sprintf("%s\n", r.Text)
		}
		result += fmt.Sprintf("\n")
	}

	twTime := time.Unix(tweet.Timestamp, 0).Format("15:04 2 Jan 2006")

	result += fmt.Sprintf("%s | 💙%d| replies %d| retweets %d", twTime, tweet.Likes, tweet.Replies, tweet.Retweets)

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

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func (p *processor) sendTweet(chatId int, id string, username string) error {
	defer timeTrack(time.Now(), "sendTweet")

	tweetResult, err := p.tw.GetTweet(id)
	if err != nil {
		return err
	}

	tweet := tweetResult.Tweet

	message := generateText(tweet, tweetResult.SelfReplay)

	tweet.Photos = append(tweet.Photos, photoFromQuoted(tweet.QuotedStatus, tweet.RetweetedStatus, tweet.InReplyToStatus)...)

	messages, err := chunkString(message, 1024)
	if err != nil {
		return err
	}

	for _, m := range messages {
		if len(tweet.Videos) > 0 {
			err = p.tg.SendVideo(chatId, m, video(tweet))
			if err != nil {
				return err
			}
			tweet.Videos = nil
			continue
		}

		if len(tweet.Photos) == 1 {
			err = p.tg.SendPhoto(chatId, m, photos(tweet)[0])
			if err != nil {
				return err
			}
			tweet.Photos = nil
			continue
		}

		if len(tweet.Photos) >= 2 {
			err = p.tg.SendPhotos(chatId, m, photos(tweet))
			if err != nil {
				return err
			}
			tweet.Photos = nil
			continue
		}

		if err := p.tg.SendMessage(chatId, m); err != nil {
			return err
		}
		continue
	}

	return nil
}

func chunkString(s string, chunkSize int) ([]string, error) {
	//regex, err := regexp.Compile(".{1,25}\\b|.{1,25}")
	//if err != nil {
	//	return nil, err
	//}
	//
	//chunks := regex.FindAllString(s, -1)
	var chunks []string
	runes := []rune(s)

	if len(runes) == 0 {
		return []string{s}, nil
	}

	for i := 0; i < len(runes); i += chunkSize {
		nn := i + chunkSize
		if nn > len(runes) {
			nn = len(runes)
		}
		chunks = append(chunks, string(runes[i:nn]))
	}
	return chunks, nil
}

func (p *processor) sendRandom(chatId int, username string) error {
	n := rand.Intn(100)
	if err := p.tg.SendMessage(chatId, fmt.Sprintf("random %d", n)); err != nil {
		return err
	}

	return nil
}

func (p *processor) sendStats(id int, userId int) error {
	isAdmin, err := p.users.IsAdmin(userId)
	if err != nil {
		return err
	}
	if !isAdmin {
		return nil
	}

	ctx := context.TODO()
	count, err := p.users.Count(ctx)
	if err != nil {
		return err
	}

	countShares, err := p.users.CountShare(ctx)
	if err != nil {
		return err
	}

	comandsStat, err := p.users.CommandsStat(ctx)
	if err != nil {
		return err
	}

	message := fmt.Sprintf("Users: %d \nUsers who share tweets: %d \n", count, countShares)

	for k, v := range comandsStat {
		message += fmt.Sprintf("\n%s: %d", k, v)
	}

	if err := p.tg.SendMessage(id, message); err != nil {
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

func parseCmd(text string) (commands.Cmd, error) {
	if _, err := parseTweeterUrl(text); err == nil {
		return commands.TweetCmd, nil
	}

	for _, cmd := range AllCommands {
		if len(text) < len(cmd) {
			continue
		}

		cmdPart := text[:len(cmd)]

		if string(cmd) == cmdPart {
			return cmd, nil
		}
	}

	return "", ErrorUnknownCommand
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
