package twitter

import (
	"github.com/list412/twitter-preview-tg-bot/internal/commands"
	"github.com/pkg/errors"
	"net/url"
	"strings"
)

type CommandParser struct {
}

func (receiver CommandParser) Parse(text string) (commands.ParsedCmdUrl, error) {
	cmdUrl := commands.ParsedCmdUrl{}
	u, err := url.Parse(text)
	if err != nil {
		return cmdUrl, err
	}

	cmdUrl.OriginalUrl = text
	cmdUrl.StrippedUrl = u.Scheme + "://" + u.Host + u.Path

	twitterHosts := []string{"twitter.com", "x.com"}

	isTwitterUrl := false
	for _, h := range twitterHosts {
		if h == u.Host {
			isTwitterUrl = true
			break
		}
	}

	if !isTwitterUrl {
		return cmdUrl, errors.New("not a twitter url")
	}

	path := strings.Split(strings.TrimLeft(u.Path, "/"), "/")
	if len(path) != 3 {
		return cmdUrl, errors.New("url don't have id")
	}

	if path[2] == "" {
		return cmdUrl, errors.New("id in url empty")
	}

	cmdUrl.Key = path[2]
	return cmdUrl, nil
}
