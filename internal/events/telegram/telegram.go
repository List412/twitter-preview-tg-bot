package telegram

import (
	"context"
	"math"
	"sync"

	"github.com/pkg/errors"

	"tweets-tg-bot/internal/clients/telegram"
	"tweets-tg-bot/internal/commands"
	"tweets-tg-bot/internal/events"
	"tweets-tg-bot/internal/events/telegram/tgTypes"
)

func New(
	tgClient *telegram.Client,
	contentManager ContentManager,
	cmdParsers commands.Parsers,
	users events.UsersServiceInterface,
	botHandler string,
) *Processor {
	usersChan := make(chan string, 10)
	usersShareTweet := make(chan string, 10)
	return &Processor{
		tg:              tgClient,
		offset:          math.MaxInt64,
		contentManager:  contentManager,
		cmdParser:       cmdParsers,
		users:           users,
		usersChan:       usersChan,
		usersShareTweet: usersShareTweet,
		botHandler:      botHandler,
	}
}

var ErrUnknownEventType = errors.New("unknown event type")
var ErrUnknownMeta = errors.New("unknown meta")

type Meta struct {
	ChatId   int
	TopicId  int
	Username string
	UserId   int
	ChatName string
}

type ContentManager interface {
	GetContent(ctx context.Context, cmd commands.Cmd, urlCmd commands.ParsedCmdUrl) (tgTypes.TweetThread, error)
}

type Processor struct { //todo rename lol
	tg              *telegram.Client
	offset          int
	contentManager  ContentManager
	cmdParser       commands.Parsers
	users           events.UsersServiceInterface
	usersChan       chan string
	usersShareTweet chan string
	botHandler      string
}

func (p *Processor) Fetch(limit int) ([]commands.Event, error) {
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

func (p *Processor) Process(ctx context.Context, event commands.Event) error {
	switch event.Type {
	case commands.Message:
		return p.processMessage(ctx, event)
	case commands.Unknown:
		return ErrUnknownEventType
	}
	return nil
}

func (p *Processor) HandleUsers() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go p.handleUsersShareTweet(&wg)
	go p.handleUsers(&wg)
	wg.Wait()
}

func (p *Processor) handleUsers(wg *sync.WaitGroup) {
	defer wg.Done()
	for user := range p.usersChan {
		err := p.users.Add(context.TODO(), user)
		if err != nil {
			println(err.Error())
		}
	}
}

func (p *Processor) handleUsersShareTweet(wg *sync.WaitGroup) {
	defer wg.Done()
	for user := range p.usersShareTweet {
		err := p.users.AddShare(context.TODO(), user)
		if err != nil {
			println(err.Error())
		}
	}
}

func (p *Processor) Close() {
	close(p.usersChan)
	close(p.usersShareTweet)
}

func (p *Processor) processMessage(ctx context.Context, e commands.Event) error {
	meta, err := meta(e)
	if err != nil {
		return err
	}

	p.usersChan <- meta.Username

	if err := p.doCmd(ctx, e.Text, meta.ChatId, meta.TopicId, meta.ChatName, meta.Username, meta.UserId); err != nil {
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
		topicId := 0
		if u.Message.IsTopicMessage {
			topicId = u.Message.MessageThreadId
		}

		res.Meta = Meta{
			ChatId:   u.Message.Chat.ID,
			TopicId:  topicId,
			Username: u.Message.From.Username,
			UserId:   u.Message.From.Id,
			ChatName: u.Message.Chat.Title,
		}
		res.Lang = u.Message.From.LanguageCode
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
