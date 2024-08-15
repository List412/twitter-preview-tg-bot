package twttrapi

import (
	"context"
	"tweets-tg-bot/internal/events/telegram/tgTypes"
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
		return tgTypes.TweetThread{}, err
	}
	tweet, err := Map(response, id)
	if err != nil {
		return tgTypes.TweetThread{}, err
	}
	return tweet, nil
}
