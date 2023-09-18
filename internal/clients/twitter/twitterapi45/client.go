package twitterapi45

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
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

const getTweet = "tweet.php"

func (c *Client) GetTweet(ctx context.Context, id string) (*Response, error) {
	q := url.Values{}
	q.Add("id", id)

	response, err := c.DoRequest(ctx, getTweet, q)
	if err != nil {
		return nil, err
	}

	log.Printf("GetTweet done %s", id)

	var tweet Response
	if err := json.Unmarshal(response, &tweet); err != nil {
		return nil, errors.Wrap(err, "unmarshal error")
	}

	return &tweet, err
}
