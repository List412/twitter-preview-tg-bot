package tiktokscraper7

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

const getVideo = ""

func (c *Client) GetVideo(ctx context.Context, id string) (*VideoParsed, error) {
	q := url.Values{}
	q.Add("url", id)
	q.Add("hd", "1")

	response, err := c.DoRequest(ctx, getVideo, q)
	if err != nil {
		return nil, err
	}

	log.Printf("GetVideo done %s", id)

	var video VideoParsed
	if err := json.Unmarshal(response, &video); err != nil {
		return nil, err
	}

	if video.Code != 0 {
		return nil, errors.New(video.Msg)
	}

	return &video, nil
}
