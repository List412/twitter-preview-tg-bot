package telegram

import (
	"context"
	"github.com/pkg/errors"
	"math"
	"sync"
	"time"
	"tweets-tg-bot/internal/clients/telegram"
	"tweets-tg-bot/internal/commands"
	"tweets-tg-bot/internal/events"
)

func New(tgClient *telegram.Client, twitterService TwitterService, users events.UsersServiceInterface) *processor {
	usersChan := make(chan string, 10)
	usersShareTweet := make(chan string, 10)
	return &processor{
		tg:              tgClient,
		offset:          math.MaxInt64,
		twitterService:  twitterService,
		users:           users,
		usersChan:       usersChan,
		usersShareTweet: usersShareTweet,
	}
}

var ErrUnknownEventType = errors.New("unknown event type")
var ErrUnknownMeta = errors.New("unknown meta")

type Meta struct {
	ChatId   int
	Username string
	UserId   int
}

type TwitterService interface {
	GetTweet(id string) (Tweet, error)
}

type Media struct {
	Photos []string
	Videos []string
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

type processor struct { //todo rename lol
	tg              *telegram.Client
	offset          int
	twitterService  TwitterService
	users           events.UsersServiceInterface
	usersChan       chan string
	usersShareTweet chan string
}

func (p *processor) Fetch(limit int) ([]commands.Event, error) {
	updates, err := p.tg.Updates(p.offset, limit)

	if err != nil {
		return nil, err
	}

	if len(updates) == 0 {
		return nil, nil
	}

	res := make([]commands.Event, len(updates))

	for i, u := range updates {
		res[i] = event(u)
	}

	p.offset = updates[len(updates)-1].Id + 1

	return res, nil
}

func (p *processor) Process(event commands.Event) error {
	switch event.Type {
	case commands.Message:
		return p.processMessage(event)
	case commands.Unknown:
		return ErrUnknownEventType
	}
	return nil
}

func (p *processor) HandleUsers() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go p.handleUsersShareTweet(&wg)
	go p.handleUsers(&wg)
	wg.Wait()
}

func (p *processor) handleUsers(wg *sync.WaitGroup) {
	defer wg.Done()
	for user := range p.usersChan {
		err := p.users.Add(context.TODO(), user)
		if err != nil {
			println(err.Error())
		}
	}
}

func (p *processor) handleUsersShareTweet(wg *sync.WaitGroup) {
	defer wg.Done()
	for user := range p.usersShareTweet {
		err := p.users.AddShare(context.TODO(), user)
		if err != nil {
			println(err.Error())
		}
	}
}

func (p *processor) Close() {
	close(p.usersChan)
	close(p.usersShareTweet)
}

func (p *processor) processMessage(e commands.Event) error {
	meta, err := meta(e)
	if err != nil {
		return err
	}

	p.usersChan <- meta.Username

	if err := p.doCmd(e.Text, meta.ChatId, meta.Username, meta.UserId); err != nil {
		return err
	}

	return nil
}

func meta(e commands.Event) (Meta, error) {
	res, ok := e.Meta.(Meta)
	if !ok {
		return Meta{}, ErrUnknownMeta
	}
	return res, nil
}

func event(u telegram.Update) commands.Event {
	messageType := fetchType(u)

	res := commands.Event{
		Type: messageType,
		Text: fetchText(u),
	}

	if res.Type == commands.Message {
		res.Meta = Meta{
			ChatId:   u.Message.Chat.ID,
			Username: u.Message.From.Username,
			UserId:   u.Message.From.Id,
		}
	}

	return res
}

func fetchText(u telegram.Update) string {
	if u.Message == nil {
		return ""
	}
	return u.Message.Text
}

func fetchType(u telegram.Update) commands.Type {
	if u.Message == nil {
		return commands.Unknown
	}
	return commands.Message
}
