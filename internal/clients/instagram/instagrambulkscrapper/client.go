package instagrambulkscrapper

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"log"
	"net/url"
	"path"
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
	DoPostRequest(ctx context.Context, host string, method string, query url.Values, body []byte) ([]byte, error)
}

const getStory = "download_story_from_url"
const getPost = "media_download_by_shortcode"

func (c *Client) GetVideo(ctx context.Context, instUrl string) (*ParsedPost, error) {
	u, err := url.Parse(instUrl)
	if err != nil {
		return nil, errors.Wrap(err, "url parse error")
	}
	upath := strings.Split(strings.Trim(u.Path, "/"), "/")
	method := getPost

	var response []byte
	if upath[0] == "stories" {
		method = getStory

		var reqeust struct {
			Url string `json:"url"`
		}
		reqeust.Url = instUrl

		bodyRequest, err := json.Marshal(reqeust)
		if err != nil {
			return nil, errors.Wrap(err, "json marshal error")
		}

		response, err = c.DoPostRequest(ctx, c.host, method, url.Values{}, bodyRequest)
		if err != nil {
			return nil, err
		}
	} else {
		q := url.Values{}
		response, err = c.DoRequest(ctx, c.host, path.Join(method, upath[1]), q)
		if err != nil {
			return nil, err
		}
	}

	log.Printf("GetPost done %s", instUrl)

	var post ParsedPost
	if err := json.Unmarshal(response, &post); err != nil {
		return nil, err
	}

	//if post.Detail != "" {
	//	return nil, fmt.Errorf("%s", post.Detail)
	//}

	return &post, nil
}
