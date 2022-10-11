package twitter

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
	"path"
)

func NewClient(host string, token string) *Client {
	return &Client{
		host:   host,
		token:  "Bearer " + token,
		client: http.Client{},
	}
}

type Client struct {
	host   string
	token  string
	client http.Client
}

const getTweet = "2/tweets"

func (c *Client) GetTweet(id string) (*Tweet, error) {
	q := url.Values{}
	q.Add("tweet.fields", "attachments")
	q.Add("user.fields", "name")
	q.Add("expansions", "attachments.media_keys")

	response, err := c.doRequest(getTweet, id, q)
	if err != nil {
		return nil, err
	}

	var tweet Tweet
	if err := json.Unmarshal(response, &tweet); err != nil {
		return nil, err
	}

	return &tweet, nil
}

func (c *Client) doRequest(method string, id string, query url.Values) ([]byte, error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(method, id),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, errors.Wrapf(err, "error while creating request: %s", method)
	}

	req.Header.Add("Authorization", c.token)

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
