package telegram

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/time/rate"
)

func NewClient(host, token, parseMode string) *Client {
	return &Client{
		host:       host,
		basePath:   basePath(token),
		client:     http.Client{},
		limiter:    rate.NewLimiter(10, 1),
		maxRetries: 4,
		parseMode:  parseMode,
	}
}

type Client struct {
	host       string
	basePath   string
	client     http.Client
	limiter    *rate.Limiter
	maxRetries int
	parseMode  string
}

const getUpdates = "getUpdates"
const sendMessage = "sendMessage"
const sendMediaGroup = "sendMediaGroup"
const sendPhoto = "sendPhoto"
const sendVideo = "sendVideo"

const getChat = "getChat"
const getChatAdmins = "getChatAdministrators"
const leaveChat = "leaveChat"

const getFile = "getFile"

func (c *Client) Download(filepath string) ([]byte, error) {
	// https://api.telegram.org/file/bot<token>/<file_path>
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join("file", c.basePath, filepath),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, errors.Wrapf(err, "error while downloading file")
	}

	return c.do(req)
}

func (c *Client) GetFile(filepath string) (*GetFileResult, error) {
	q := url.Values{}
	q.Add("file_id", filepath)
	data, err := c.get(getFile, q)
	if err != nil {
		return nil, err
	}
	var response GetFileResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, errors.Wrapf(err, "error while parsing response body")
	}

	if !response.Ok {
		return nil, errors.Errorf("error while parsing response body")
	}

	return &response.Result, nil
}

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

func (c *Client) SendMessage(chatId, topicId int, text string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatId))
	q.Add("text", text)
	q.Add("disable_notification", "true")
	q.Add("parse_mode", c.parseMode)

	if topicId > 0 {
		q.Add("message_thread_id", strconv.Itoa(topicId))
	}

	_, err := c.get(sendMessage, q)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) ReplyToMessage(chatId, topicId, replyToMessageId int, text string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatId))
	q.Add("text", text)
	q.Add("disable_notification", "true")
	q.Add("parse_mode", c.parseMode)
	q.Add("reply_to_message_id", strconv.Itoa(replyToMessageId))

	if topicId > 0 {
		q.Add("message_thread_id", strconv.Itoa(topicId))
	}

	_, err := c.get(sendMessage, q)
	if err != nil {
		return err
	}
	return nil
}

// SendPhotos
// send photos
func (c *Client) SendPhotos(chatId, topicId int, text string, mediaUrls []MediaForEncoding) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatId))
	q.Add("disable_notification", "true")

	if topicId > 0 {
		q.Add("message_thread_id", strconv.Itoa(topicId))
	}

	encodedMedia, err := encodedMediaObjects(mediaUrls, text)
	if err != nil {
		return err
	}
	q.Add("media", encodedMedia)

	retry := 0
	for retry <= c.maxRetries {
		_, err = c.get(sendMediaGroup, q)
		if errors.Is(err, ErrToManyRequests) && retry < c.maxRetries {
			retry++
			continue
		}
		if err != nil {
			return errors.Wrap(err, "get")
		}
		return nil
	}

	return nil
}

// SendMedia
// send photo/video
func (c *Client) SendMedia(chatId, topicId int, text string, mediaUrls []MediaForEncoding, allMedia []MediaObject) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatId))
	q.Add("disable_notification", "true")

	if topicId > 0 {
		q.Add("message_thread_id", strconv.Itoa(topicId))
	}

	encodedMedia, err := encodedMediaObjects(mediaUrls, text)
	if err != nil {
		return errors.Wrap(err, "encodedMediaObjects")
	}
	q.Add("media", encodedMedia)

	retry := 0
	for retry <= c.maxRetries {
		_, err = c.postMultipart(sendMediaGroup, q, allMedia)
		if errors.Is(err, ErrToManyRequests) && retry < c.maxRetries {
			retry++
			continue
		}
		if err != nil {
			return errors.Wrap(err, "postMultipart")
		}
		return nil
	}

	return nil
}

func (c *Client) SendPhoto(chatId, topicId int, text string, photo string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatId))
	q.Add("photo", photo)
	q.Add("caption", text)
	q.Add("disable_notification", "true")
	q.Add("parse_mode", c.parseMode)

	if topicId > 0 {
		q.Add("message_thread_id", strconv.Itoa(topicId))
	}

	resp, err := c.get(sendPhoto, q)
	if err != nil {
		return err
	}

	_ = resp

	return nil
}

func (c *Client) SendVideo(chatId, topicId int, text string, video MediaObject) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatId))
	videoPath := video.Url
	q.Add("video", videoPath)
	q.Add("caption", text)
	q.Add("disable_notification", "true")
	q.Add("parse_mode", c.parseMode)

	if topicId > 0 {
		q.Add("message_thread_id", strconv.Itoa(topicId))
	}

	_, err := c.get(sendVideo, q)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) LeaveChat(ctx context.Context, chatId int) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatId))

	resp, err := c.get(leaveChat, q)
	if err != nil {
		return err
	}

	_ = resp

	return nil
}

func (c *Client) GetChat(ctx context.Context, id int) (ChatFullInfo, error) {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(id))

	resp, err := c.get(getChat, q)
	if err != nil {
		return ChatFullInfo{}, err
	}

	chatInfo := ChatFullInfo{}
	if err := json.Unmarshal(resp, &chatInfo); err != nil {
		return ChatFullInfo{}, err
	}

	if !chatInfo.Ok {
		return chatInfo, errors.New("invalid response")
	}

	return chatInfo, nil
}

func (c *Client) GetChatAdmins(ctx context.Context, id int) (ChatAdmins, error) {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(id))

	resp, err := c.get(getChatAdmins, q)
	if err != nil {
		return ChatAdmins{}, err
	}

	chatInfo := ChatAdmins{}
	if err := json.Unmarshal(resp, &chatInfo); err != nil {
		return ChatAdmins{}, err
	}

	if !chatInfo.Ok {
		return chatInfo, errors.New("invalid response")
	}

	return chatInfo, nil
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
		if resp.StatusCode == http.StatusTooManyRequests {
			tgErr, err := c.parseErrorStruct(body)
			if err != nil {
				return nil, err
			}
			slog.Info(fmt.Sprintf("To many requests, retry in %d sec...", tgErr.Parameters.RetryAfter))
			time.Sleep(time.Second * time.Duration(tgErr.Parameters.RetryAfter))
			return nil, ErrToManyRequests
		}

		return nil, errors.Wrapf(c.parseError(body), "response status %s", resp.Status)
	}
	return body, nil
}

func (c *Client) postMultipart(method string, query url.Values, files []MediaObject) ([]byte, error) {
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

	if query != nil {
		req.URL.RawQuery = query.Encode()
	}

	return c.do(req)
}

func basePath(token string) string {
	return "bot" + token
}

type mediaType = string

const MediaTypePhoto mediaType = "photo"
const MediaTypeVideo mediaType = "video"

func encodedMediaObjects(mediaForEncoding []MediaForEncoding, text string) (string, error) {
	var mediaObjects []EncodedMediaObject

	for j, mediaForEncoding := range mediaForEncoding {
		if len(mediaForEncoding.Media) == 0 {
			continue
		}

		mediaUrls := mediaForEncoding.Media
		currentMediaType := mediaForEncoding.MediaType

		for i, v := range mediaUrls {
			mediaPath := v.Url
			if v.NeedUpload {
				mediaPath = fmt.Sprintf("attach://%s", v.Name)
			}
			media := EncodedMediaObject{
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

func encodedMediaObject(mediaUrl MediaObject, text string, mediaType mediaType) (string, error) {
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
