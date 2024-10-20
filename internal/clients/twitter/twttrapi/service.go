package twttrapi

import (
	"context"
	"github.com/pkg/errors"
	"tweets-tg-bot/internal/events/telegram/tgTypes"
)

func NewService(client ClientI) *Service {
	return &Service{client: client}
}

type Service struct {
	client ClientI
}

type ClientI interface {
	RapidApiClient
	GetTweet(ctx context.Context, id string) (*ParsedThread, error)
}

func (s Service) GetTweet(ctx context.Context, id string) (tgTypes.TweetThread, error) {
	response, err := s.client.GetTweet(ctx, id)
	if err != nil {
		return tgTypes.TweetThread{}, errors.Wrap(err, "get tweet")
	}
	tweet, err := Map(response, id)
	if err != nil {
		return tgTypes.TweetThread{}, errors.Wrap(err, "map")
	}
	return tweet, nil
}
