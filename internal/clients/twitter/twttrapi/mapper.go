package twttrapi

import (
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/list412/twitter-preview-tg-bot/internal/events/telegram"
	"github.com/list412/twitter-preview-tg-bot/internal/events/telegram/tgTypes"
)

type Downloader interface {
	FileSize(url string) (uint64, error)
	Download(urls []tgTypes.MediaObject) ([]tgTypes.MediaObject, error)
}

type Mapper struct {
	Downloader Downloader
}

func (m Mapper) Map(parsedTweet *ParsedThread, id string) (tgTypes.TweetThread, error) {
	if parsedTweet.Errors != nil || parsedTweet.Error != nil {
		errMsg := ""
		if parsedTweet.Error != nil {
			errMsg = *parsedTweet.Error
		}
		if errMsg != "" && len(parsedTweet.Errors) > 0 {
			errMsg = parsedTweet.Errors[0].Message
		}
		return tgTypes.TweetThread{}, errors.Wrap(telegram.ErrApiResponse, errMsg)
	}

	tweet := tgTypes.TweetThread{}

	currentEntry, entryId, err := m.getCurrentEntry(parsedTweet, id)
	if err != nil {
		return tgTypes.TweetThread{}, err
	}

	tweetResult := m.getTweetData(currentEntry)

	tweetTime, err := time.Parse("Mon Jan 02 15:04:05 -0700 2006", tweetResult.CreatedAt)
	if err != nil {
		tweetTime = time.Now()
	}
	tweet.Time = tweetTime
	tweet.Likes = tweetResult.FavoriteCount
	tweet.Quotes = tweetResult.QuoteCount
	tweet.Retweets = tweetResult.RetweetCount
	tweet.Replies = tweetResult.ReplyCount
	tweet.Views = m.getViewsCount(currentEntry)

	userResult := m.getUserData(currentEntry)
	tweet.UserName = userResult.Name
	tweet.UserId = userResult.ScreenName

	entries, err := m.getEntries(parsedTweet)
	if err != nil {
		return tgTypes.TweetThread{}, errors.Wrap(err, "no entries")
	}
	content, err := m.parseEntries(entries, entryId)
	if err != nil {
		return tgTypes.TweetThread{}, errors.Wrap(err, "error parsing entries")
	}

	tweet.Tweets = content

	userNote := m.getMainTweet(currentEntry).BirdwatchPivot
	if userNote != nil {
		tweet.UserNote = tgTypes.UserNote{Text: userNote.Note.Summary.Text, Title: userNote.Shorttitle}
	}

	return tweet, nil
}

func (m Mapper) parseEntries(entries []Entity, entryId int) ([]tgTypes.TweetContent, error) {
	var results []tgTypes.TweetContent
	for i, entry := range entries {
		if i == entryId || (entry.Content.Content.TweetDisplayType == "SelfThread" && entryId == 0) {
			tweetContent := tgTypes.TweetContent{}
			tweet := m.getTweetFromEntity(entry, i)
			media, err := m.getMedia(*tweet.Legacy)
			if err != nil {
				return nil, errors.Wrap(err, "could not get media")
			}
			tweetContent.Media = media
			text := m.getText(tweet)

			tweetContent.Text = m.replaceUrlInText(text, m.getResultFromTweet(tweet))

			results = append(results, tweetContent)
		}

		if len(entry.Content.Items) > 0 && entry.Content.Items[0].Item.Content.TweetDisplayType == "SelfThread" && entryId == 0 {
			for itemIndex, _ := range entry.Content.Items {
				tweetContent := tgTypes.TweetContent{}
				tweet := m.getTweetFromEntity(entry, itemIndex)
				media, err := m.getMedia(*tweet.Legacy)
				if err != nil {
					return nil, errors.Wrap(err, "could not get media")
				}
				tweetContent.Media = media
				text := m.getText(tweet)

				tweetContent.Text = m.replaceUrlInText(text, m.getResultFromTweet(tweet))

				results = append(results, tweetContent)
			}
		}
	}
	return results, nil
}

