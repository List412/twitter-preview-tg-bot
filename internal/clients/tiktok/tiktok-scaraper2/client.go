package tiktok_scaraper2

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

const getVideo = "video/info_v2"

func (c *Client) GetVideo(ctx context.Context, id string) (*VideoParsed, error) {
	q := url.Values{}
	q.Add("video_url", id)

	response, err := c.DoRequest(ctx, getVideo, q)
	if err != nil {
		return nil, err
	}

	log.Printf("GetVideo done %s", id)

	var video VideoParsed
	if err := json.Unmarshal(response, &video); err != nil {
		return nil, err
	}

	return &video, nil
}
