package socialapi1instagram

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"time"
	"tweets-tg-bot/internal/downloader"
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

	var additionalContents []tgTypes.TweetContent

	switch contentType {
	case "story":
		if !post.Data.IsVideo {
			mediaObject, err := getMediaFromImageVersions(post.Data.ImageVersions)
			if err != nil {
				return tweet, errors.Wrap(err, "error getting media object from image")
			}
			media.Photos = append(media.Photos, mediaObject)
		} else {
			mediaObject, err := getMediaFromVideoVersions(post.Data.VideoVersions, post.Data.VideoUrl)
			if err != nil {
				return tweet, errors.Wrap(err, "error getting media from video")
			}
			media.Videos = append(media.Videos, mediaObject)
		}
	case "post":
		mediaObject, err := getMediaFromImageVersions(post.Data.ImageVersions)
		if err != nil {
			return tweet, errors.Wrap(err, "error getting media object from image")
		}
		media.Photos = append(media.Photos, mediaObject)
	case "album":
		photos, err := getMediaFromCarousel(post.Data.CarouselMedia)
		if err != nil {
			return tweet, errors.Wrap(err, "error getting photo object from carousel")
		}
		chunkedPhotos := chunkMedia(photos, 10)
		for i, chunk := range chunkedPhotos {
			if i == 0 {
				media.Photos = append(media.Photos, chunk...)
			} else {
				additionalContents = append(additionalContents, tgTypes.TweetContent{Media: tgTypes.Media{Photos: chunk}})
			}
		}
	case "reel":
		mediaObject, err := getMediaFromVideoVersions(post.Data.VideoVersions, post.Data.VideoUrl)
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

	if author != nil {
		tweet.UserName = author.FullName
		tweet.UserId = author.Username
	}

	if post.Data.Caption != nil {
		content.Text = post.Data.Caption.Text
	}
	content.Media = media
	tweet.Tweets = append(tweet.Tweets, content)

	if len(additionalContents) > 0 {
		tweet.Tweets = append(tweet.Tweets, additionalContents...)
	}

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
				Name:       "image",
				Url:        image.Url,
				NeedUpload: false,
			}
		}
	}

	if maxSize == 0 {
		return media, errors.New("no images found")
	}

	return media, nil
}

func getMediaFromCarousel(carousel []CarouselMedia) ([]tgTypes.MediaObject, error) {
	var result []tgTypes.MediaObject

	for _, m := range carousel {
		media, err := getMediaFromImageVersions(m.ImageVersions)
		if err != nil {
			return nil, errors.Wrap(err, "error getting media from image")
		}
		result = append(result, media)
	}
	return result, nil
}

func getMediaFromVideoVersions(v []*VideoVersion, videoUrl string) (tgTypes.MediaObject, error) {
	maxSize := 0
	media := tgTypes.MediaObject{}

	maxFileSize := uint64(50 * 1024 * 1024)

	if videoUrl != "" {
		fileSize, err := downloader.FileSize(videoUrl)
		if err == nil && fileSize <= maxFileSize {
			return tgTypes.MediaObject{
				Name:       "video",
				Url:        videoUrl,
				NeedUpload: true,
			}, nil
		}
	}

	for _, video := range v {
		size := video.Height * video.Width
		fileSize, err := downloader.FileSize(video.Url)
		if err != nil {
			continue
		}
		if size > maxSize && fileSize <= maxFileSize {
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
