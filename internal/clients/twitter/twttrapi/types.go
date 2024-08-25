package twttrapi

type Variant struct {
	ContentType string `json:"content_type"`
	Url         string `json:"url"`
}

type TweetData struct {
	ConversationIdStr string `json:"conversation_id_str"`
	CreatedAt         string `json:"created_at"`
	DisplayTextRange  []int  `json:"display_text_range"`

	Entities struct {
		Urls []struct {
			Indices     []int  `json:"indices"`
			Url         string `json:"url"`
			ExpandedUrl string `json:"expanded_url"`
			DisplayUrl  string `json:"display_url"`
		} `json:"urls"`
	} `json:"entities"`

	ExtendedEntities struct {
		Media []struct {
			DisplayUrl    string `json:"display_url"`
			ExpandedUrl   string `json:"expanded_url"`
			MediaUrlHttps string `json:"media_url_https"`
			Type          string `json:"type"`
			MediaKey      string `json:"media_key"`
			VideoInfo     struct {
				Variants []Variant `json:"variants"`
			} `json:"video_info"`
		} `json:"media"`
	} `json:"extended_entities"`
	FavoriteCount int    `json:"favorite_count"`
	Favorited     bool   `json:"favorited"`
	FullText      string `json:"full_text"`
	IsQuoteStatus bool   `json:"is_quote_status"`
	Lang          string `json:"lang"`

	QuoteCount   int    `json:"quote_count"`
	ReplyCount   int    `json:"reply_count"`
	RetweetCount int    `json:"retweet_count"`
	Retweeted    bool   `json:"retweeted"`
	UserIdStr    string `json:"user_id_str"`
}

type UserData struct {
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
}

type NoteResult struct {
}

type BirdwatchPivot struct {
	Note struct {
		Summary struct {
			Text string `json:"text"`
		} `json:"summary"`
	} `json:"note"`
	Shorttitle string `json:"shorttitle"`
}

type Tweet struct {
	Tweet *Tweet `json:"tweet"`
	Core  *struct {
		UserResult struct {
			Result struct {
				Legacy UserData `json:"legacy"`
			} `json:"result"`
		} `json:"user_result"`
	} `json:"core"`
	Legacy *TweetData `json:"legacy"`

	BirdwatchPivot *BirdwatchPivot `json:"birdwatch_pivot"`

	NoteTweet *struct {
		NoteTweetResults struct {
			Result struct {
				Text string `json:"text"`
			} `json:"result"`
		} `json:"note_tweet_results"`
	} `json:"note_tweet"`

	ViewCountInfo *struct {
		Count string `json:"count"`
		State string `json:"state"`
	} `json:"view_count_info"`
}

type ParsedTweet struct {
	Data struct {
		TweetResult struct {
			Result *Tweet `json:"result"`
		} `json:"tweet_result"`
	} `json:"data"`
	Errors []struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
	Error *string `json:"error"`
}

type Entity struct {
	Content struct {
		Content struct {
			TweetDisplayType string `json:"tweetDisplayType,omitempty"`
			TweetResult      struct {
				Result *Tweet `json:"result"`
			} `json:"tweetResult,omitempty"`
			CursorType string `json:"cursorType,omitempty"`
			Value      string `json:"value,omitempty"`
		} `json:"content,omitempty"`
		Items []struct {
			EntryId string `json:"entryId"`
			Item    struct {
				Content struct {
					Typename         string `json:"__typename"`
					TweetDisplayType string `json:"tweetDisplayType,omitempty"`
					TweetResult      struct {
						Result *Tweet `json:"result"`
					} `json:"tweetResult,omitempty"`
				} `json:"content"`
			} `json:"item"`
		} `json:"items,omitempty"`
		ModuleDisplayType string `json:"moduleDisplayType,omitempty"`
	} `json:"content"`
	EntryId   string `json:"entryId"`
	SortIndex string `json:"sortIndex"`
}

type ParsedThread struct {
	Data struct {
		TimelineResponse struct {
			Instructions []struct {
				Entries   []Entity `json:"entries,omitempty"`
				Direction string   `json:"direction,omitempty"`
			} `json:"instructions"`
			Metadata struct {
				ReaderModeConfig struct {
					IsReaderModeAvailable bool `json:"is_reader_mode_available"`
				} `json:"readerModeConfig"`
			} `json:"metadata"`
		} `json:"timeline_response"`
	} `json:"data"`
	Errors []struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
	Error *string `json:"error"`
}

