package events

import (
	"context"
	"tweets-tg-bot/internal/commands"
)

type Fetcher interface {
	Fetch(limit int) ([]commands.Event, error)
}

type Processor interface {
	Process(event commands.Event) error
	HandleUsers()
	Close()
}

type UsersServiceInterface interface {
	Add(ctx context.Context, userName string) error
	AddShare(ctx context.Context, userName string) error
	IsAdmin(userId int) (bool, error)
	Count(ctx context.Context) (int, error)
	CountShare(ctx context.Context) (int, error)
	Command(cmd commands.Cmd, userName string)
	CommandsStat(ctx context.Context) (map[string]int, error)
}
