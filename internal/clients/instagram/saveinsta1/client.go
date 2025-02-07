package saveinsta1

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
	DoPostRequest(ctx context.Context, host string, method string, query url.Values, reqBody []byte) ([]byte, error)
}

const getVideo = "media"

func (c *Client) GetVideo(ctx context.Context, id string) (*ParsedPost, error) {
	q := url.Values{}

	payloadStruct := struct {
		Url string `json:"url"`
	}{
		Url: id,
	}

	body, err := json.Marshal(payloadStruct)
	if err != nil {
		return nil, err
	}

	response, err := c.DoPostRequest(ctx, c.host, getVideo, q, body)
	if err != nil {
		return nil, err
	}

	log.Printf("GetPost done %s", id)

	var post ParsedPost
	if err := json.Unmarshal(response, &post); err != nil {
		return nil, err
	}

	if !post.Success {
		return nil, fmt.Errorf("%s", post.Message)
	}

	return &post, nil
}
