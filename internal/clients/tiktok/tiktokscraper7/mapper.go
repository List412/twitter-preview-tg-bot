package tiktokscraper7

import (
	"fmt"
	"github.com/pkg/errors"
	"time"
	"tweets-tg-bot/internal/events/telegram/tgTypes"
)

func Map(parsedVideo *VideoParsed, id string) (tgTypes.TweetThread, error) {
	tweet := tgTypes.TweetThread{}

	video := parsedVideo.Data

	author := video.Author

	t := time.Unix(int64(video.CreateTime), 0)
	tweet.Time = t
	tweet.Likes = video.DiggCount
	tweet.Views = fmt.Sprintf("%d", video.PlayCount)
	tweet.Replies = video.CommentCount
	tweet.Retweets = video.ShareCount

	tweet.UserName = author.Nickname
	tweet.UserId = author.UniqueId

	content := tgTypes.TweetContent{}
	content.Text = video.Title
	media := tgTypes.Media{}

	videoUrl := video.Play
	size := video.Size

	maxSize := 50 * 1024 * 1024
	if video.WmSize <= maxSize {
		videoUrl = video.Wmplay
		size = video.WmSize
	}

	if video.HdSize <= maxSize {
		videoUrl = video.Hdplay
		size = video.HdSize
	}

	if size == 0 {
		return tgTypes.TweetThread{}, errors.New("video is empty")
	}

	media.Videos = append(media.Videos, tgTypes.MediaObject{
		Name:       video.Id,
		Url:        videoUrl,
		NeedUpload: true,
	})

	content.Media = media

	tweet.Tweets = append(tweet.Tweets, content)

	return tweet, nil
}
