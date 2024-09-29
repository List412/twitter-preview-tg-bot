package tiktok89

import (
	"context"
	"tweets-tg-bot/internal/events/telegram/tgTypes"
)

type Service struct {
	client *Client
}

func NewService(client *Client) *Service {
	return &Service{client: client}
}

func (s Service) GetVideo(ctx context.Context, id string) (tgTypes.TweetThread, error) {
	response, err := s.client.GetVideo(ctx, id)
	if err != nil {
		return tgTypes.TweetThread{}, err
	}
	tweet, err := Map(response)
	if err != nil {
		return tgTypes.TweetThread{}, err
	}
	return tweet, nil
}
