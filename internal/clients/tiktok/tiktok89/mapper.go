package tiktok89

import (
	"fmt"
	"time"
	"tweets-tg-bot/internal/events/telegram/tgTypes"
)

func Map(parsedVideo *VideoParsed, id string) (tgTypes.TweetThread, error) {
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
	media.Videos = append(media.Videos, tgTypes.MediaObject{
		Name:       "def",
		Url:        video.DownloadAddr.UrlList[0],
		NeedUpload: true,
	})

	content.Media = media

	tweet.Tweets = append(tweet.Tweets, content)

	return tweet, nil
}
