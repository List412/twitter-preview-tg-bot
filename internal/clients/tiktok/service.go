package tiktok

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"runtime/debug"
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
		result, err := s.getVideoOrError(ctx, api, id)
		if err != nil {
			log.Println("GetVideo", err)
			continue
		}
		result.Source = "tiktok"
		return result, nil
	}
	return tgTypes.TweetThread{}, errors.New("failed to retrieve video")
}

func (s *Service) getVideoOrError(ctx context.Context, api Api, id string) (tweet tgTypes.TweetThread, err error) {
	defer func() {
		r := recover()
		if r != nil {
			err = errors.Wrap(r.(error), fmt.Sprintf("%s\n", debug.Stack()))
		}
	}()

	return api.GetVideo(ctx, id)
}
