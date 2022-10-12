package twimg_cdn

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
	"path"
)

func NewClient(host string) *Client {
	return &Client{
		host:   host,
		client: http.Client{},
	}
}

type Client struct {
	host   string
	token  string
	client http.Client
}

const getTweet = "tweet"

// followers count
// https://cdn.syndication.twimg.com/widgets/followbutton/info.json?screen_names=konstruktors

func (c *Client) GetTweet(id string) (*Tweet, error) {
	q := url.Values{}
	q.Add("id", id)

	response, err := c.doRequest(getTweet, q)
	if err != nil {
		return nil, err
	}

	var tweet Tweet
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

	req.URL.RawQuery = query.Encode()

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
