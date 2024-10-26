package event_consumer

import (
	"context"
	"log/slog"
	"sync"
	"time"
	"tweets-tg-bot/internal/commands"
	"tweets-tg-bot/internal/events"
)

func NewConsumer(fetcher events.Fetcher, processor events.Processor, batchSize int) consumer {
	return consumer{
		fetcher:   fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

type consumer struct {
	fetcher   events.Fetcher
	processor events.Processor
	batchSize int
}

func (c *consumer) Start(ctx context.Context) error {

	for {
		select {
		case <-ctx.Done():
			c.processor.Close()
			return nil
		default:
			gotEvents, err := c.fetcher.Fetch(c.batchSize)
			if err != nil {
				slog.Error("consumer", "error", err.Error())
				continue
			}

			if len(gotEvents) == 0 {
				time.Sleep(1 * time.Second)
				continue
			}

			go c.handleEvents(gotEvents)
		}
	}
}

/**
1 потеря событий при ошибке: сохранять кудато? ретрай? забить?
3 счетчик ошибок или ретурн
*/

func (c *consumer) handleEvents(eventsBatch []commands.Event) {
	wg := sync.WaitGroup{}

	wg.Add(len(eventsBatch))

	for _, event := range eventsBatch {
		go func(event commands.Event) {
			defer wg.Done()
			if err := c.processor.Process(event); err != nil {
				slog.Error("can't handle event", "error", err.Error())
			}
		}(event)
	}

	wg.Wait()
}