func (m Mapper) getMedia(tweet TweetData) (tgTypes.Media, error) {
	result := tgTypes.Media{}

	for _, media := range tweet.ExtendedEntities.Media {
		switch media.Type {
		case "photo":
			result.Photos = append(result.Photos, tgTypes.MediaObject{
				Url:        media.MediaUrlHttps,
				Name:       media.MediaKey,
				NeedUpload: true,
			})
		case "animated_gif":
			fallthrough
		case "video":
			variant, err := m.chooseVideoVariant(media.VideoInfo.Variants)
			if err != nil {
				return tgTypes.Media{}, err
			}
			variant.Name = media.MediaKey
			result.Videos = append(result.Videos, *variant)
		}
	}

	return result, nil
}

func (m Mapper) getEntries(tweet *ParsedThread) ([]Entity, error) {
	if len(tweet.Data.TimelineResponse.Instructions) > 0 && tweet.Data.TimelineResponse.Instructions[0].Entries != nil {
		return tweet.Data.TimelineResponse.Instructions[0].Entries, nil
	}

	if len(tweet.Data.TimelineResponse.Instructions) > 1 && tweet.Data.TimelineResponse.Instructions[1].Entries != nil {
		return tweet.Data.TimelineResponse.Instructions[1].Entries, nil
	}

	return nil, errors.New("no timeline instructions found")
}

func (m Mapper) getUserData(entry Entity) UserData {
	return m.getMainTweet(entry).Core.UserResult.Result.Legacy
}

func (m Mapper) getCurrentEntry(tweet *ParsedThread, id string) (Entity, int, error) {
	entryId := "tweet-" + id

	entries, err := m.getEntries(tweet)
	if err != nil {
		return Entity{}, 0, errors.Wrap(err, "could not get entries")
	}

	for i, entry := range entries {
		if entryId == entry.EntryId {
			return entry, i, nil
		}
	}
	return Entity{}, 0, errors.New("no entry with provided id")
}

func (m Mapper) getTweetData(entry Entity) TweetData {
	mainTweet := m.getMainTweet(entry)
	return *mainTweet.Legacy
}

func (m Mapper) getViewsCount(entry Entity) string {
	return m.getMainTweet(entry).ViewCountInfo.Count
}

func (m Mapper) getMainTweet(entry Entity) *Tweet {
	mainTweet := entry.Content.Content.TweetResult.Result
	if mainTweet.Legacy == nil {
		mainTweet = mainTweet.Tweet
	}

	return mainTweet
}

func (m Mapper) getTweetFromEntity(entity Entity, index int) *Tweet {
	if entity.Content.Content.TweetResult.Result != nil {
		tweet := entity.Content.Content.TweetResult.Result
		if tweet.Legacy == nil {
			tweet = tweet.Tweet
		}

		return tweet
	}

	tweet := entity.Content.Items[index].Item.Content.TweetResult.Result
	if tweet.Legacy == nil {
		tweet = tweet.Tweet
	}

	return tweet
}

func (m Mapper) getResultFromTweet(tw *Tweet) *TweetData {
	if tw.Core == nil && tw.Tweet != nil {
		tw = tw.Tweet
	}

	return tw.Legacy
}

func (m Mapper) getText(tw *Tweet) string {
	if tw.Core == nil && tw.Tweet != nil {
		tw = tw.Tweet
	}

	if tw.NoteTweet != nil {
		return tw.NoteTweet.NoteTweetResults.Result.Text
	}

	return tw.Legacy.FullText
}

func (m Mapper) replaceUrlInText(text string, tweet *TweetData) string {
	if len(tweet.Entities.Urls) == 0 {
		return text
	}

	for _, url := range tweet.Entities.Urls {
		text = strings.Replace(text, url.Url, url.ExpandedUrl, 1)
	}

	return text
}

func (m Mapper) chooseVideoVariant(variants []Variant) (*tgTypes.MediaObject, error) {
	for i := len(variants) - 1; i >= 0; i-- {
		if variants[i].ContentType == "video/mp4" {
			size, err := m.Downloader.FileSize(variants[i].Url)
			if err != nil {
				continue
			}

			if size <= 50*1024*1024 {
				return &tgTypes.MediaObject{
					Name:       fmt.Sprintf("video_%d", i),
					Url:        variants[i].Url,
					NeedUpload: true,
				}, nil
			}
		}
	}
	return nil, errors.New("could not find video variant")
}
