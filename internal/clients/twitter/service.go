package twitter

import (
	"time"
	"tweets-tg-bot/internal/clients/twitter/twttrapi"
	"tweets-tg-bot/internal/events/telegram"
	"tweets-tg-bot/internal/events/telegram/tgTypes"
)

type Client interface {
	GetTweet(id string) (*twttrapi.ParsedTweet, error)
}

type Service struct {
	client Client
}

func NewService(client Client) *Service {
	return &Service{client: client}
}

func (s Service) GetTweet(id string) (tgTypes.Tweet, error) {
	parsedTweet, err := s.client.GetTweet(id)
	if err != nil {
		return tgTypes.Tweet{}, err
	}
	if parsedTweet.Errors != nil || parsedTweet.Error != nil {
		return tgTypes.Tweet{}, telegram.ErrApiResponse
	}

	tweet := tgTypes.Tweet{Media: tgTypes.Media{}}

	tweetResult := getTweetData(parsedTweet)

	tweet.Text = getText(parsedTweet)

	tweet.Time, err = time.Parse("Mon Jan 02 15:04:05 -0700 2006", tweetResult.CreatedAt)
	if err != nil {
		tweet.Time = time.Now()
	}

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

func getUserData(tweet *twttrapi.ParsedTweet) twttrapi.UserData {
	if tweet.Data.TweetResult.Result.Core != nil {
		return tweet.Data.TweetResult.Result.Core.UserResult.Result.Legacy
	}
	return tweet.Data.TweetResult.Result.Tweet.Core.UserResult.Result.Legacy
}

func getTweetData(tweet *twttrapi.ParsedTweet) twttrapi.TweetData {
	if tweet.Data.TweetResult.Result.Legacy != nil {
		return *tweet.Data.TweetResult.Result.Legacy
	}
	return *tweet.Data.TweetResult.Result.Tweet.Legacy
}

func getText(tweet *twttrapi.ParsedTweet) string {
	tw := tweet.Data.TweetResult.Result
	if tw.Core == nil && tw.Tweet != nil {
		tw = tw.Tweet
	}

	if tw.NoteTweet != nil {
		return tw.NoteTweet.NoteTweetResults.Result.Text
	}

	return tw.Legacy.FullText
}

func getViewsCount(tweet *twttrapi.ParsedTweet) string {
	tw := tweet.Data.TweetResult.Result
	if tw.Core == nil && tw.Tweet != nil {
		tw = tw.Tweet
	}
	return tw.ViewCountInfo.Count
}
