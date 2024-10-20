package tiktok

import (
	"github.com/pkg/errors"
	"net/url"
	"strings"
)

type CommandParser struct {
}

func (p CommandParser) Parse(text string) (string, error) {
	u, err := url.Parse(text)

	hosts := []string{"tiktok.com", "vt.tiktok.com"}

	if err != nil {
		return "", err
	}

	isTiktokUrl := false
	for _, h := range hosts {
		if h == u.Host {
			isTiktokUrl = true
			break
		}
	}

	if !isTiktokUrl {
		return "", errors.New("not a tiktok url")
	}

	path := strings.Split(strings.Trim(u.Path, "/"), "/")
	if len(path) != 1 {
		return "", errors.New("url don't have id")
	}

	if path[0] == "" {
		return "", errors.New("id in url empty")
	}
	return u.String(), nil
}
