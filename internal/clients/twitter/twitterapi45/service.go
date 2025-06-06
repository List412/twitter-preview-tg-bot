package twitterapi45

import (
	"context"
	"github.com/list412/tweets-tg-bot/internal/events/telegram/tgTypes"
	"github.com/pkg/errors"
)

func NewService(client *Client) *Service {
	return &Service{client: client}
}

type Service struct {
	client *Client
}

func (s Service) GetTweet(ctx context.Context, id string) (tgTypes.TweetThread, error) {
	response, err := s.client.GetTweet(ctx, id)
	if err != nil {
		return tgTypes.TweetThread{}, errors.Wrap(err, "get tweet")
	}
	tweet, err := Map(response)
	if err != nil {
		return tgTypes.TweetThread{}, errors.Wrap(err, "map")
	}
	return tweet, nil
}
