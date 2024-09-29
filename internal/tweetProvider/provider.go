package tweetProvider

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"runtime/debug"
	"sync"
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

	combinedErrorMessage := ""
	go func() {
		for err := range errChan {
			log.Print(err.Error())
			combinedErrorMessage += err.Error() + "\n"
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
		result.Source = "twitter"
		return result, nil
	case <-allDone:
		return tgTypes.TweetThread{}, errors.New(combinedErrorMessage)
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

	go p.runGetTweet(ctx, name, api, id, rChan, errChan, doneChan)

	select {
	case <-ctx.Done():
		errChan <- errors.Wrapf(ctx.Err(), "api: %s", name)
	case result := <-rChan:
		resultChan <- result
	case <-doneChan:
		return
	}
}

func (p *Provider) runGetTweet(
	ctx context.Context,
	name string,
	api GetTweetApi,
	id string,
	rChan chan<- tgTypes.TweetThread,
	errChan chan<- error,
	doneChan chan<- struct{},
) {
	defer recoverPanic(name, errChan)
	defer close(doneChan)
	log.Printf("%s start", name)
	defer log.Printf("%s finish", name)

	retry := 0
	for {
		log.Printf("%s attempt #%d id %s", name, retry, id)
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
}

func recoverPanic(name string, errChan chan<- error) {
	if r := recover(); r != nil {
		errChan <- errors.New(fmt.Sprintf("api %s: %+v; stack: %s", name, r, debug.Stack()))
	}
}
