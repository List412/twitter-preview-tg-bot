package tiktok

import (
	"context"
	"github.com/pkg/errors"
	"log"
	"tweets-tg-bot/internal/events/telegram/tgTypes"
)

type Api interface {
	GetVideo(ctx context.Context, id string) (tgTypes.TweetThread, error)
}

type Service struct {
	apis []Api
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) RegisterApi(api ...Api) {
	s.apis = append(s.apis, api...)
}

func (s *Service) GetVideo(ctx context.Context, id string) (tgTypes.TweetThread, error) {
	for _, api := range s.apis {
		result, err := api.GetVideo(ctx, id)
		if err != nil {
			log.Println("GetVideo", err)
			continue
		}
		result.Source = "tiktok"
		return result, nil
	}
	return tgTypes.TweetThread{}, errors.New("failed to retrieve video")
}
