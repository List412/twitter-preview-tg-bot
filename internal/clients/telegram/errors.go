package telegram

import (
	"encoding/json"
	"github.com/pkg/errors"
)

var ErrNoEnoughRightToSendPhoto = errors.New("not enough rights to send photos to the chat")
var ErrNoEnoughRightToSendVideo = errors.New("not enough rights to send video to the chat")
var ErrToManyRequests = errors.New("too many requests")

type tgError struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
	Parameters  struct {
		RetryAfter int `json:"retry_after"`
	} `json:"parameters"`
}

func (c *Client) parseErrorStruct(body []byte) (tgError, error) {
	tgErr := tgError{}
	err := json.Unmarshal(body, &tgErr)
	if err != nil {
		return tgErr, errors.Wrap(err, "failed to parse Telegram Error")
	}
	return tgErr, nil
}

func (c *Client) parseError(body []byte) error {
	tgErr := tgError{}
	err := json.Unmarshal(body, &tgErr)
	if err != nil {
		return errors.New(string(body))
	}

	switch tgErr.Description {
	case "Bad Request: not enough rights to send photos to the chat":
		return ErrNoEnoughRightToSendPhoto
	case "Bad Request: not enough rights to send videos to the chat":
		return ErrNoEnoughRightToSendVideo
	default:
		return errors.New(tgErr.Description)
	}
}
