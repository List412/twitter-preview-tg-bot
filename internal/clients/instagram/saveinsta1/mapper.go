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

	content := tgTypes.TweetContent{}
	media := tgTypes.Media{}

	mediaCode := "name"

	for i, result := range post.Result {
		if i == 0 {
			tweet.Time = time.Unix(int64(result.Meta.TakenAt), 0)
			tweet.Likes = result.Meta.LikeCount
			tweet.Replies = result.Meta.CommentCount

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
					NeedUpload: false,
				})
			}
		}
	}

	content.Media = media
	tweet.Tweets = append(tweet.Tweets, content)

	return tweet, nil
}
