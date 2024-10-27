package instagram

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"runtime/debug"
	"time"
	"tweets-tg-bot/internal/commands"
	"tweets-tg-bot/internal/events/telegram/tgTypes"
)

type Api interface {
	GetPost(ctx context.Context, id string) (tgTypes.TweetThread, error)
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

func (s *Service) GetContent(ctx context.Context, cmdUrl commands.ParsedCmdUrl) (tgTypes.TweetThread, error) {
	retries := 2
	var lastErr error
	for retries > 0 {
		retries--
		for _, api := range s.apis {
			result, err := s.getPostOrError(ctx, api, cmdUrl)
			if err != nil {
				log.Println("GetPost", err)
				lastErr = err
				continue
			}
			result.Source = "instagram"
			return result, nil
		}
		time.Sleep(15 * time.Second)
	}

	return tgTypes.TweetThread{}, errors.Wrap(lastErr, "failed to retrieve post")
}

func (s *Service) getPostOrError(ctx context.Context, api Api, cmdUrl commands.ParsedCmdUrl) (tweet tgTypes.TweetThread, err error) {
	defer func() {
		r := recover()
		if r != nil {
			err = errors.Wrap(r.(error), fmt.Sprintf("%s\n", debug.Stack()))
		}
	}()

	return api.GetPost(ctx, cmdUrl.StrippedUrl)
}
