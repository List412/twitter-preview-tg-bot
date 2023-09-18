package twttrapi

import (
	"time"
	"tweets-tg-bot/internal/events/telegram"
	"tweets-tg-bot/internal/events/telegram/tgTypes"
)

func Map(parsedTweet *ParsedTweet) (tgTypes.Tweet, error) {
	if parsedTweet.Errors != nil || parsedTweet.Error != nil {
		return tgTypes.Tweet{}, telegram.ErrApiResponse
	}

	tweet := tgTypes.Tweet{Media: tgTypes.Media{}}

	tweetResult := getTweetData(parsedTweet)

	tweet.Text = getText(parsedTweet)

	tweetTime, err := time.Parse("Mon Jan 02 15:04:05 -0700 2006", tweetResult.CreatedAt)
	if err != nil {
		tweetTime = time.Now()
	}
	tweet.Time = tweetTime

	for _, media := range tweetResult.ExtendedEntities.Media {
		switch media.Type {
		case "photo":
			tweet.Media.Photos = append(tweet.Media.Photos, tgTypes.MediaObject{Url: media.MediaUrlHttps})
		case "animated_gif":
			fallthrough
		case "video":
			for i := len(media.VideoInfo.Variants) - 1; i >= 0; i-- {
				if media.VideoInfo.Variants[i].ContentType == "video/mp4" {
					tweet.Media.Videos = append(tweet.Media.Videos, tgTypes.MediaObject{
						Name:       media.MediaKey,
						Url:        media.VideoInfo.Variants[i].Url,
						Data:       nil,
						NeedUpload: true,
					})
					break
				}
			}
		}
	}

	tweet.Likes = tweetResult.FavoriteCount
	tweet.Quotes = tweetResult.QuoteCount
	tweet.Retweets = tweetResult.RetweetCount
	tweet.Replies = tweetResult.ReplyCount
	tweet.Views = getViewsCount(parsedTweet)

	userResult := getUserData(parsedTweet)
	tweet.UserName = userResult.Name
	tweet.UserId = userResult.ScreenName

	return tweet, nil
}

func getUserData(tweet *ParsedTweet) UserData {
	if tweet.Data.TweetResult.Result.Core != nil {
		return tweet.Data.TweetResult.Result.Core.UserResult.Result.Legacy
	}
	return tweet.Data.TweetResult.Result.Tweet.Core.UserResult.Result.Legacy
}

func getTweetData(tweet *ParsedTweet) TweetData {
	if tweet.Data.TweetResult.Result.Legacy != nil {
		return *tweet.Data.TweetResult.Result.Legacy
	}
	return *tweet.Data.TweetResult.Result.Tweet.Legacy
}

func getText(tweet *ParsedTweet) string {
	tw := tweet.Data.TweetResult.Result
	if tw.Core == nil && tw.Tweet != nil {
		tw = tw.Tweet
	}

	if tw.NoteTweet != nil {
		return tw.NoteTweet.NoteTweetResults.Result.Text
	}

	return tw.Legacy.FullText
}

func getViewsCount(tweet *ParsedTweet) string {
	tw := tweet.Data.TweetResult.Result
	if tw.Core == nil && tw.Tweet != nil {
		tw = tw.Tweet
	}
	return tw.ViewCountInfo.Count
}
