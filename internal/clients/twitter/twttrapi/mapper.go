package twttrapi

import (
	"github.com/pkg/errors"
	"strings"
	"time"
	"tweets-tg-bot/internal/downloader"
	"tweets-tg-bot/internal/events/telegram"
	"tweets-tg-bot/internal/events/telegram/tgTypes"
)

func Map(parsedTweet *ParsedThread) (tgTypes.TweetThread, error) {
	if parsedTweet.Errors != nil || parsedTweet.Error != nil {
		return tgTypes.TweetThread{}, telegram.ErrApiResponse
	}

	tweet := tgTypes.TweetThread{}

	tweetResult := getTweetData(parsedTweet)

	tweetTime, err := time.Parse("Mon Jan 02 15:04:05 -0700 2006", tweetResult.CreatedAt)
	if err != nil {
		tweetTime = time.Now()
	}
	tweet.Time = tweetTime
	tweet.Likes = tweetResult.FavoriteCount
	tweet.Quotes = tweetResult.QuoteCount
	tweet.Retweets = tweetResult.RetweetCount
	tweet.Replies = tweetResult.ReplyCount
	tweet.Views = getViewsCount(parsedTweet)

	userResult := getUserData(parsedTweet)
	tweet.UserName = userResult.Name
	tweet.UserId = userResult.ScreenName

	entries := getEntries(parsedTweet)
	content, err := parseEntries(entries)
	if err != nil {
		return tgTypes.TweetThread{}, errors.Wrap(err, "error parsing entries")
	}

	tweet.Tweets = content

	return tweet, nil
}

func parseEntries(entries []Entity) ([]tgTypes.TweetContent, error) {
	var results []tgTypes.TweetContent
	for i, entry := range entries {
		if i == 0 || entry.Content.Content.TweetDisplayType == "SelfThread" {
			tweetContent := tgTypes.TweetContent{}
			tweet := getTweetFromEntity(entry, i)
			media, err := getMedia(*tweet.Legacy)
			if err != nil {
				return nil, errors.Wrap(err, "could not get media")
			}
			tweetContent.Media = media
			text := getText(tweet)

			tweetContent.Text = replaceUrlInText(text, getResultFromTweet(tweet))

			results = append(results, tweetContent)
		}

		if len(entry.Content.Items) > 0 && entry.Content.Items[0].Item.Content.TweetDisplayType == "SelfThread" {
			for itemIndex, _ := range entry.Content.Items {
				tweetContent := tgTypes.TweetContent{}
				tweet := getTweetFromEntity(entry, itemIndex)
				media, err := getMedia(*tweet.Legacy)
				if err != nil {
					return nil, errors.Wrap(err, "could not get media")
				}
				tweetContent.Media = media
				text := getText(tweet)

				tweetContent.Text = replaceUrlInText(text, getResultFromTweet(tweet))

				results = append(results, tweetContent)
			}
		}
	}
	return results, nil
}

func getMedia(tweet TweetData) (tgTypes.Media, error) {
	result := tgTypes.Media{}

	for _, media := range tweet.ExtendedEntities.Media {
		switch media.Type {
		case "photo":
			result.Photos = append(result.Photos, tgTypes.MediaObject{Url: media.MediaUrlHttps})
		case "animated_gif":
			fallthrough
		case "video":
			variant, err := chooseVideoVariant(media.VideoInfo.Variants)
			if err != nil {
				return tgTypes.Media{}, err
			}
			variant.Name = media.MediaKey
			result.Videos = append(result.Videos, *variant)
		}
	}

	return result, nil
}

func getEntries(tweet *ParsedThread) []Entity {
	return tweet.Data.TimelineResponse.Instructions[0].Entries
}

func getUserData(tweet *ParsedThread) UserData {
	return getMainTweet(tweet).Core.UserResult.Result.Legacy
}

func getTweetData(tweet *ParsedThread) TweetData {
	return *getMainTweet(tweet).Legacy
}

func getViewsCount(tweet *ParsedThread) string {
	return getMainTweet(tweet).ViewCountInfo.Count
}

func getMainTweet(tweet *ParsedThread) *Tweet {
	return tweet.Data.TimelineResponse.Instructions[0].Entries[0].Content.Content.TweetResult.Result
}

func getTweetFromEntity(entity Entity, index int) *Tweet {
	if entity.Content.Content.TweetResult.Result != nil {
		return entity.Content.Content.TweetResult.Result
	}
	return entity.Content.Items[index].Item.Content.TweetResult.Result
}

func getResultFromTweet(tw *Tweet) *TweetData {
	if tw.Core == nil && tw.Tweet != nil {
		tw = tw.Tweet
	}

	return tw.Legacy
}

func getText(tw *Tweet) string {
	if tw.Core == nil && tw.Tweet != nil {
		tw = tw.Tweet
	}

	if tw.NoteTweet != nil {
		return tw.NoteTweet.NoteTweetResults.Result.Text
	}

	return tw.Legacy.FullText
}

func replaceUrlInText(text string, tweet *TweetData) string {
	if len(tweet.Entities.Urls) == 0 {
		return text
	}

	for _, url := range tweet.Entities.Urls {
		text = strings.Replace(text, url.Url, url.ExpandedUrl, 1)
	}

	return text
}

func chooseVideoVariant(variants []Variant) (*tgTypes.MediaObject, error) {
	for i := len(variants) - 1; i >= 0; i-- {
		if variants[i].ContentType == "video/mp4" {
			size, err := downloader.FileSize(variants[i].Url)
			if err != nil {
				return nil, err
			}

			if size <= 50*1024*1024 {
				return &tgTypes.MediaObject{
					Url:        variants[i].Url,
					NeedUpload: true,
				}, nil
			}
		}
	}
	return nil, nil
}
