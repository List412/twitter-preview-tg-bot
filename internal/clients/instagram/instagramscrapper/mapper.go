package instagramscrapper

import (
	"fmt"
	"github.com/list412/twitter-preview-tg-bot/internal/downloader"
	"github.com/list412/twitter-preview-tg-bot/internal/events/telegram/tgTypes"
	"github.com/pkg/errors"
	"time"
)

func Map(post *ParsedPost) (tgTypes.TweetThread, error) {
	tweet := tgTypes.TweetThread{}
	content := tgTypes.TweetContent{}
	media := tgTypes.Media{}

	author := post.User
	if author.Username == "" {
		author = post.Owner
	}

	contentType := post.ProductType

	var additionalContents []tgTypes.TweetContent

	switch contentType {
	case "story":
		if post.VideoDuration == 0 {
			mediaObject, err := getMediaFromImageVersions(post.ImageVersions2.Candidates)
			if err != nil {
				return tweet, errors.Wrap(err, "error getting media object from image")
			}
			media.Photos = append(media.Photos, mediaObject)
		} else {
			mediaObject, err := getMediaFromVideoVersions(post.VideoVersions)
			if err != nil {
				return tweet, errors.Wrap(err, "error getting media from video")
			}
			media.Videos = append(media.Videos, mediaObject)
		}
	case "feed":
		fallthrough
	case "post":
		mediaObject, err := getMediaFromImageVersions(post.ImageVersions2.Candidates)
		if err != nil {
			return tweet, errors.Wrap(err, "error getting media object from image")
		}
		media.Photos = append(media.Photos, mediaObject)
	case "carousel_container":
		fallthrough
	case "album":
		photos, err := getMediaFromCarousel(post.CarouselMedia)
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
	case "clips":
		fallthrough
	case "reel":
		mediaObject, err := getMediaFromVideoVersions(post.VideoVersions)
		if err != nil {
			return tweet, errors.Wrap(err, "error getting media from video")
		}
		media.Videos = append(media.Videos, mediaObject)
	default:

	}

	t := time.Unix(int64(post.TakenAt), 0)
	tweet.Time = t

	tweet.UserName = author.FullName
	tweet.UserId = author.Username

	content.Text = post.Caption.Text

	content.Media = media
	tweet.Tweets = append(tweet.Tweets, content)

	if len(additionalContents) > 0 {
		tweet.Tweets = append(tweet.Tweets, additionalContents...)
	}

	return tweet, nil
}

func getMediaFromImageVersions(v []ImageCandidate) (tgTypes.MediaObject, error) {
	maxSize := 0
	media := tgTypes.MediaObject{}

	for _, image := range v {
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
		media, err := getMediaFromImageVersions(m.ImageVersions2.Candidates)
		if err != nil {
			return nil, errors.Wrap(err, "error getting media from image")
		}
		result = append(result, media)
	}
	return result, nil
}

func getMediaFromVideoVersions(v []VideoVersion) (tgTypes.MediaObject, error) {
	maxSize := 0
	media := tgTypes.MediaObject{}

	maxFileSize := uint64(50 * 1024 * 1024)

	for _, video := range v {
		size := video.Height * video.Width
		fileSize, err := downloader.FileSize(video.Url)
		if err != nil {
			continue
		}
		if size > maxSize && fileSize <= maxFileSize {
			maxSize = size
			media = tgTypes.MediaObject{
				Name:       "video",
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
