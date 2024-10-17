package socialapi1instagram

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"time"
	"tweets-tg-bot/internal/events/telegram/tgTypes"
)

func Map(post *ParsedPost) (tgTypes.TweetThread, error) {
	tweet := tgTypes.TweetThread{}
	content := tgTypes.TweetContent{}
	media := tgTypes.Media{}

	stats := post.Data.Metrics
	author := post.Data.User
	if author == nil {
		author = post.Data.Owner
	}

	contentType := post.Data.MediaName

	switch contentType {
	case "post":
		mediaObject, err := getMediaFromImageVersions(post.Data.ImageVersions)
		if err != nil {
			return tweet, errors.Wrap(err, "error getting media object from image")
		}
		media.Photos = append(media.Photos, mediaObject)
	case "video":
	case "reel":
		mediaObject, err := getMediaFromVideoVersions(post.Data.VideoVersions)
		if err != nil {
			return tweet, errors.Wrap(err, "error getting media from video")
		}
		media.Videos = append(media.Videos, mediaObject)
	default:

	}

	t := time.Unix(int64(post.Data.TakenAt), 0)
	tweet.Time = t
	if stats != nil {
		tweet.Likes = stats.LikeCount
		tweet.Views = strconv.Itoa(stats.ViewCount)
		tweet.Replies = stats.CommentCount
		tweet.Retweets = stats.ShareCount
	}

	tweet.UserName = author.FullName
	tweet.UserId = author.Username

	content.Text = post.Data.Caption.Text
	content.Media = media
	tweet.Tweets = append(tweet.Tweets, content)

	return tweet, nil
}

func getMediaFromImageVersions(v *ImageVersion) (tgTypes.MediaObject, error) {
	maxSize := 0
	media := tgTypes.MediaObject{}

	for _, image := range v.Items {
		size := image.Width * image.Height
		if size > maxSize {
			maxSize = size
			media = tgTypes.MediaObject{
				Name: "image",
				Url:  image.Url,
			}
		}
	}

	if maxSize == 0 {
		return media, errors.New("no images found")
	}

	return media, nil
}

func getMediaFromVideoVersions(v []*VideoVersion) (tgTypes.MediaObject, error) {
	maxSize := 0
	media := tgTypes.MediaObject{}
	for _, video := range v {
		size := video.Height * video.Width
		if size > maxSize {
			maxSize = size
			media = tgTypes.MediaObject{
				Name:       video.Id,
				Url:        video.Url,
				NeedUpload: true,
			}
		}
	}

	if maxSize == 0 {
		return media, fmt.Errorf("no video url")
	}

	return media, nil
}
