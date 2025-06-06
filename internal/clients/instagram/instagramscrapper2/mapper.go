package instagramscrapper2

import (
	"fmt"
	"strconv"
	"time"

	"github.com/list412/tweets-tg-bot/internal/events/telegram/tgTypes"
)

func Map(post *ParsedPost) (tgTypes.TweetThread, error) {
	tweet := tgTypes.TweetThread{}

	data := post.Data.ShortcodeMedia
	owner := data.Owner

	tweet.UserId = owner.Username
	tweet.UserName = owner.FullName

	tweet.Views = strconv.Itoa(data.VideoViewCount)
	tweet.Likes = data.EdgeMediaPreviewLike.Count

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
			Name:       fmt.Sprintf("%s_%s.mp4", data.Id, data.Shortcode),
			Url:        data.VideoUrl,
			NeedUpload: true,
		})
	} else {
		media.Photos = append(media.Photos, tgTypes.MediaObject{
			Name:       fmt.Sprintf("%s_%s.jpg", data.Id, data.Shortcode),
			Url:        data.DisplayUrl,
			NeedUpload: false,
		})
	}

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