type T struct {
	Data struct {
		TimelineResponse struct {
			Instructions []struct {
				Typename string `json:"__typename"`
				Entries  []struct {
					Content struct {
						Typename string `json:"__typename"`
						Content  struct {
							Typename            string `json:"__typename"`
							HasModeratedReplies bool   `json:"hasModeratedReplies,omitempty"`
							TweetDisplayType    string `json:"tweetDisplayType,omitempty"`
							TweetResult         struct {
								Result struct {
									Typename       string `json:"__typename"`
									BirdwatchPivot struct {
										CallToAction struct {
											DestinationUrl string `json:"destination_url"`
											Prompt         string `json:"prompt"`
											Title          string `json:"title"`
										} `json:"call_to_action"`
										DestinationUrl string `json:"destination_url"`
										IconType       string `json:"icon_type"`
										Note           struct {
											RestId  string `json:"rest_id"`
											Summary struct {
												Entities []struct {
													FromIndex int `json:"fromIndex"`
													Ref       struct {
														Typename string `json:"__typename"`
														Url      string `json:"url"`
														UrlType  string `json:"urlType"`
													} `json:"ref"`
													ToIndex int `json:"toIndex"`
												} `json:"entities"`
												Text string `json:"text"`
											} `json:"summary"`
										} `json:"note"`
										Shorttitle string `json:"shorttitle"`
										Subtitle   struct {
											Entities []struct {
												FromIndex int `json:"fromIndex"`
												Ref       struct {
													Typename string `json:"__typename"`
													Url      string `json:"url"`
													UrlType  string `json:"urlType"`
												} `json:"ref"`
												ToIndex int `json:"toIndex"`
											} `json:"entities"`
											Text string `json:"text"`
										} `json:"subtitle"`
										Title string `json:"title"`
									} `json:"birdwatch_pivot"`
									ConversationMuted bool `json:"conversation_muted"`
									Core              struct {
										UserResult struct {
											Result struct {
												Typename                   string `json:"__typename"`
												AffiliatesHighlightedLabel struct {
												} `json:"affiliates_highlighted_label"`
												BusinessAccount struct {
												} `json:"business_account"`
												CreatorSubscriptionsCount int  `json:"creator_subscriptions_count"`
												ExclusiveTweetFollowing   bool `json:"exclusive_tweet_following"`
												HasNftAvatar              bool `json:"has_nft_avatar"`
												IsBlueVerified            bool `json:"is_blue_verified"`
												Legacy                    struct {
													AdvertiserAccountServiceLevels []interface{} `json:"advertiser_account_service_levels"`
													AdvertiserAccountType          string        `json:"advertiser_account_type"`
													AnalyticsType                  string        `json:"analytics_type"`
													CanDm                          bool          `json:"can_dm"`
													CanMediaTag                    bool          `json:"can_media_tag"`
													CreatedAt                      string        `json:"created_at"`
													Description                    string        `json:"description"`
													Entities                       struct {
														Description struct {
															Hashtags     []interface{} `json:"hashtags"`
															Symbols      []interface{} `json:"symbols"`
															Urls         []interface{} `json:"urls"`
															UserMentions []interface{} `json:"user_mentions"`
														} `json:"description"`
														Url struct {
															Urls []struct {
																DisplayUrl  string `json:"display_url"`
																ExpandedUrl string `json:"expanded_url"`
																Indices     []int  `json:"indices"`
																Url         string `json:"url"`
															} `json:"urls"`
														} `json:"url"`
													} `json:"entities"`
													FastFollowersCount      int           `json:"fast_followers_count"`
													FavouritesCount         int           `json:"favourites_count"`
													FollowersCount          int           `json:"followers_count"`
													FriendsCount            int           `json:"friends_count"`
													GeoEnabled              bool          `json:"geo_enabled"`
													HasCustomTimelines      bool          `json:"has_custom_timelines"`
													HasExtendedProfile      bool          `json:"has_extended_profile"`
													IdStr                   string        `json:"id_str"`
													IsTranslator            bool          `json:"is_translator"`
													Location                string        `json:"location"`
													MediaCount              int           `json:"media_count"`
													Name                    string        `json:"name"`
													NormalFollowersCount    int           `json:"normal_followers_count"`
													PinnedTweetIdsStr       []string      `json:"pinned_tweet_ids_str"`
													ProfileBackgroundColor  string        `json:"profile_background_color"`
													ProfileBannerUrl        string        `json:"profile_banner_url"`
													ProfileImageUrlHttps    string        `json:"profile_image_url_https"`
													ProfileInterstitialType string        `json:"profile_interstitial_type"`
													ProfileLinkColor        string        `json:"profile_link_color"`
													Protected               bool          `json:"protected"`
													ScreenName              string        `json:"screen_name"`
													StatusesCount           int           `json:"statuses_count"`
													TranslatorTypeEnum      string        `json:"translator_type_enum"`
													Url                     string        `json:"url"`
													Verified                bool          `json:"verified"`
													WithheldInCountries     []interface{} `json:"withheld_in_countries"`
												} `json:"legacy"`
												PrivateSuperFollowing bool `json:"private_super_following"`
												Professional          struct {
													Category []struct {
														Id   int    `json:"id"`
														Name string `json:"name"`
													} `json:"category"`
													ProfessionalType        string `json:"professional_type"`
													QuickPromoteEligibility struct {
														IsEligible bool `json:"is_eligible"`
													} `json:"quick_promote_eligibility"`
												} `json:"professional"`
												ProfileImageShape   string `json:"profile_image_shape"`
												RestId              string `json:"rest_id"`
												SuperFollowEligible bool   `json:"super_follow_eligible"`
												SuperFollowedBy     bool   `json:"super_followed_by"`
												SuperFollowing      bool   `json:"super_following"`
											} `json:"result"`
										} `json:"user_result"`
									} `json:"core"`
									EditControl struct {
										Typename           string   `json:"__typename"`
										EditTweetIds       []string `json:"edit_tweet_ids"`
										EditableUntilMsecs string   `json:"editable_until_msecs"`
										EditsRemaining     string   `json:"edits_remaining"`
										IsEditEligible     bool     `json:"is_edit_eligible"`
									} `json:"edit_control"`
									IsTranslatable bool `json:"is_translatable"`
									Legacy         struct {
										BookmarkCount     int    `json:"bookmark_count"`
										Bookmarked        bool   `json:"bookmarked"`
										ConversationIdStr string `json:"conversation_id_str"`
										CreatedAt         string `json:"created_at"`
										DisplayTextRange  []int  `json:"display_text_range"`
										Entities          struct {
											Hashtags     []interface{} `json:"hashtags"`
											Symbols      []interface{} `json:"symbols"`
											Urls         []interface{} `json:"urls"`
											UserMentions []interface{} `json:"user_mentions"`
										} `json:"entities"`
										FavoriteCount int    `json:"favorite_count"`
										Favorited     bool   `json:"favorited"`
										FullText      string `json:"full_text"`
										IsQuoteStatus bool   `json:"is_quote_status"`
										Lang          string `json:"lang"`
										QuoteCount    int    `json:"quote_count"`
										ReplyCount    int    `json:"reply_count"`
										RetweetCount  int    `json:"retweet_count"`
										Retweeted     bool   `json:"retweeted"`
										UserIdStr     string `json:"user_id_str"`
									} `json:"legacy"`
									QuickPromoteEligibility struct {
										Eligibility string `json:"eligibility"`
									} `json:"quick_promote_eligibility"`
									RestId        string `json:"rest_id"`
									UnmentionData struct {
									} `json:"unmention_data"`
									ViewCountInfo struct {
										Count string `json:"count"`
										State string `json:"state"`
									} `json:"view_count_info"`
								} `json:"result"`
							} `json:"tweetResult,omitempty"`
							CursorType string `json:"cursorType,omitempty"`
							Value      string `json:"value,omitempty"`
						} `json:"content,omitempty"`
						ClientEventInfo struct {
							Details struct {
								ConversationDetails struct {
									ConversationSection string `json:"conversationSection"`
								} `json:"conversationDetails"`
							} `json:"details"`
						} `json:"clientEventInfo,omitempty"`
						Items []struct {
							EntryId string `json:"entryId"`
							Item    struct {
								ClientEventInfo struct {
									Details struct {
										ConversationDetails struct {
											ConversationSection string `json:"conversationSection"`
										} `json:"conversationDetails"`
										TimelinesDetails struct {
											ControllerData string `json:"controllerData"`
										} `json:"timelinesDetails"`
									} `json:"details"`
								} `json:"clientEventInfo"`
								Content struct {
									Typename         string `json:"__typename"`
									TweetDisplayType string `json:"tweetDisplayType"`
									TweetResult      struct {
										Result struct {
											Typename          string `json:"__typename"`
											ConversationMuted bool   `json:"conversation_muted"`
											Core              struct {
												UserResult struct {
													Result struct {
														Typename                   string `json:"__typename"`
														AffiliatesHighlightedLabel struct {
														} `json:"affiliates_highlighted_label"`
														BusinessAccount struct {
														} `json:"business_account"`
														CreatorSubscriptionsCount       int  `json:"creator_subscriptions_count"`
														ExclusiveTweetFollowing         bool `json:"exclusive_tweet_following"`
														HasHiddenLikesOnProfile         bool `json:"has_hidden_likes_on_profile,omitempty"`
														HasHiddenSubscriptionsOnProfile bool `json:"has_hidden_subscriptions_on_profile,omitempty"`
														HasNftAvatar                    bool `json:"has_nft_avatar"`
														IsBlueVerified                  bool `json:"is_blue_verified"`
														Legacy                          struct {
															AdvertiserAccountServiceLevels []string `json:"advertiser_account_service_levels"`
															AdvertiserAccountType          string   `json:"advertiser_account_type"`
															AnalyticsType                  string   `json:"analytics_type"`
															CanDm                          bool     `json:"can_dm"`
															CanMediaTag                    bool     `json:"can_media_tag"`
															CreatedAt                      string   `json:"created_at"`
															Description                    string   `json:"description"`
															Entities                       struct {
																Description struct {
																	Hashtags []interface{} `json:"hashtags"`
																	Symbols  []interface{} `json:"symbols"`
																	Urls     []struct {
																		DisplayUrl  string `json:"display_url"`
																		ExpandedUrl string `json:"expanded_url"`
																		Indices     []int  `json:"indices"`
																		Url         string `json:"url"`
																	} `json:"urls"`
																	UserMentions []struct {
																		IdStr      string `json:"id_str"`
																		Indices    []int  `json:"indices"`
																		Name       string `json:"name"`
																		ScreenName string `json:"screen_name"`
																	} `json:"user_mentions"`
																} `json:"description"`
																Url struct {
																	Urls []struct {
																		DisplayUrl  string `json:"display_url"`
																		ExpandedUrl string `json:"expanded_url"`
																		Indices     []int  `json:"indices"`
																		Url         string `json:"url"`
																	} `json:"urls"`
																} `json:"url,omitempty"`
															} `json:"entities"`
															FastFollowersCount      int           `json:"fast_followers_count"`
															FavouritesCount         int           `json:"favourites_count"`
															FollowersCount          int           `json:"followers_count"`
															FriendsCount            int           `json:"friends_count"`
															GeoEnabled              bool          `json:"geo_enabled"`
															HasCustomTimelines      bool          `json:"has_custom_timelines"`
															HasExtendedProfile      bool          `json:"has_extended_profile"`
															IdStr                   string        `json:"id_str"`
															IsTranslator            bool          `json:"is_translator"`
															Location                string        `json:"location"`
															MediaCount              int           `json:"media_count"`
															Name                    string        `json:"name"`
															NormalFollowersCount    int           `json:"normal_followers_count"`
															PinnedTweetIdsStr       []string      `json:"pinned_tweet_ids_str"`
															ProfileBackgroundColor  string        `json:"profile_background_color"`
															ProfileBannerUrl        string        `json:"profile_banner_url,omitempty"`
															ProfileImageUrlHttps    string        `json:"profile_image_url_https"`
															ProfileInterstitialType string        `json:"profile_interstitial_type"`
															ProfileLinkColor        string        `json:"profile_link_color"`
															Protected               bool          `json:"protected"`
															ScreenName              string        `json:"screen_name"`
															StatusesCount           int           `json:"statuses_count"`
															TranslatorTypeEnum      string        `json:"translator_type_enum"`
															Verified                bool          `json:"verified"`
															WithheldInCountries     []interface{} `json:"withheld_in_countries"`
															ProfileLocationPlace    struct {
																Country     string `json:"country"`
																CountryCode string `json:"country_code"`
																FullName    string `json:"full_name"`
																Id          string `json:"id"`
																Name        string `json:"name"`
																PlaceType   string `json:"place_type"`
															} `json:"profile_location_place,omitempty"`
															Url string `json:"url,omitempty"`
														} `json:"legacy"`
														PrivateSuperFollowing bool `json:"private_super_following"`
														Professional          struct {
															Category []struct {
																Id   int    `json:"id"`
																Name string `json:"name"`
															} `json:"category"`
															ProfessionalType        string `json:"professional_type"`
															QuickPromoteEligibility struct {
																IsEligible bool `json:"is_eligible"`
															} `json:"quick_promote_eligibility"`
														} `json:"professional,omitempty"`
														ProfileImageShape   string `json:"profile_image_shape"`
														RestId              string `json:"rest_id"`
														SuperFollowEligible bool   `json:"super_follow_eligible"`
														SuperFollowedBy     bool   `json:"super_followed_by"`
														SuperFollowing      bool   `json:"super_following"`
													} `json:"result"`
												} `json:"user_result"`
											} `json:"core"`
											EditControl struct {
												Typename           string   `json:"__typename"`
												EditTweetIds       []string `json:"edit_tweet_ids"`
												EditableUntilMsecs string   `json:"editable_until_msecs"`
												EditsRemaining     string   `json:"edits_remaining"`
												IsEditEligible     bool     `json:"is_edit_eligible"`
											} `json:"edit_control"`
											IsTranslatable bool `json:"is_translatable"`
											Legacy         struct {
												BookmarkCount     int    `json:"bookmark_count"`
												Bookmarked        bool   `json:"bookmarked"`
												ConversationIdStr string `json:"conversation_id_str"`
												CreatedAt         string `json:"created_at"`
												DisplayTextRange  []int  `json:"display_text_range"`
												Entities          struct {
													Hashtags     []interface{} `json:"hashtags"`
													Symbols      []interface{} `json:"symbols"`
													Urls         []interface{} `json:"urls"`
													UserMentions []struct {
														IdStr      string `json:"id_str"`
														Indices    []int  `json:"indices"`
														Name       string `json:"name"`
														ScreenName string `json:"screen_name"`
													} `json:"user_mentions"`
												} `json:"entities"`
												FavoriteCount        int    `json:"favorite_count"`
												Favorited            bool   `json:"favorited"`
												FullText             string `json:"full_text"`
												InReplyToScreenName  string `json:"in_reply_to_screen_name"`
												InReplyToStatusIdStr string `json:"in_reply_to_status_id_str"`
												InReplyToUserIdStr   string `json:"in_reply_to_user_id_str"`
												IsQuoteStatus        bool   `json:"is_quote_status"`
												Lang                 string `json:"lang"`
												QuoteCount           int    `json:"quote_count"`
												ReplyCount           int    `json:"reply_count"`
												RetweetCount         int    `json:"retweet_count"`
												Retweeted            bool   `json:"retweeted"`
												UserIdStr            string `json:"user_id_str"`
											} `json:"legacy"`
											QuickPromoteEligibility struct {
												Eligibility string `json:"eligibility"`
											} `json:"quick_promote_eligibility"`
											RestId        string `json:"rest_id"`
											UnmentionData struct {
											} `json:"unmention_data"`
											ViewCountInfo struct {
												Count string `json:"count"`
												State string `json:"state"`
											} `json:"view_count_info"`
										} `json:"result"`
									} `json:"tweetResult"`
								} `json:"content"`
							} `json:"item"`
						} `json:"items,omitempty"`
						ModuleDisplayType string `json:"moduleDisplayType,omitempty"`
					} `json:"content"`
					EntryId   string `json:"entryId"`
					SortIndex string `json:"sortIndex"`
				} `json:"entries,omitempty"`
				Direction string `json:"direction,omitempty"`
			} `json:"instructions"`
		} `json:"timeline_response"`
	} `json:"data"`
}
