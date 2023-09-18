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
	GetTweet(ctx context.Context, id string) (tgTypes.Tweet, error)
}

func NewProvider() *provider {
	return &provider{}
}

type provider struct {
	tweetApis map[string]GetTweetApi
}

func (p *provider) RegisterApi(name string, getTweetApi GetTweetApi) {
	if p.tweetApis == nil {
		p.tweetApis = map[string]GetTweetApi{}
	}
	p.tweetApis[name] = getTweetApi
}

func (p *provider) GetTweet(id string) (tgTypes.Tweet, error) {

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	resultChan := make(chan tgTypes.Tweet)
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
		return tgTypes.Tweet{}, telegram.ErrApiResponse
	}
}

func (p *provider) runApi(
	ctx context.Context,
	wg *sync.WaitGroup,
	name string,
	api GetTweetApi,
	id string,
	resultChan chan<- tgTypes.Tweet,
	errChan chan<- error,
) {
	defer wg.Done()
	rChan := make(chan tgTypes.Tweet)
	defer close(rChan)
	doneChan := make(chan struct{})

	go func() {
		log.Printf("%s start", name)
		defer log.Printf("%s finish", name)
		result, err := api.GetTweet(ctx, id)
		if err != nil {
			errChan <- errors.Wrapf(err, "api: %s", name)
		} else {
			rChan <- result
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
