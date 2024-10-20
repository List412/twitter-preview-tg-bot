package twitterapi45

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
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

const getTweet = "tweet.php"

func (c *Client) GetTweet(ctx context.Context, id string) (*Response, error) {
	q := url.Values{}
	q.Add("id", id)

	response, err := c.DoRequest(ctx, c.host, getTweet, q)
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
