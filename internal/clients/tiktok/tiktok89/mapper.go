package tiktok89

import (
	"fmt"
	"github.com/pkg/errors"
	"time"
	"tweets-tg-bot/internal/events/telegram/tgTypes"
)

func Map(parsedVideo *VideoParsed) (tgTypes.TweetThread, error) {
	tweet := tgTypes.TweetThread{}

	stats := parsedVideo.Statistics
	author := parsedVideo.Author
	video := parsedVideo.Video

	t := time.Unix(int64(parsedVideo.CreateTime), 0)
	tweet.Time = t
	tweet.Likes = stats.DiggCount
	tweet.Views = fmt.Sprintf("%d", stats.PlayCount)
	tweet.Replies = stats.CommentCount
	tweet.Retweets = stats.RepostCount

	tweet.UserName = author.Nickname
	tweet.UserId = author.Uid

	content := tgTypes.TweetContent{}
	content.Text = parsedVideo.Desc
	media := tgTypes.Media{}

	videoUrl, err := videoVariants(video)
	if err != nil {
		return tweet, errors.Wrap(err, "error getting video url")
	}

	media.Videos = append(media.Videos, tgTypes.MediaObject{
		Name:       "def",
		Url:        videoUrl,
		NeedUpload: true,
	})

	content.Media = media

	tweet.Tweets = append(tweet.Tweets, content)

	return tweet, nil
}

func videoVariants(video Video) (string, error) {
	for _, br := range video.BitRate {
		if br.PlayAddr.DataSize <= 50*1024*1024 {
			if len(br.PlayAddr.UrlList) > 0 {
				return br.PlayAddr.UrlList[0], nil
			}
		}
	}

	if len(video.DownloadAddr.UrlList) > 0 {
		return video.DownloadAddr.UrlList[0], nil
	} else if len(video.PlayAddr.UrlList) > 0 {
		return video.PlayAddr.UrlList[0], nil
	}

	return "", fmt.Errorf("no video url")
}
