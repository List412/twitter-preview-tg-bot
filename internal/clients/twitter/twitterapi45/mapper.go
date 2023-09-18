package twitterapi45

import (
	"strconv"
	"time"
	"tweets-tg-bot/internal/events/telegram/tgTypes"
)

func Map(parsedTweet *Response) (tgTypes.Tweet, error) {
	tweet := tgTypes.Tweet{}

	tweet.Text = parsedTweet.Text
	tweet.Likes = parsedTweet.Likes
	tweet.Quotes = parsedTweet.Quotes
	tweet.Replies = parsedTweet.Replies
	tweet.Retweets = parsedTweet.Retweets
	tweet.UserId = parsedTweet.Author.Name
	tweet.UserName = parsedTweet.Author.ScreenName
	//tweet.Bookmarks = parsedTweet.Bookmarks todo?

	tweetTime, err := time.Parse("Mon Jan 02 15:04:05 -0700 2006", parsedTweet.CreatedAt)
	if err != nil {
		tweetTime = time.Now()
	}
	tweet.Time = tweetTime

	for _, photo := range parsedTweet.Media.Photo {
		media := tgTypes.MediaObject{
			Url: photo.MediaUrlHttps,
		}
		tweet.Media.Photos = append(tweet.Media.Photos, media)
	}

	for i, video := range parsedTweet.Media.Video {
		media := tgTypes.MediaObject{
			Name:       video.Variants[1].ContentType + "_" + strconv.Itoa(i),
			Url:        video.Variants[1].Url,
			NeedUpload: true,
		}
		tweet.Media.Videos = append(tweet.Media.Photos, media)
	}

	return tweet, nil
}
