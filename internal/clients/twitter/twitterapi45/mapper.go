package twitterapi45

import (
	"strconv"
	"time"
	"tweets-tg-bot/internal/downloader"
	"tweets-tg-bot/internal/events/telegram/tgTypes"
)

func Map(parsedTweet *Response) (tgTypes.TweetThread, error) {
	tweet := tgTypes.TweetThread{}
	tweetContent := tgTypes.TweetContent{}

	tweetContent.Text = parsedTweet.Text
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
			Url:        photo.MediaUrlHttps,
			Name:       photo.Id,
			NeedUpload: false,
		}
		tweetContent.Media.Photos = append(tweetContent.Media.Photos, media)
	}

	for i, video := range parsedTweet.Media.Video {
		variant, err := getVideoFromVariants(video.Variants)
		if err != nil {
			return tgTypes.TweetThread{}, err
		}
		media := tgTypes.MediaObject{
			Name:       variant.ContentType + "_" + strconv.Itoa(i),
			Url:        variant.Url,
			NeedUpload: true,
		}
		tweetContent.Media.Videos = append(tweetContent.Media.Videos, media)
	}

	tweet.Tweets = append(tweet.Tweets, tweetContent)
	return tweet, nil
}

func getVideoFromVariants(variants []Variant) (Variant, error) {
	for i := len(variants) - 1; i >= 0; i-- {
		if variants[i].ContentType == "video/mp4" {
			size, err := downloader.FileSize(variants[i].Url)
			if err != nil {
				return Variant{}, err
			}

			if size <= 50*1024*1024 {
				return variants[i], nil
			}
		}
	}
	return variants[0], nil
}
