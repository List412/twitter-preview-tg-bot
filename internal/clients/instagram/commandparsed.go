package instagram

import (
	"net/url"
	"strings"

	"github.com/pkg/errors"

	"github.com/list412/tweets-tg-bot/internal/commands"
)

type CommandParser struct {
}

func (p CommandParser) Parse(text string) (commands.ParsedCmdUrl, error) {
	u, err := url.Parse(text)
	if err != nil {
		return commands.ParsedCmdUrl{}, err
	}

	hosts := []string{"instagram.com", "www.instagram.com"}
	isInstagramUrl := false
	for _, h := range hosts {
		if h == u.Host {
			isInstagramUrl = true
			break
		}
	}

	parsedUrl := commands.ParsedCmdUrl{}
	parsedUrl.OriginalUrl = text
	parsedUrl.StrippedUrl = u.Scheme + "://" + u.Host + u.Path

	if !isInstagramUrl {
		return commands.ParsedCmdUrl{}, errors.New("not a instagram url")
	}

	path := strings.Split(strings.Trim(u.Path, "/"), "/")
	if len(path) == 0 {
		return commands.ParsedCmdUrl{}, errors.New("url don't have id")
	}
	if path[0] == "" {
		return commands.ParsedCmdUrl{}, errors.New("media type is empty")
	}
	if len(path) <= 1 {
		return commands.ParsedCmdUrl{}, errors.New("media id is empty")
	}
	switch path[0] {
	case "share":
		{
			if len(path) < 3 {
				return commands.ParsedCmdUrl{}, errors.New("media code not found")
			}
			switch path[1] {
			case "reel":
				{
					parsedUrl.Key = path[2]
				}
			}
		}
	case "reel":
		fallthrough
	case "p":
		if len(path) != 2 || path[1] == "" {
			return commands.ParsedCmdUrl{}, errors.New("media code is empty")
		}
		parsedUrl.Key = path[1]
	case "stories":
		if len(path) != 3 || path[2] == "" {
			return commands.ParsedCmdUrl{}, errors.New("media code is empty")
		}
		parsedUrl.Key = path[2]
	}

	return parsedUrl, nil
}
