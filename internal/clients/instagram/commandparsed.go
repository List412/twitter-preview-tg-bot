package instagram

import (
	"github.com/pkg/errors"
	"net/url"
	"strings"
)

type CommandParser struct {
}

func (p CommandParser) Parse(text string) (string, error) {
	u, err := url.Parse(text)

	hosts := []string{"instagram.com", "www.instagram.com"}

	if err != nil {
		return "", err
	}

	isInstagramUrl := false
	for _, h := range hosts {
		if h == u.Host {
			isInstagramUrl = true
			break
		}
	}

	if !isInstagramUrl {
		return "", errors.New("not a instagram url")
	}

	path := strings.Split(strings.Trim(u.Path, "/"), "/")
	if len(path) < 2 || len(path) > 3 {
		return "", errors.New("url don't have id")
	}

	if path[0] == "" {
		return "", errors.New("media type is empty")
	}
	if path[1] == "" {
		return "", errors.New("media code is empty")
	}
	return u.String(), nil
}
