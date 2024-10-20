package twttrapi

import (
	"context"
	"encoding/json"
	"log"
	"net/url"
)

func NewClient(rapidApiClient RapidApiClient, host string) *Client {
	return &Client{rapidApiClient, host}
}

type Client struct {
	RapidApiClient
	host string
}

type RapidApiClient interface {
	DoRequest(ctx context.Context, host string, method string, query url.Values) ([]byte, error)
}

const getTweet = "get-tweet-conversation"

func (c Client) GetTweet(ctx context.Context, id string) (*ParsedThread, error) {
	q := url.Values{}
	q.Add("tweet_id", id)

	response, err := c.DoRequest(ctx, c.host, getTweet, q)
	if err != nil {
		return nil, err
	}

	log.Printf("GetTweet done %s", id)

	var tweet ParsedThread
	if err := json.Unmarshal(response, &tweet); err != nil {
		return nil, err
	}

	return &tweet, nil
}
