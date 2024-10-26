package downloader

import "tweets-tg-bot/internal/events/telegram/tgTypes"

type Mock struct {
}

func (d Mock) FileSize(url string) (uint64, error) {
	return uint64(100), nil
}

func (d Mock) Download(urls []tgTypes.MediaObject) ([]tgTypes.MediaObject, error) {
	return Download(urls)
}
