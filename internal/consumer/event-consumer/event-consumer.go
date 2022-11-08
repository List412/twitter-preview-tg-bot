package event_consumer

import (
	"log"
	"sync"
	"time"
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

func (c *consumer) Start() error {

	for {
		gotEvents, err := c.fetcher.Fetch(c.batchSize)
		if err != nil {
			log.Printf("[ERR] consumer: %s", err.Error())
			continue
		}

		if len(gotEvents) == 0 {
			time.Sleep(1 * time.Second)
			continue
		}

		if err := c.handleEvents(gotEvents); err != nil {
			log.Printf("can't handle events %s", err.Error())
			continue
		}
	}
}

/**
1 потеря событий при ошибке: сохранять кудато? ретрай? забить?
3 счетчик ошибок или ретурн
*/

func (c *consumer) handleEvents(eventsBatch []events.Event) error {
	wg := sync.WaitGroup{}

	wg.Add(len(eventsBatch))

	for _, event := range eventsBatch {
		go func(event events.Event) {
			defer wg.Done()
			if err := c.processor.Process(event); err != nil {
				log.Printf("can't handle event: %s", err.Error())
			}
		}(event)
	}

	wg.Wait()

	return nil
}
