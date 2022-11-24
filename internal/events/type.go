package events

import (
	"context"
)

type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}

type Processor interface {
	Process(event Event) error
	HandleUsers()
	Close()
}

type UsersServiceInterface interface {
	Add(ctx context.Context, userName string) error
	AddShare(ctx context.Context, userName string) error
	IsAdmin(userId int) (bool, error)
	Count(ctx context.Context) (int, error)
	CountShare(ctx context.Context) (int, error)
}

type Type int

const (
	Unknown Type = iota
	Message
)

type Event struct {
	Type Type
	Text string
	Meta interface{}
}
