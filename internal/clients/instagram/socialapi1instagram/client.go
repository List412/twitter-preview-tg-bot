package socialapi1instagram

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

const getVideo = "v1/post_info"

func (c *Client) GetVideo(ctx context.Context, id string) (*ParsedPost, error) {
	q := url.Values{}
	q.Add("code_or_id_or_url", id)
	q.Add("include_insights", "true")

	response, err := c.DoRequest(ctx, getVideo, q)
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
