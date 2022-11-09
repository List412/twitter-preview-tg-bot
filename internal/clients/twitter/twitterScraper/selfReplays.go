package twitterScraper

import (
	"github.com/pkg/errors"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const tweetIdPlaceholder = "{tweet_id_here}"

const variables = "{\n        \"focalTweetId\": \"{tweet_id_here}\",\n        \"with_rux_injections\": false,\n        \"includePromotedContent\": false,\n        \"withCommunity\": false,\n        \"withQuickPromoteEligibilityTweetFields\": false,\n        \"withBirdwatchNotes\": false,\n        \"withSuperFollowsUserFields\": false,\n        \"withDownvotePerspective\": false,\n        \"withReactionsMetadata\": false,\n        \"withReactionsPerspective\": false,\n        \"withSuperFollowsTweetFields\": false,\n        \"withVoice\": false,\n        \"withV2Timeline\": false\n    }"
const features = "{\n        \"verified_phone_label_enabled\": false,\n        \"responsive_web_graphql_timeline_navigation_enabled\": true,\n        \"unified_cards_ad_metadata_container_dynamic_card_content_query_enabled\": true,\n        \"tweetypie_unmention_optimization_enabled\": true,\n        \"responsive_web_uc_gql_enabled\": true,\n        \"vibe_api_enabled\": true,\n        \"responsive_web_edit_tweet_api_enabled\": true,\n        \"graphql_is_translatable_rweb_tweet_is_translatable_enabled\": false,\n        \"standardized_nudges_misinfo\": true,\n        \"tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled\": false,\n        \"interactive_text_enabled\": true,\n        \"responsive_web_text_conversations_enabled\": false,\n        \"responsive_web_enhance_cards_enabled\": true\n    }"

func (s Scraper) GetReplaysRequest(id string) (*http.Request, error) {
	u := url.URL{
		Scheme: "https",
		Host:   "twitter.com",
		Path:   "i/api/graphql/zZXycP0V6H7m-2r0mOnFcA/TweetDetail", // todo find somewhere queryId
	}

	q := url.Values{}
	q.Add("variables", strings.Replace(variables, tweetIdPlaceholder, id, 1))
	q.Add("features", features)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = q.Encode()
	return req, nil
}

func (s Scraper) GetTweetSelfReplays(id string) (*PageScrapperResult, error) {
	req, err := s.GetReplaysRequest(id)
	if err != nil {
		return nil, err
	}
	var tweet TweetHuge
	err = s.tw.RequestAPI(req, &tweet)
	if err != nil {
		return nil, err
	}

	log.Printf("Scrapp tweet done %s", id)

	if len(tweet.Data.ThreadedConversationWithInjections.Instructions) < 1 {
		return nil, errors.New("no entries in tweet")
	}

	entries := tweet.Data.ThreadedConversationWithInjections.Instructions[0].Entries

	if len(entries) < 2 {
		return nil, errors.New("no entries with replays")
	}

	selfReplayEntry := entries[1]

	selfReplays := selfReplayEntry.Content.Items

	var scrapperReslt PageScrapperResult

	var result []SelfReplay

	for _, e := range selfReplays {
		legacy := e.Item.ItemContent.TweetResults.Result.Legacy

		if legacy.SelfThread.IdStr == "" {
			continue
		}

		if legacy.SelfThread.IdStr == legacy.IdStr {
			continue
		}

		reply := SelfReplay{}

		reply.Text = legacy.FullText
		reply.Id = legacy.IdStr
		result = append(result, reply)
	}

	var collabs []Collab

	if entries[0].Content.ItemContent.TweetResults.Result.Legacy.CollabControl != nil &&
		len(entries[0].Content.ItemContent.TweetResults.Result.Legacy.CollabControl.CollaboratorsResults) > 1 {
		for _, c := range entries[0].Content.ItemContent.TweetResults.Result.Legacy.CollabControl.CollaboratorsResults {
			collab := Collab{
				Name:       c.Result.Legacy.Name,
				ScreenName: c.Result.Legacy.ScreenName,
			}
			collabs = append(collabs, collab)
		}
	}

	scrapperReslt.SelfReplay = result
	scrapperReslt.CollabUsers = collabs

	return &scrapperReslt, nil
}
