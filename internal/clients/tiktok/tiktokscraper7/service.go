package tiktokscraper7

import (
	"context"
	"github.com/list412/twitter-preview-tg-bot/internal/events/telegram/tgTypes"
	"github.com/pkg/errors"
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
		return tgTypes.TweetThread{}, errors.Wrap(err, "get video")
	}
	tweet, err := Map(response)
	if err != nil {
		return tgTypes.TweetThread{}, errors.Wrap(err, "convert response")
	}
	return tweet, nil
}
