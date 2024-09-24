package twitter

import (
	"github.com/pkg/errors"
	"net/url"
	"strings"
)

type CommandParser struct {
}

func (receiver CommandParser) Parse(text string) (string, error) {
	u, err := url.Parse(text)

	twitterHosts := []string{"twitter.com", "x.com"}

	if err != nil {
		return "", err
	}

	isTwitterUrl := false
	for _, h := range twitterHosts {
		if h == u.Host {
			isTwitterUrl = true
			break
		}
	}

	if !isTwitterUrl {
		return "", errors.New("not a twitter url")
	}

	path := strings.Split(strings.TrimLeft(u.Path, "/"), "/")
	if len(path) != 3 {
		return "", errors.New("url don't have id")
	}

	if path[2] == "" {
		return "", errors.New("id in url empty")
	}
	return path[2], nil
}
