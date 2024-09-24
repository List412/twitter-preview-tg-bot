package tiktok_scaraper2

import (
	"strconv"
	"time"
	"tweets-tg-bot/internal/events/telegram/tgTypes"
)

func Map(parsedVideo *VideoParsed, id string) (tgTypes.TweetThread, error) {
	tweet := tgTypes.TweetThread{}

	stats := parsedVideo.ItemInfo.ItemStruct.StatsV2
	author := parsedVideo.ItemInfo.ItemStruct.Author
	video := parsedVideo.ItemInfo.ItemStruct.Video

	t := time.Unix(int64(parsedVideo.ItemInfo.ItemStruct.CreateTime), 0)
	tweet.Time = t
	tweet.Likes = parseInt(stats.DiggCount)
	tweet.Views = stats.PlayCount
	tweet.Replies = parseInt(stats.CommentCount)
	tweet.Retweets = parseInt(stats.RepostCount)

	tweet.UserName = author.Nickname
	tweet.UserId = author.Id

	content := tgTypes.TweetContent{}
	content.Text = parsedVideo.ItemInfo.ItemStruct.Desc
	media := tgTypes.Media{}
	media.Videos = append(media.Videos, tgTypes.MediaObject{
		Name:       "def",
		Url:        video.DownloadAddr,
		NeedUpload: true,
	})

	content.Media = media

	tweet.Tweets = append(tweet.Tweets, content)

	return tweet, nil
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}
