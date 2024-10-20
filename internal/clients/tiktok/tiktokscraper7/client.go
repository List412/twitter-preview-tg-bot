package tiktokscraper7

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

const getVideo = ""

func (c *Client) GetVideo(ctx context.Context, id string) (*VideoParsed, error) {
	q := url.Values{}
	q.Add("url", id)
	q.Add("hd", "1")

	response, err := c.DoRequest(ctx, c.host, getVideo, q)
	if err != nil {
		return nil, err
	}

	log.Printf("GetVideo done %s", id)

	var video VideoParsed
	if err := json.Unmarshal(response, &video); err != nil {
		return nil, errors.Wrap(err, "unmarshal video response")
	}

	if video.Code != 0 {
		return nil, errors.New(video.Msg)
	}

	return &video, nil
}
