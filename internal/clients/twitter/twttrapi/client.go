package twttrapi

import (
	"context"
	"encoding/json"
	"log"
	"net/url"

	"github.com/pkg/errors"
	"golang.org/x/time/rate"
)

func NewClient(rapidApiClient RapidApiClient, host string) *Client {
	return &Client{rapidApiClient, host, rate.NewLimiter(1, 1)}
}

type Client struct {
	RapidApiClient
	host        string
	rateLimiter *rate.Limiter
}

type RapidApiClient interface {
	DoRequest(ctx context.Context, host string, method string, query url.Values) ([]byte, error)
}

const getTweet = "get-tweet-conversation"
const getTweetSimple = "get-tweet"

func (c Client) GetTweet(ctx context.Context, id string) (*ParsedThread, error) {
	response, err := c.callApi(ctx, getTweet, id)
	if err != nil {
		return nil, err
	}
	var tweet ParsedThread
	if err := json.Unmarshal(response, &tweet); err != nil {
		return nil, err
	}
	return &tweet, nil
}

func (c Client) GetTweetSimple(ctx context.Context, id string) (*ParsedThread, error) {
	response, err := c.callApi(ctx, getTweetSimple, id)
	if err != nil {
		return nil, err
	}
	var tweetS ParsedSimple
	if err := json.Unmarshal(response, &tweetS); err != nil {
		return nil, err
	}

	if len(tweetS.Errors) > 0 {
		return nil, errors.New(tweetS.Errors[0].Message)
	}

	tweet := ParsedThread{}
	entity := Entity{}
	entity.Content.Content.TweetResult.Result = tweetS.Data.TweetResult.Result
	entity.EntryId = "tweet-" + id
	tweet.Data.TimelineResponse.Instructions = append(tweet.Data.TimelineResponse.Instructions, Instructions{
		Entries: []Entity{
			entity,
		},
		Direction: "",
	})

	return &tweet, nil
}

func (c Client) callApi(ctx context.Context, method string, id string) ([]byte, error) {
	q := url.Values{}
	q.Add("tweet_id", id)

	_ = c.rateLimiter.Wait(ctx)
	response, err := c.DoRequest(ctx, c.host, method, q)
	if err != nil {
		return nil, err
	}

	log.Printf("GetTweet done %s", id)

	return response, nil
}
