package scrapper

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func NewClient(host string) *Client {
	return &Client{
		host:   host,
		client: http.Client{},
	}
}

const userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:105.0) Gecko/20100101 Firefox/105.0"
const guestAuthToken = "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA"
const tweetIdPlaceholder = "{tweet_id_here}"

const variables = "{\n        \"focalTweetId\": \"{tweet_id_here}\",\n        \"with_rux_injections\": false,\n        \"includePromotedContent\": false,\n        \"withCommunity\": false,\n        \"withQuickPromoteEligibilityTweetFields\": false,\n        \"withBirdwatchNotes\": false,\n        \"withSuperFollowsUserFields\": false,\n        \"withDownvotePerspective\": false,\n        \"withReactionsMetadata\": false,\n        \"withReactionsPerspective\": false,\n        \"withSuperFollowsTweetFields\": false,\n        \"withVoice\": false,\n        \"withV2Timeline\": false\n    }"
const features = "{\n        \"verified_phone_label_enabled\": false,\n        \"responsive_web_graphql_timeline_navigation_enabled\": true,\n        \"unified_cards_ad_metadata_container_dynamic_card_content_query_enabled\": true,\n        \"tweetypie_unmention_optimization_enabled\": true,\n        \"responsive_web_uc_gql_enabled\": true,\n        \"vibe_api_enabled\": true,\n        \"responsive_web_edit_tweet_api_enabled\": true,\n        \"graphql_is_translatable_rweb_tweet_is_translatable_enabled\": false,\n        \"standardized_nudges_misinfo\": true,\n        \"tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled\": false,\n        \"interactive_text_enabled\": true,\n        \"responsive_web_text_conversations_enabled\": false,\n        \"responsive_web_enhance_cards_enabled\": true\n    }"

type Client struct {
	host   string
	token  string
	client http.Client
}

func (c *Client) GetTweetSelfReplays(id string) (*PageScrapperResult, error) {
	page, err := c.GetTweetPage(id)
	if err != nil {
		return nil, err
	}

	log.Printf("Scrapp tweet done %s", id)

	var tweet TweetHuge
	err = json.Unmarshal(page, &tweet)
	if err != nil {
		return nil, err
	}

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

func (c *Client) GetTweetPage(id string) ([]byte, error) {
	err := c.updateToken() // todo update token on error
	if err != nil {
		return nil, err
	}

	u := url.URL{
		Scheme: "https",
		Host:   c.host,
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

	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("Authorization", guestAuthToken)
	req.Header.Add("x-guest-token", c.token)

	log.Printf("%s %s %s\n", req.RemoteAddr, req.Method, req.URL)
	response, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) updateToken() error {
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return err
	}

	req.Header.Add("User-Agent", userAgent)

	response, err := c.client.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(response.Body)

	reg, err := regexp.Compile("gt=[0-9]*")
	if err != nil {
		return err
	}
	gt := reg.FindString(string(body))

	if len(gt) > 4 {
		gt = gt[3:]
	}

	// todo mutex? nah
	c.token = gt

	return nil
}
