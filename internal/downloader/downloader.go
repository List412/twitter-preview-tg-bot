package downloader

import (
	"io"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/list412/tweets-tg-bot/internal/events/telegram/tgTypes"
)

type Downloader struct {
}

func (d Downloader) FileSize(url string) (uint64, error) {
	return FileSize(url)
}

func (d Downloader) Download(urls []tgTypes.MediaObject) ([]tgTypes.MediaObject, error) {
	return Download(urls)
}

func Download(urls []tgTypes.MediaObject) ([]tgTypes.MediaObject, error) {
	errGr := errgroup.Group{}

	for i, v := range urls {
		if !v.NeedUpload {
			continue
		}
		index := i
		url := v.Url
		errGr.Go(func() error {
			body, err := DownloadFile(url)
			if err != nil {
				return err
			}
			urls[index].Data = body
			return nil
		})
	}
	err := errGr.Wait()
	if err != nil {
		return nil, err
	}
	return urls, nil
}

func DownloadFile(url string) ([]byte, error) {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func FileSize(url string) (uint64, error) {
	resp, err := http.Head(url)
	if err != nil {
		return 0, err
	}

	contentLength := resp.Header.Get("Content-Length")
	result, err := strconv.ParseUint(contentLength, 10, 64)
	if err != nil {
		return 0, errors.New("invalid content length")
	}

	return result, nil
}
