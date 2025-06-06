package instagramscrapper2

import (
	"context"

	"github.com/pkg/errors"

	"github.com/list412/twitter-preview-tg-bot/internal/commands"
	"github.com/list412/twitter-preview-tg-bot/internal/events/telegram/tgTypes"
)

type Service struct {
	client *Client
}

func NewService(client *Client) *Service {
	return &Service{client: client}
}

func (s Service) GetPost(ctx context.Context, id commands.ParsedCmdUrl) (tgTypes.TweetThread, error) {
	response, err := s.client.GetVideo(ctx, id.Key)
	if err != nil {
		return tgTypes.TweetThread{}, errors.Wrap(err, "get post")
	}
	tweet, err := Map(response)
	if err != nil {
		return tgTypes.TweetThread{}, errors.Wrap(err, "convert response")
	}
	return tweet, nil
}
