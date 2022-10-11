package telegram

import (
	"github.com/pkg/errors"
	"math"
	"tweets-tg-bot/internal/clients/telegram"
	twimg_cdn "tweets-tg-bot/internal/clients/twitter/twimg-cdn"
	"tweets-tg-bot/internal/events"
)

func New(tgClient *telegram.Client, twClient *twimg_cdn.Client) *processor {
	return &processor{
		tg:     tgClient,
		offset: math.MaxInt64,
		tw:     twClient,
	}
}

var ErrUnknownEventType = errors.New("unknown event type")
var ErrUnknownMeta = errors.New("unknown meta")

type Meta struct {
	ChatId   int
	Username string
}

type processor struct { //todo rename lol
	tg     *telegram.Client
	offset int
	tw     *twimg_cdn.Client
}

func (p *processor) Fetch(limit int) ([]events.Event, error) {
	updates, err := p.tg.Updates(p.offset, limit)

	if err != nil {
		return nil, err
	}

	if len(updates) == 0 {
		return nil, nil
	}

	res := make([]events.Event, len(updates))

	for i, u := range updates {
		res[i] = event(u)
	}

	p.offset = updates[len(updates)-1].Id + 1

	return res, nil
}

func (p *processor) Process(event events.Event) error {
	switch event.Type {
	case events.Message:
		return p.processMessage(event)
	case events.Unknown:
		return ErrUnknownEventType
	}
	return nil
}

func (p *processor) processMessage(e events.Event) error {
	meta, err := meta(e)
	if err != nil {
		return err
	}

	if err := p.doCmd(e.Text, meta.ChatId, meta.Username); err != nil {
		return err
	}

	return nil
}

func meta(e events.Event) (Meta, error) {
	res, ok := e.Meta.(Meta)
	if !ok {
		return Meta{}, ErrUnknownMeta
	}
	return res, nil
}

func event(u telegram.Update) events.Event {
	messageType := fetchType(u)

	res := events.Event{
		Type: messageType,
		Text: fetchText(u),
	}

	if res.Type == events.Message {
		res.Meta = Meta{
			ChatId:   u.Message.Chat.ID,
			Username: u.Message.From.Username,
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

func fetchType(u telegram.Update) events.Type {
	if u.Message == nil {
		return events.Unknown
	}
	return events.Message
}
