package tiktok_scaraper2

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

const getVideo = "video/info_v2"

func (c *Client) GetVideo(ctx context.Context, id string) (*VideoParsed, error) {
	q := url.Values{}
	q.Add("video_url", id)

	response, err := c.DoRequest(ctx, c.host, getVideo, q)
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
