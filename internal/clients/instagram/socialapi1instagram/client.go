package socialapi1instagram

import (
	"context"
	"encoding/json"
	"fmt"
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

const getVideo = "v1/post_info"

func (c *Client) GetVideo(ctx context.Context, id string) (*ParsedPost, error) {
	q := url.Values{}
	q.Add("code_or_id_or_url", id)
	q.Add("include_insights", "true")

	response, err := c.DoRequest(ctx, c.host, getVideo, q)
	if err != nil {
		return nil, err
	}

	log.Printf("GetPost done %s", id)

	var post ParsedPost
	if err := json.Unmarshal(response, &post); err != nil {
		return nil, err
	}

	if post.Detail != "" {
		return nil, fmt.Errorf("%s", post.Detail)
	}

	return &post, nil
}
