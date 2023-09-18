package rapidApi

import (
	"context"
	"github.com/pkg/errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
)

func NewClient(host string, token string) Client {
	return Client{
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

func (c *Client) DoRequest(ctx context.Context, method string, query url.Values) ([]byte, error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(method),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
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
