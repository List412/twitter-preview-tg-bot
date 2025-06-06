package twttrapi

import (
	"context"
	"github.com/list412/twitter-preview-tg-bot/internal/events/telegram/tgTypes"
	"github.com/pkg/errors"
)

func NewService(client ClientI, mapper MapperI) *Service {
	return &Service{client: client, mapper: mapper}
}

type Service struct {
	client ClientI
	mapper MapperI
}

type MapperI interface {
	Map(parsedTweet *ParsedThread, id string) (tgTypes.TweetThread, error)
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
	tweet, err := s.mapper.Map(response, id)
	if err != nil {
		return tgTypes.TweetThread{}, errors.Wrap(err, "map")
	}
	return tweet, nil
}
