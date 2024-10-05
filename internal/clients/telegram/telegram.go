package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/time/rate"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"tweets-tg-bot/internal/events/telegram/tgTypes"
)

func NewClient(host string, token string) *Client {
	return &Client{
		host:     host,
		basePath: basePath(token),
		client:   http.Client{},
		limiter:  rate.NewLimiter(1, 1),
	}
}

type Client struct {
	host     string
	basePath string
	client   http.Client
	limiter  *rate.Limiter
}

const getUpdates = "getUpdates"
const sendMessage = "sendMessage"
const sendMediaGroup = "sendMediaGroup"
const sendPhoto = "sendPhoto"
const sendVideo = "sendVideo"

func (c *Client) Updates(offset int, limit int) ([]Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	data, err := c.get(getUpdates, q)
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
	q.Add("disable_notification", "true")
	q.Add("parse_mode", "HTML")

	_, err := c.get(sendMessage, q)
	if err != nil {
		return err
	}
	return nil
}

// SendPhotos
// send photos
func (c *Client) SendPhotos(chatId int, text string, mediaUrls []MediaForEncoding) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatId))
	q.Add("disable_notification", "true")

	encodedMedia, err := encodedMediaObjects(mediaUrls, text)
	if err != nil {
		return err
	}
	q.Add("media", encodedMedia)

	_, err = c.get(sendMediaGroup, q)
	if err != nil {
		return err
	}

	return nil
}

// SendMedia
// send photo/video
func (c *Client) SendMedia(chatId int, text string, mediaUrls []MediaForEncoding, allMedia []tgTypes.MediaObject) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatId))
	q.Add("disable_notification", "true")

	encodedMedia, err := encodedMediaObjects(mediaUrls, text)
	if err != nil {
		return errors.Wrap(err, "encodedMediaObjects")
	}
	q.Add("media", encodedMedia)

	_, err = c.postMultipart(sendMediaGroup, q, allMedia)
	if err != nil {
		return errors.Wrap(err, "postMultipart")
	}

	return nil
}

func (c *Client) SendPhoto(chatId int, text string, photo string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatId))
	q.Add("photo", photo)
	q.Add("caption", text)
	q.Add("disable_notification", "true")
	q.Add("parse_mode", "HTML")

	resp, err := c.get(sendPhoto, q)
	if err != nil {
		return err
	}

	_ = resp

	return nil
}

func (c *Client) SendVideo(chatId int, text string, video tgTypes.MediaObject) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatId))
	videoPath := video.Url
	q.Add("video", videoPath)
	q.Add("caption", text)
	q.Add("disable_notification", "true")
	q.Add("parse_mode", "HTML")

	_, err := c.get(sendVideo, q)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) do(req *http.Request) ([]byte, error) {
	err := c.limiter.Wait(req.Context())
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "error while making request")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "error while reading response body")
	}
	if resp.StatusCode >= http.StatusBadRequest {
		return nil, c.parseError(body)
	}
	return body, nil
}

func (c *Client) postMultipart(method string, query url.Values, files []tgTypes.MediaObject) ([]byte, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}

	for _, v := range files {
		part, err := w.CreateFormFile(v.Name, v.Name)
		if err != nil {
			return nil, err
		}
		r := bytes.NewReader(v.Data)
		io.Copy(part, r)
	}

	w.Close()
	req, err := http.NewRequest(http.MethodPost, u.String(), &b)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = query.Encode()

	req.Header.Set("Content-Type", w.FormDataContentType())

	return c.do(req)
}

func (c *Client) get(method string, query url.Values) ([]byte, error) {
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

	return c.do(req)
}

func basePath(token string) string {
	return "bot" + token
}

type mediaType = string

const MediaTypePhoto mediaType = "photo"
const MediaTypeVideo mediaType = "video"

func encodedMediaObjects(mediaForEncoding []MediaForEncoding, text string) (string, error) {
	var mediaObjects []MediaObject

	for j, mediaForEncoding := range mediaForEncoding {
		if len(mediaForEncoding.Media) == 0 {
			continue
		}

		mediaUrls := mediaForEncoding.Media
		currentMediaType := mediaForEncoding.MediaType

		for i, v := range mediaUrls {
			mediaPath := v.Url
			if v.NeedUpload || mediaForEncoding.ForceNeedUpload {
				mediaPath = fmt.Sprintf("attach://%s", v.Name)
			}
			media := MediaObject{
				Type:  currentMediaType,
				Media: mediaPath,
			}
			if i == 0 && j == 0 {
				media.Caption = text
				media.ParseMode = "HTML"
			}

			mediaObjects = append(mediaObjects, media)
		}
	}

	encoded, err := json.Marshal(mediaObjects)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func encodedMediaObject(mediaUrl tgTypes.MediaObject, text string, mediaType mediaType) (string, error) {
	mediaPath := mediaUrl.Url
	if mediaUrl.NeedUpload {
		mediaPath = fmt.Sprintf("attach://%s", mediaUrl.Name)
	}
	media := struct {
		Type      string
		Video     string
		Caption   string
		ParseMode string
	}{
		Type:      mediaType,
		Video:     mediaPath,
		Caption:   text,
		ParseMode: "HTML",
	}

	encoded, err := json.Marshal(media)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}
