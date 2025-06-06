package socialapi1instagram

import (
	"context"
	"github.com/list412/twitter-preview-tg-bot/internal/commands"
	"github.com/list412/twitter-preview-tg-bot/internal/events/telegram/tgTypes"
	"github.com/pkg/errors"
)

type Service struct {
	client *Client
}

func NewService(client *Client) *Service {
	return &Service{client: client}
}

func (s Service) GetPost(ctx context.Context, id commands.ParsedCmdUrl) (tgTypes.TweetThread, error) {
	response, err := s.client.GetVideo(ctx, id.StrippedUrl)
	if err != nil {
		return tgTypes.TweetThread{}, errors.Wrap(err, "get post")
	}
	tweet, err := Map(response)
	if err != nil {
		return tgTypes.TweetThread{}, errors.Wrap(err, "convert response")
	}
	return tweet, nil
}
