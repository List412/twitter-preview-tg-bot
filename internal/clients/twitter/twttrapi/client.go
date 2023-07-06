package twttrapi

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
)

func NewClient(host string, token string) *Client {
	return &Client{
		host:   host,
		token:  token,
		client: http.Client{},
	}
}

type Client struct {
	host   string
	token  string
	client http.Client
}

const getTweet = "get-tweet"

func (c *Client) GetTweet(id string) (*ParsedTweet, error) {
	q := url.Values{}
	q.Add("tweet_id", id)

	response, err := c.doRequest(getTweet, q)
	if err != nil {
		return nil, err
	}

	log.Printf("GetTweet done %s", id)

	var tweet ParsedTweet
	if err := json.Unmarshal(response, &tweet); err != nil {
		return nil, err
	}

	return &tweet, nil
}

func (c *Client) doRequest(method string, query url.Values) ([]byte, error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, errors.Wrapf(err, "error while creating request: %s", method)
	}

	req.Header.Add("X-RapidAPI-Key", c.token)
	req.Header.Add("X-RapidAPI-Host", c.host)

	req.URL.RawQuery = query.Encode()

	log.Printf("%s %s %s\n", req.RemoteAddr, req.Method, req.URL)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "error while making request: %s", method)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "error while reading response %s body", method)
	}
	return body, nil
}
