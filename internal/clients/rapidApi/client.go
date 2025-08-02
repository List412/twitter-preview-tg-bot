package rapidApi

import (
	"bytes"
	"context"
	"io"
	"log"
	"log/slog"
	"net/http"
	"net/url"
	"path"

	"github.com/pkg/errors"
)

func NewClient(token string) Client {
	return Client{
		token:  token,
		client: http.Client{},
	}
}

type Client struct {
	token  string
	client http.Client
}

func (c Client) DoRequest(ctx context.Context, host string, method string, query url.Values) ([]byte, error) {
	u := url.URL{
		Scheme: "https",
		Host:   host,
		Path:   path.Join(method),
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, errors.Wrapf(err, "error while creating request: %s", method)
	}

	req.Header.Add("X-RapidAPI-Key", c.token)
	req.Header.Add("X-RapidAPI-Host", host)

	req.URL.RawQuery = query.Encode()
	log.Printf("%s %s %s\n", req.RemoteAddr, req.Method, req.URL)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "error while making request: %s", method)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.Info("Request failed", "status", resp.StatusCode, "path", u.String(), "headers:", resp.Header)
		return nil, errors.Errorf("error while making request: %s, status: %s", method, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "error while reading response %s body", method)
	}
	return body, nil
}

func (c Client) DoPostRequest(ctx context.Context, host string, method string, query url.Values, reqBody []byte) ([]byte, error) {
	u := url.URL{
		Scheme: "https",
		Host:   host,
		Path:   path.Join(method),
	}

	bodyReader := bytes.NewReader(reqBody)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u.String(), bodyReader)
	if err != nil {
		return nil, errors.Wrapf(err, "error while creating request: %s", method)
	}

	req.Header.Add("X-RapidAPI-Key", c.token)
	req.Header.Add("X-RapidAPI-Host", host)

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
