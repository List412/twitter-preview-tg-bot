package saveinsta1

import (
	"fmt"
	"github.com/pkg/errors"
	"time"
	"tweets-tg-bot/internal/events/telegram/tgTypes"
)

func Map(post *ParsedPost) (tgTypes.TweetThread, error) {
	tweet := tgTypes.TweetThread{}
	if len(post.Result) == 0 {
		return tweet, errors.New("empty post.Result")
	}

	mediaCode := "name"
	text := ""
	media := tgTypes.Media{}
	for i, result := range post.Result {
		if i == 0 {
			tweet.Time = time.Unix(int64(result.Meta.TakenAt), 0)
			tweet.Likes = result.Meta.LikeCount
			tweet.Replies = result.Meta.CommentCount

			text = result.Meta.Title
			mediaCode = result.Meta.Shortcode
		}

		for j, url := range result.Urls {
			switch url.Extension {
			case "jpg":
				media.Photos = append(media.Photos, tgTypes.MediaObject{
					Name:       fmt.Sprintf("%s_%d_%d.jpg", mediaCode, i, j),
					Url:        url.Url,
					NeedUpload: false,
				})
			case "mp4":
				media.Videos = append(media.Videos, tgTypes.MediaObject{
					Name:       fmt.Sprintf("%s_%d_%d.mp4", mediaCode, i, j),
					Url:        url.Url,
					NeedUpload: true,
				})
			}
		}
	}

	for _, photoChunk := range chunkMedia(media.Photos, 10) {
		content := tgTypes.TweetContent{}
		mediaChunk := tgTypes.Media{}

		mediaChunk.Photos = photoChunk

		content.Media = mediaChunk
		tweet.Tweets = append(tweet.Tweets, content)
	}

	for _, videoChunk := range chunkMedia(media.Videos, 10) {
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
