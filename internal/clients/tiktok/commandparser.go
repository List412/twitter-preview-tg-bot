package tiktok

import (
	"github.com/list412/tweets-tg-bot/internal/commands"
	"github.com/pkg/errors"
	"net/url"
	"strings"
)

type CommandParser struct {
}

func (p CommandParser) Parse(text string) (commands.ParsedCmdUrl, error) {
	cmdUrl := commands.ParsedCmdUrl{}
	u, err := url.Parse(text)
	if err != nil {
		return cmdUrl, err
	}

	cmdUrl.OriginalUrl = text
	cmdUrl.StrippedUrl = u.Scheme + "://" + u.Host + u.Path

	hosts := []string{"tiktok.com", "vt.tiktok.com"}

	isTiktokUrl := false
	for _, h := range hosts {
		if h == u.Host {
			isTiktokUrl = true
			break
		}
	}

	if !isTiktokUrl {
		return cmdUrl, errors.New("not a tiktok url")
	}

	path := strings.Split(strings.Trim(u.Path, "/"), "/")
	if len(path) != 1 {
		return cmdUrl, errors.New("url don't have id")
	}

	if path[0] == "" {
		return cmdUrl, errors.New("id in url empty")
	}
	cmdUrl.Key = path[0]
	return cmdUrl, nil
}
