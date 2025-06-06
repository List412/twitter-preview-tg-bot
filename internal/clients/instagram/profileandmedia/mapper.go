package profileandmedia

import (
	"fmt"
	"github.com/list412/tweets-tg-bot/internal/events/telegram/tgTypes"
	"github.com/pkg/errors"
	"time"
)

func Map(post *ParsedPost) (tgTypes.TweetThread, error) {
	tweet := tgTypes.TweetThread{}
	if post.Graphql.ShortcodeMedia.Id == "" {
		return tweet, errors.New("no graphql media provided")
	}

	media := tgTypes.Media{}
	for _, m := range post.Graphql.ShortcodeMedia.EdgeSidecarToChildren.Edges {
		if m.Node.IsVideo {
			media.Videos = append(media.Videos, tgTypes.MediaObject{
				Name:       m.Node.Id,
				Url:        m.Node.VideoUrl,
				NeedUpload: true,
			})
		} else {
			media.Photos = append(media.Photos, tgTypes.MediaObject{
				Name:       m.Node.Id,
				Url:        m.Node.DisplayUrl,
				NeedUpload: false,
			})
		}
	}

	if len(media.Videos) == 0 && len(media.Photos) == 0 {
		if post.Graphql.ShortcodeMedia.IsVideo {
			media.Videos = append(media.Videos, tgTypes.MediaObject{
				Name:       post.Graphql.ShortcodeMedia.Id,
				Url:        post.Graphql.ShortcodeMedia.VideoUrl,
				NeedUpload: true,
			})
			tweet.Views = fmt.Sprintf("%d", post.Graphql.ShortcodeMedia.VideoViewCount)
		} else {
			media.Photos = append(media.Photos, tgTypes.MediaObject{
				Name:       post.Graphql.ShortcodeMedia.Id,
				Url:        post.Graphql.ShortcodeMedia.DisplayUrl,
				NeedUpload: false,
			})
		}
	}

	owner := post.Graphql.ShortcodeMedia.Owner
	tweet.UserName = owner.FullName
	tweet.UserId = owner.Username

	tweet.Time = time.Unix(int64(post.Graphql.ShortcodeMedia.TakenAtTimestamp), 64)
	tweet.Likes = post.Graphql.ShortcodeMedia.EdgeMediaPreviewLike.Count
	tweet.Replies = post.Graphql.ShortcodeMedia.EdgeMediaPreviewComment.Count

	text := ""
	if len(post.Graphql.ShortcodeMedia.EdgeMediaToCaption.Edges) > 0 {
		text = post.Graphql.ShortcodeMedia.EdgeMediaToCaption.Edges[0].Node.Text
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
