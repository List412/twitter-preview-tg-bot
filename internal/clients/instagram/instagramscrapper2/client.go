package instagramscrapper2

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

const getVideo = "media_info"

func (c *Client) GetVideo(ctx context.Context, id string) (*ParsedPost, error) {
	q := url.Values{}
	q.Add("short_code", id)

	response, err := c.DoRequest(ctx, c.host, getVideo, q)
	if err != nil {
		return nil, err
	}

	log.Printf("GetPost done %s", id)

	var post ParsedPost
	if err := json.Unmarshal(response, &post); err != nil {
		return nil, err
	}

	return &post, nil
}
