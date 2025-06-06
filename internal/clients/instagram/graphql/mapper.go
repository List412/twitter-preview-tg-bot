package graphql

import (
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/list412/tweets-tg-bot/internal/downloader"
	"github.com/list412/tweets-tg-bot/internal/events/telegram/tgTypes"
)

func Map(post *ParsedPost) (tgTypes.TweetThread, error) {
	tweet := tgTypes.TweetThread{}

	data := post.Data.XdtShortcodeMedia
	owner := data.Owner

	tweet.UserId = owner.Username
	tweet.UserName = owner.FullName

	tweet.Views = strconv.Itoa(data.VideoViewCount)
	tweet.Likes = data.EdgeMediaPreviewLike.Count
	tweet.Replies = data.EdgeMediaPreviewComment.Count

	text := ""
	if len(data.EdgeMediaToCaption.Edges) > 0 {
		text = data.EdgeMediaToCaption.Edges[0].Node.Text
		timeInt, err := strconv.Atoi(data.EdgeMediaToCaption.Edges[0].Node.CreatedAt)
		if err == nil {
			tweet.Time = time.Unix(int64(timeInt), 0)
		}
	}

	media := tgTypes.Media{}
	if data.IsVideo {
		media.Videos = append(media.Videos, tgTypes.MediaObject{
			Name:       fmt.Sprintf("%s_%s.mp4", data.ID, data.Shortcode),
			Url:        data.VideoURL,
			NeedUpload: true,
		})
	} else {
		if len(data.EdgeSidecarToChildren.Edges) > 0 {
			for i, child := range data.EdgeSidecarToChildren.Edges {
				if child.Node.IsVideo {
					media.Videos = append(media.Videos, tgTypes.MediaObject{
						Name:       fmt.Sprintf("%s_%s_%d.mp4", child.Node.Shortcode, child.Node.Id, i),
						Url:        child.Node.VideoUrl,
						NeedUpload: true,
					})
				} else {
					media.Photos = append(media.Photos, tgTypes.MediaObject{
						Name:       fmt.Sprintf("%s_%s_%d.jpg", child.Node.Shortcode, child.Node.Id, i),
						Url:        child.Node.DisplayUrl,
						NeedUpload: false,
					})
				}
			}
		} else {
			media.Photos = append(media.Photos, tgTypes.MediaObject{
				Name:       fmt.Sprintf("%s_%s.jpg", data.ID, data.Shortcode),
				Url:        data.DisplayURL,
				NeedUpload: false,
			})
		}
	}

	for _, photoChunk := range chunkMedia(media.Photos, 10) {
		content := tgTypes.TweetContent{}
		mediaChunk := tgTypes.Media{}

		mediaChunk.Photos = photoChunk

		content.Media = mediaChunk
		tweet.Tweets = append(tweet.Tweets, content)
	}

	for _, videoChunk := range chunkVideos(media.Videos, 10) {
		content := tgTypes.TweetContent{}
		mediaChunk := tgTypes.Media{}

		mediaChunk.Videos = videoChunk

		content.Media = mediaChunk
		tweet.Tweets = append(tweet.Tweets, content)
	}

	if tweet.Tweets == nil {
		tweet.Tweets = append(tweet.Tweets, tgTypes.TweetContent{
			Text: text,
		})
	} else {
		tweet.Tweets[0].Text = text
	}

	return tweet, nil
}

func chunkMedia(media []tgTypes.MediaObject, length int) [][]tgTypes.MediaObject {
	var result [][]tgTypes.MediaObject

	mediaIndex := 0
	for mediaIndex < len(media) {
		chunkLength := min(length, len(media[mediaIndex:]))
		result = append(result, media[mediaIndex:mediaIndex+chunkLength])
		mediaIndex += chunkLength
	}
	return result
}

func chunkVideos(media []tgTypes.MediaObject, length int) [][]tgTypes.MediaObject {
	var result [][]tgTypes.MediaObject
	maxChunkVideosSize := 50 * 1024 * 1024

	mediaIndex := 0

	currentChunkSize := 0

	var tmpSlice []tgTypes.MediaObject

	if len(media) < 1 {
		return result
	}

	if len(media) == 1 {
		result = append(result, media)
		return result
	}

	for _, m := range media {
		s, err := downloader.FileSize(m.Url)
		if err != nil {
			continue
		}
		size := int(s)
		if size > maxChunkVideosSize {
			continue
		}
		if currentChunkSize+size > maxChunkVideosSize || len(tmpSlice)+1 >= length {
			mediaIndex++
			result = append(result, tmpSlice)
			currentChunkSize = 0
		} else {
			currentChunkSize += size
		}

		tmpSlice = append(tmpSlice, m)
	}

	result = append(result, tmpSlice[len(result[0]):])

	return result
}

func checkFileSize(src string) bool {
	size, err := downloader.FileSize(src)
	if err != nil {
		return false
	}
	return size <= 50*1024*1024
}

func chooseVideoVariant(variants []DisplayResource) (*tgTypes.MediaObject, error) {
	for i := len(variants) - 1; i >= 0; i-- {
		size, err := downloader.FileSize(variants[i].Src)
		if err != nil {
			continue
		}

		if size <= 50*1024*1024 {
			return &tgTypes.MediaObject{
				Name:       fmt.Sprintf("video_%d", i),
				Url:        variants[i].Src,
				NeedUpload: true,
			}, nil
		}
	}

	return nil, errors.New("could not find video variant")
}
