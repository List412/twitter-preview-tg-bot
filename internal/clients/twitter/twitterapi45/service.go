package twitterapi45

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

func (s Service) GetTweet(ctx context.Context, id string) (tgTypes.Tweet, error) {
	response, err := s.client.GetTweet(ctx, id)
	if err != nil {
		return tgTypes.Tweet{}, err
	}
	tweet, err := Map(response)
	if err != nil {
		return tgTypes.Tweet{}, err
	}
	return tweet, nil
}
