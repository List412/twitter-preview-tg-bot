package twttrapi

import (
	"context"
	"encoding/json"
	"log"
	"net/url"
	"tweets-tg-bot/internal/clients/rapidApi"
)

func NewClient(host string, token string) *Client {
	return &Client{rapidApi.NewClient(host, token)}
}

type Client struct {
	rapidApi.Client
}

const getTweet = "get-tweet-conversation"

func (c *Client) GetTweet(ctx context.Context, id string) (*ParsedThread, error) {
	q := url.Values{}
	q.Add("tweet_id", id)

	response, err := c.DoRequest(ctx, getTweet, q)
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
