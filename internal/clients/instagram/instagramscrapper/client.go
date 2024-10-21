package instagramscrapper

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"log"
	"net/url"
	"strings"
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

const getVideo = "reel_by_shortcode"
const getStory = "story_by_id"
const getPost = "post_by_shortcode"

func (c *Client) GetVideo(ctx context.Context, instUrl string) (*ParsedPost, error) {
	u, err := url.Parse(instUrl)
	if err != nil {
		return nil, errors.Wrap(err, "url parse error")
	}
	path := strings.Split(strings.Trim(u.Path, "/"), "/")
	method := getVideo

	id := ""
	param := "id"

	switch path[0] {
	case "reel":
		method = getVideo
		param = "shortcode"
		id = path[1]
	case "p":
		method = getPost
		param = "shortcode"
		id = path[1]
	case "stories":
		method = getStory
		id = path[2]
	}

	return c.getVideo(ctx, method, id, param)
}

func (c *Client) getVideo(ctx context.Context, method string, id string, param string) (*ParsedPost, error) {
	q := url.Values{}
	q.Add(param, id)

	response, err := c.DoRequest(ctx, c.host, method, q)
	if err != nil {
		return nil, err
	}

	log.Printf("GetPost done %s", id)

	var post ParsedPost
	if err := json.Unmarshal(response, &post); err != nil {
		return nil, err
	}

	//if post.Detail != "" {
	//	return nil, fmt.Errorf("%s", post.Detail)
	//}

	return &post, nil
}
