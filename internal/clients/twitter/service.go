package twitter

import (
	"time"
	"tweets-tg-bot/internal/clients/twitter/twttrapi"
	"tweets-tg-bot/internal/events/telegram"
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

func (s Service) GetTweet(id string) (telegram.Tweet, error) {
	parsedTweet, err := s.client.GetTweet(id)
	if err != nil {
		return telegram.Tweet{}, err
	}

	tweet := telegram.Tweet{Media: telegram.Media{}}
	tweetResult := parsedTweet.Data.TweetResult.Result.Legacy
	tweet.Text = tweetResult.FullText

	tweet.Time, err = time.Parse("Mon Jan 02 15:04:05 -0700 2006", tweetResult.CreatedAt)
	if err != nil {
		tweet.Time = time.Now()
	}

	for _, media := range tweetResult.ExtendedEntities.Media {
		switch media.Type {
		case "photo":
			tweet.Media.Photos = append(tweet.Media.Photos, media.ExpandedUrl)
		case "video":
			for i := len(media.VideoInfo.Variants) - 1; i >= 0; i-- {
				if media.VideoInfo.Variants[i].ContentType == "video/mp4" {
					tweet.Media.Videos = append(tweet.Media.Videos, media.VideoInfo.Variants[i].Url)
					break
				}
			}
		}
	}

	tweet.Likes = tweetResult.FavoriteCount
	tweet.Quotes = tweetResult.QuoteCount
	tweet.Retweets = tweetResult.RetweetCount
	tweet.Replies = tweetResult.ReplyCount
	tweet.Views = parsedTweet.Data.TweetResult.Result.ViewCountInfo.Count

	userResult := parsedTweet.Data.TweetResult.Result.Core.UserResult.Result.Legacy
	tweet.UserName = userResult.Name
	tweet.UserId = userResult.ScreenName

	return tweet, nil
}
