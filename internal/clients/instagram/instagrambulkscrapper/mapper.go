package instagrambulkscrapper

import (
	"fmt"
	"github.com/list412/tweets-tg-bot/internal/events/telegram/tgTypes"
)

func Map(post *ParsedPost) (tgTypes.TweetThread, error) {
	tweet := tgTypes.TweetThread{}
	content := tgTypes.TweetContent{}
	media := tgTypes.Media{}

	author := post.Data.Owner

	var additionalContents []tgTypes.TweetContent

	if post.Data.VideoHd != "" {
		media.Videos = append(media.Videos, tgTypes.MediaObject{
			Name:       "vid",
			Url:        post.Data.VideoHd,
			Data:       nil,
			NeedUpload: false,
		})
	} else if post.Data.ImageHd != "" {
		media.Photos = append(media.Photos, tgTypes.MediaObject{
			Name:       "vid",
			Url:        post.Data.ImageHd,
			Data:       nil,
			NeedUpload: false,
		})
	} else if len(post.Data.ChildMediasHd) > 0 {
		for i, childMedia := range post.Data.ChildMediasHd {
			switch childMedia.Type {
			case "image":
				media.Photos = append(media.Photos, tgTypes.MediaObject{
					Name:       fmt.Sprintf("photo_%d", i),
					Url:        childMedia.Url,
					Data:       nil,
					NeedUpload: false,
				})
			case "video":
				media.Videos = append(media.Videos, tgTypes.MediaObject{
					Name:       fmt.Sprintf("video_%d", i),
					Url:        childMedia.Url,
					Data:       nil,
					NeedUpload: false,
				})
			}
		}
	} else if post.Data.MainMediaHd != "" {
		switch post.Data.MainMediaType {
		case "image":
			media.Photos = append(media.Photos, tgTypes.MediaObject{
				Name:       post.Data.MainMediaType,
				Url:        post.Data.MainMediaHd,
				Data:       nil,
				NeedUpload: false,
			})
		case "video":
			media.Videos = append(media.Videos, tgTypes.MediaObject{
				Name:       post.Data.MainMediaType,
				Url:        post.Data.MainMediaHd,
				Data:       nil,
				NeedUpload: false,
			})
		}
	}

	tweet.UserName = author.FullName
	tweet.UserId = author.Username

	content.Text = post.Data.Caption

	content.Media = media
	tweet.Tweets = append(tweet.Tweets, content)

	if len(additionalContents) > 0 {
		tweet.Tweets = append(tweet.Tweets, additionalContents...)
	}

	return tweet, nil
}
