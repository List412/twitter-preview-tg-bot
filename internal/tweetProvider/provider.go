package tweetProvider

import (
	"context"
	"github.com/pkg/errors"
	"log"
	"sync"
	"tweets-tg-bot/internal/events/telegram"
	"tweets-tg-bot/internal/events/telegram/tgTypes"
)

type GetTweetApi interface {
	GetTweet(ctx context.Context, id string) (tgTypes.TweetThread, error)
}

func NewProvider() *Provider {
	return &Provider{}
}

type Provider struct {
	tweetApis map[string]GetTweetApi
}

func (p *Provider) RegisterApi(name string, getTweetApi GetTweetApi) {
	if p.tweetApis == nil {
		p.tweetApis = map[string]GetTweetApi{}
	}
	p.tweetApis[name] = getTweetApi
}

func (p *Provider) GetTweet(id string) (tgTypes.TweetThread, error) {

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	resultChan := make(chan tgTypes.TweetThread)
	defer close(resultChan)
	errChan := make(chan error, len(p.tweetApis))

	wg := &sync.WaitGroup{}
	for name, api := range p.tweetApis {
		wg.Add(1)
		go p.runApi(ctx, wg, name, api, id, resultChan, errChan)
	}

	go func() {
		for err := range errChan {
			log.Print(err.Error())
		}
	}()

	allDone := make(chan struct{})

	go func() {
		defer close(allDone)
		defer close(errChan)
		wg.Wait()
		allDone <- struct{}{}
	}()

	select {
	case result := <-resultChan:
		return result, nil
	case <-allDone:
		return tgTypes.TweetThread{}, telegram.ErrApiResponse
	}
}

func (p *Provider) runApi(
	ctx context.Context,
	wg *sync.WaitGroup,
	name string,
	api GetTweetApi,
	id string,
	resultChan chan<- tgTypes.TweetThread,
	errChan chan<- error,
) {
	defer wg.Done()
	rChan := make(chan tgTypes.TweetThread)
	defer close(rChan)
	doneChan := make(chan struct{})

	go func() {
		log.Printf("%s start", name)
		defer log.Printf("%s finish", name)

		retry := 0
		for {
			result, err := api.GetTweet(ctx, id)
			if err != nil {
				retry++
				if retry <= 3 {
					continue
				}
				errChan <- errors.Wrapf(err, "api: %s", name)
			} else {
				rChan <- result
			}
			break
		}

		close(doneChan)
	}()

	select {
	case <-ctx.Done():
		errChan <- errors.Wrapf(ctx.Err(), "api: %s", name)
	case result := <-rChan:
		resultChan <- result
	case <-doneChan:
		return
	}
}
