package telegram

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

func NewClient(host string, token string) *Client {
	return &Client{
		host:     host,
		basePath: basePath(token),
		client:   http.Client{},
	}
}

type Client struct {
	host     string
	basePath string
	client   http.Client
}

const getUpdates = "getUpdates"
const sendMessage = "sendMessage"
const sendPhotos = "sendMediaGroup"
const sendPhoto = "sendPhoto"
const sendVideo = "sendVideo"

func (c *Client) Updates(offset int, limit int) ([]Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	data, err := c.doRequest(getUpdates, q)
	if err != nil {
		return nil, err
	}
	var response UpdateResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, errors.Wrapf(err, "error while parsing response body")
	}

	return response.Result, nil
}

func (c *Client) SendMessage(chatId int, text string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatId))
	q.Add("text", text)

	_, err := c.doRequest(sendMessage, q)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) SendPhotos(chatId int, text string, photos []string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatId))
	media, err := encodedPhotos(photos, text)
	if err != nil {
		return err
	}
	q.Add("media", media)

	_, err = c.doRequest(sendPhotos, q)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) SendPhoto(chatId int, text string, photo string, button *Button) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatId))
	q.Add("photo", photo)
	q.Add("caption", text)

	// todo?
	//if button != nil {
	//	buttons := make([][]Button, 1)
	//	buttons[0] = []Button{*button}
	//	data, err := json.Marshal(InlineKeyboardMarkup{buttons})
	//	if err == nil {
	//		q.Add("reply_markup", string(data))
	//	}
	//}

	_, err := c.doRequest(sendPhoto, q)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) SendVideo(chatId int, text string, video string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatId))
	q.Add("video", video)
	q.Add("caption", text)

	resp, err := c.doRequest(sendVideo, q)
	if err != nil {
		return err
	}

	_ = resp

	return nil
}

func (c *Client) doRequest(method string, query url.Values) ([]byte, error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, errors.Wrapf(err, "error while creating request: %s", method)
	}

	req.URL.RawQuery = query.Encode()

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "error while making request: %s", method)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "error while reading response %s body", method)
	}
	return body, nil
}

func basePath(token string) string {
	return "bot" + token
}

func encodedPhotos(photos []string, text string) (string, error) {
	InputMediaPhoto := make([]Photo, len(photos))
	for i, p := range photos {
		photo := Photo{
			Type:  "photo",
			Media: p,
		}
		if i == 0 {
			photo.Caption = text
		}
		InputMediaPhoto[i] = photo
	}

	encoded, err := json.Marshal(InputMediaPhoto)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}
