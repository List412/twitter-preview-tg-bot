package commands

import "github.com/pkg/errors"

type Parser interface {
	Parse(text string) (ParsedCmdUrl, error)
}

type Parsers struct {
	parsers map[Cmd]Parser
}

func (p *Parsers) Parse(text string) (Cmd, ParsedCmdUrl, error) {
	for cmd, parser := range p.parsers {
		parsed, err := parser.Parse(text)
		if err != nil {
			continue
		}
		return cmd, parsed, nil
	}
	return "", ParsedCmdUrl{}, errors.New("parser not found")
}

func (p *Parsers) RegisterParser(cmd Cmd, parser Parser) {
	if p.parsers == nil {
		p.parsers = make(map[Cmd]Parser)
	}

	p.parsers[cmd] = parser
}
