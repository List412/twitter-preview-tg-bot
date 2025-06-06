package clients

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/list412/twitter-preview-tg-bot/internal/commands"
	"github.com/list412/twitter-preview-tg-bot/internal/events/telegram/tgTypes"
)

type ContentProvider interface {
	GetContent(ctx context.Context, urlCmd commands.ParsedCmdUrl) (tgTypes.TweetThread, error)
}

type Manager struct {
	apiServices  map[commands.Cmd]ContentProvider
	enabledFlags map[commands.Cmd]bool
}

func (m *Manager) RegisterService(cmd commands.Cmd, api ContentProvider) *Manager {
	if m.apiServices == nil {
		m.apiServices = make(map[commands.Cmd]ContentProvider)
	}
	if m.enabledFlags == nil {
		m.enabledFlags = make(map[commands.Cmd]bool)
	}

	m.apiServices[cmd] = api
	m.enabledFlags[cmd] = true

	return m
}

func (m *Manager) GetContent(ctx context.Context, cmd commands.Cmd, urlCmd commands.ParsedCmdUrl) (tgTypes.TweetThread, error) {
	enabled, ok := m.enabledFlags[cmd]
	if !ok || !enabled {
		return tgTypes.TweetThread{}, NewErrServiceDisabled(cmd)
	}

	apiService, ok := m.apiServices[cmd]
	if !ok {
		return tgTypes.TweetThread{}, ErrServiceNotFound
	}

	result, err := apiService.GetContent(ctx, urlCmd)
	if err != nil {
		return tgTypes.TweetThread{}, errors.Wrap(err, fmt.Sprintf("Getting %s content from %s failed", urlCmd, cmd))
	}

	return result, nil
}

func NewManager() *Manager {
	return &Manager{
		apiServices:  make(map[commands.Cmd]ContentProvider),
		enabledFlags: make(map[commands.Cmd]bool),
	}
}
