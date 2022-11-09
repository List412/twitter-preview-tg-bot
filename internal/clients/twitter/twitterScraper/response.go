package twitterScraper

type TweetWeb struct {
	Data struct {
		ThreadedConversationWithInjections struct {
			Instructions []struct {
				Type    string `json:"type"`
				Entries []struct {
					EntryId   string `json:"entryId"`
					SortIndex string `json:"sortIndex"`
					Content   struct {
						EntryType   string `json:"entryType"`
						Typename    string `json:"__typename"`
						ItemContent struct {
							ItemType     string `json:"itemType"`
							Typename     string `json:"__typename"`
							TweetResults struct {
								Result struct {
									Typename string `json:"__typename"`
									RestId   string `json:"rest_id"`
									Core     struct {
										UserResults struct {
											Result struct {
												Typename string `json:"__typename"`
												Id       string `json:"id"`
												RestId   string `json:"rest_id"`
												Legacy   struct {
													CreatedAt string `json:"created_at"`

													FastFollowersCount   int    `json:"fast_followers_count"`
													FavouritesCount      int    `json:"favourites_count"`
													FollowersCount       int    `json:"followers_count"`
													FriendsCount         int    `json:"friends_count"`
													HasCustomTimelines   bool   `json:"has_custom_timelines"`
													IsTranslator         bool   `json:"is_translator"`
													ListedCount          int    `json:"listed_count"`
													Location             string `json:"location"`
													MediaCount           int    `json:"media_count"`
													Name                 string `json:"name"`
													NormalFollowersCount int    `json:"normal_followers_count"`

													ScreenName string `json:"screen_name"`
												} `json:"legacy"`
											} `json:"result"`
										} `json:"user_results"`
									} `json:"core"`
									Legacy struct {
										CreatedAt         string `json:"created_at"`
										ConversationIdStr string `json:"conversation_id_str"`
										Entities          struct {
											UserMentions []struct {
												IdStr      string `json:"id_str"`
												Name       string `json:"name"`
												ScreenName string `json:"screen_name"`
												Indices    []int  `json:"indices"`
											} `json:"user_mentions"`
											Urls     []interface{} `json:"urls"`
											Hashtags []struct {
												Indices []int  `json:"indices"`
												Text    string `json:"text"`
											} `json:"hashtags"`
											Symbols []interface{} `json:"symbols"`
										} `json:"entities"`
										FavoriteCount        int    `json:"favorite_count"`
										Favorited            bool   `json:"favorited"`
										FullText             string `json:"full_text"`
										IsQuoteStatus        bool   `json:"is_quote_status"`
										Lang                 string `json:"lang"`
										QuoteCount           int    `json:"quote_count"`
										ReplyCount           int    `json:"reply_count"`
										RetweetCount         int    `json:"retweet_count"`
										Retweeted            bool   `json:"retweeted"`
										Source               string `json:"source"`
										UserIdStr            string `json:"user_id_str"`
										IdStr                string `json:"id_str"`
										InReplyToScreenName  string `json:"in_reply_to_screen_name,omitempty"`
										InReplyToStatusIdStr string `json:"in_reply_to_status_id_str,omitempty"`
										InReplyToUserIdStr   string `json:"in_reply_to_user_id_str,omitempty"`
										SelfThread           *struct {
											IdStr string `json:"id_str"`
										} `json:"self_thread"`
									} `json:"legacy"`
								} `json:"result"`
							} `json:"tweet_results,omitempty"`
						} `json:"itemContent,omitempty"`
					} `json:"content"`
				} `json:"entries,omitempty"`
			} `json:"instructions"`
		} `json:"threaded_conversation_with_injections"`
	} `json:"data"`
}

type TweetWeb2 struct {
	Data struct {
		ThreadedConversationWithInjections struct {
			Instructions []struct {
				Type    string `json:"type"`
				Entries []struct {
					EntryId string `json:"entryId"`
					Item    struct {
						ItemContent struct {
							TweetResults struct {
								Result struct {
									Legacy struct {
										CreatedAt         string `json:"created_at"`
										ConversationIdStr string `json:"conversation_id_str"`
										DisplayTextRange  []int  `json:"display_text_range"`
										Entities          struct {
											UserMentions []interface{} `json:"user_mentions"`
											Urls         []interface{} `json:"urls"`
											Hashtags     []interface{} `json:"hashtags"`
											Symbols      []interface{} `json:"symbols"`
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
										Source               string `json:"source"`
										UserIdStr            string `json:"user_id_str"`
										IdStr                string `json:"id_str"`
										SelfThread           struct {
											IdStr string `json:"id_str"`
										} `json:"self_thread"`
									} `json:"legacy"`
								} `json:"result"`
							} `json:"tweet_results"`
						} `json:"itemContent"`
					} `json:"item"`
				}
			}
		}
	}
}

type TweetHuge struct {
	Data struct {
		ThreadedConversationWithInjections struct {
			Instructions []struct {
				Type    string `json:"type"`
				Entries []struct {
					EntryId   string `json:"entryId"`
					SortIndex string `json:"sortIndex"`
					Content   struct {
						EntryType   string `json:"entryType"`
						Typename    string `json:"__typename"`
						ItemContent struct {
							ItemType     string `json:"itemType"`
							Typename     string `json:"__typename"`
							TweetResults struct {
								Result struct {
									Typename string `json:"__typename"`
									RestId   string `json:"rest_id"`
									Core     struct {
										UserResults struct {
											Result struct {
												Typename                   string `json:"__typename"`
												Id                         string `json:"id"`
												RestId                     string `json:"rest_id"`
												AffiliatesHighlightedLabel struct {
												} `json:"affiliates_highlighted_label"`
												HasNftAvatar bool `json:"has_nft_avatar"`
												Legacy       struct {
													CreatedAt           string `json:"created_at"`
													DefaultProfile      bool   `json:"default_profile"`
													DefaultProfileImage bool   `json:"default_profile_image"`
													Description         string `json:"description"`
													Entities            struct {
														Description struct {
															Urls []interface{} `json:"urls"`
														} `json:"description"`
														Url struct {
															Urls []struct {
																DisplayUrl  string `json:"display_url"`
																ExpandedUrl string `json:"expanded_url"`
																Url         string `json:"url"`
																Indices     []int  `json:"indices"`
															} `json:"urls"`
														} `json:"url"`
													} `json:"entities"`
													FastFollowersCount      int      `json:"fast_followers_count"`
													FavouritesCount         int      `json:"favourites_count"`
													FollowersCount          int      `json:"followers_count"`
													FriendsCount            int      `json:"friends_count"`
													HasCustomTimelines      bool     `json:"has_custom_timelines"`
													IsTranslator            bool     `json:"is_translator"`
													ListedCount             int      `json:"listed_count"`
													Location                string   `json:"location"`
													MediaCount              int      `json:"media_count"`
													Name                    string   `json:"name"`
													NormalFollowersCount    int      `json:"normal_followers_count"`
													PinnedTweetIdsStr       []string `json:"pinned_tweet_ids_str"`
													PossiblySensitive       bool     `json:"possibly_sensitive"`
													ProfileBannerExtensions struct {
														MediaColor struct {
															R struct {
																Ok struct {
																	Palette []struct {
																		Percentage float64 `json:"percentage"`
																		Rgb        struct {
																			Blue  int `json:"blue"`
																			Green int `json:"green"`
																			Red   int `json:"red"`
																		} `json:"rgb"`
																	} `json:"palette"`
																} `json:"ok"`
															} `json:"r"`
														} `json:"mediaColor"`
													} `json:"profile_banner_extensions"`
													ProfileBannerUrl       string `json:"profile_banner_url"`
													ProfileImageExtensions struct {
														MediaColor struct {
															R struct {
																Ok struct {
																	Palette []struct {
																		Percentage float64 `json:"percentage"`
																		Rgb        struct {
																			Blue  int `json:"blue"`
																			Green int `json:"green"`
																			Red   int `json:"red"`
																		} `json:"rgb"`
																	} `json:"palette"`
																} `json:"ok"`
															} `json:"r"`
														} `json:"mediaColor"`
													} `json:"profile_image_extensions"`
													ProfileImageUrlHttps    string        `json:"profile_image_url_https"`
													ProfileInterstitialType string        `json:"profile_interstitial_type"`
													Protected               bool          `json:"protected"`
													ScreenName              string        `json:"screen_name"`
													StatusesCount           int           `json:"statuses_count"`
													TranslatorType          string        `json:"translator_type"`
													Url                     string        `json:"url"`
													Verified                bool          `json:"verified"`
													WithheldInCountries     []interface{} `json:"withheld_in_countries"`
												} `json:"legacy"`
												Professional struct {
													RestId           string `json:"rest_id"`
													ProfessionalType string `json:"professional_type"`
													Category         []struct {
														Id       int    `json:"id"`
														Name     string `json:"name"`
														IconName string `json:"icon_name"`
													} `json:"category"`
												} `json:"professional"`
											} `json:"result"`
										} `json:"user_results"`
									} `json:"core"`
									UnmentionData struct {
									} `json:"unmention_data"`
									EditControl struct {
										EditTweetIds       []string `json:"edit_tweet_ids"`
										EditableUntilMsecs string   `json:"editable_until_msecs"`
										IsEditEligible     bool     `json:"is_edit_eligible"`
										EditsRemaining     string   `json:"edits_remaining"`
									} `json:"edit_control"`
									Legacy struct {
										CollabControl     *CollabControl `json:"collab_control"`
										CreatedAt         string         `json:"created_at"`
										ConversationIdStr string         `json:"conversation_id_str"`
										DisplayTextRange  []int          `json:"display_text_range"`
										Entities          struct {
											UserMentions []interface{} `json:"user_mentions"`
											Urls         []interface{} `json:"urls"`
											Hashtags     []interface{} `json:"hashtags"`
											Symbols      []interface{} `json:"symbols"`
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
										Source        string `json:"source"`
										UserIdStr     string `json:"user_id_str"`
										IdStr         string `json:"id_str"`
										SelfThread    struct {
											IdStr string `json:"id_str"`
										} `json:"self_thread"`
									} `json:"legacy"`
								} `json:"result"`
							} `json:"tweet_results,omitempty"`
							TweetDisplayType    string `json:"tweetDisplayType,omitempty"`
							HasModeratedReplies bool   `json:"hasModeratedReplies,omitempty"`
							Value               string `json:"value,omitempty"`
							CursorType          string `json:"cursorType,omitempty"`
						} `json:"itemContent,omitempty"`
						Items []struct {
							EntryId string `json:"entryId"`
							Item    struct {
								ItemContent struct {
									ItemType     string `json:"itemType"`
									Typename     string `json:"__typename"`
									TweetResults struct {
										Result struct {
											Typename string `json:"__typename"`
											RestId   string `json:"rest_id"`
											Core     struct {
												UserResults struct {
													Result struct {
														Typename                   string `json:"__typename"`
														Id                         string `json:"id"`
														RestId                     string `json:"rest_id"`
														AffiliatesHighlightedLabel struct {
														} `json:"affiliates_highlighted_label"`
														HasNftAvatar bool `json:"has_nft_avatar"`
														Legacy       struct {
															CreatedAt           string `json:"created_at"`
															DefaultProfile      bool   `json:"default_profile"`
															DefaultProfileImage bool   `json:"default_profile_image"`
															Description         string `json:"description"`
															Entities            struct {
																Description struct {
																	Urls []interface{} `json:"urls"`
																} `json:"description"`
																Url struct {
																	Urls []struct {
																		DisplayUrl  string `json:"display_url"`
																		ExpandedUrl string `json:"expanded_url"`
																		Url         string `json:"url"`
																		Indices     []int  `json:"indices"`
																	} `json:"urls"`
																} `json:"url,omitempty"`
															} `json:"entities"`
															FastFollowersCount      int      `json:"fast_followers_count"`
															FavouritesCount         int      `json:"favourites_count"`
															FollowersCount          int      `json:"followers_count"`
															FriendsCount            int      `json:"friends_count"`
															HasCustomTimelines      bool     `json:"has_custom_timelines"`
															IsTranslator            bool     `json:"is_translator"`
															ListedCount             int      `json:"listed_count"`
															Location                string   `json:"location"`
															MediaCount              int      `json:"media_count"`
															Name                    string   `json:"name"`
															NormalFollowersCount    int      `json:"normal_followers_count"`
															PinnedTweetIdsStr       []string `json:"pinned_tweet_ids_str"`
															PossiblySensitive       bool     `json:"possibly_sensitive"`
															ProfileBannerExtensions struct {
																MediaColor struct {
																	R struct {
																		Ok struct {
																			Palette []struct {
																				Percentage float64 `json:"percentage"`
																				Rgb        struct {
																					Blue  int `json:"blue"`
																					Green int `json:"green"`
																					Red   int `json:"red"`
																				} `json:"rgb"`
																			} `json:"palette"`
																		} `json:"ok"`
																	} `json:"r"`
																} `json:"mediaColor"`
															} `json:"profile_banner_extensions,omitempty"`
															ProfileBannerUrl       string `json:"profile_banner_url,omitempty"`
															ProfileImageExtensions struct {
																MediaColor struct {
																	R struct {
																		Ok struct {
																			Palette []struct {
																				Percentage float64 `json:"percentage"`
																				Rgb        struct {
																					Blue  int `json:"blue"`
																					Green int `json:"green"`
																					Red   int `json:"red"`
																				} `json:"rgb"`
																			} `json:"palette"`
																		} `json:"ok"`
																	} `json:"r"`
																} `json:"mediaColor"`
															} `json:"profile_image_extensions"`
															ProfileImageUrlHttps    string        `json:"profile_image_url_https"`
															ProfileInterstitialType string        `json:"profile_interstitial_type"`
															Protected               bool          `json:"protected"`
															ScreenName              string        `json:"screen_name"`
															StatusesCount           int           `json:"statuses_count"`
															TranslatorType          string        `json:"translator_type"`
															Url                     string        `json:"url,omitempty"`
															Verified                bool          `json:"verified"`
															WithheldInCountries     []interface{} `json:"withheld_in_countries"`
														} `json:"legacy"`
														Professional struct {
															RestId           string `json:"rest_id"`
															ProfessionalType string `json:"professional_type"`
															Category         []struct {
																Id       int    `json:"id"`
																Name     string `json:"name"`
																IconName string `json:"icon_name"`
															} `json:"category"`
														} `json:"professional,omitempty"`
													} `json:"result"`
												} `json:"user_results"`
											} `json:"core"`
											UnmentionData struct {
											} `json:"unmention_data"`
											EditControl struct {
												EditTweetIds       []string `json:"edit_tweet_ids"`
												EditableUntilMsecs string   `json:"editable_until_msecs"`
												IsEditEligible     bool     `json:"is_edit_eligible"`
												EditsRemaining     string   `json:"edits_remaining"`
											} `json:"edit_control"`
											Legacy struct {
												CollabControl     *CollabControl `json:"collab_control"`
												CreatedAt         string         `json:"created_at"`
												ConversationIdStr string         `json:"conversation_id_str"`
												DisplayTextRange  []int          `json:"display_text_range"`
												Entities          struct {
													UserMentions []struct {
														IdStr      string `json:"id_str"`
														Name       string `json:"name"`
														ScreenName string `json:"screen_name"`
														Indices    []int  `json:"indices"`
													} `json:"user_mentions"`
													Urls     []interface{} `json:"urls"`
													Hashtags []interface{} `json:"hashtags"`
													Symbols  []interface{} `json:"symbols"`
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
												Source               string `json:"source"`
												UserIdStr            string `json:"user_id_str"`
												IdStr                string `json:"id_str"`
												SelfThread           struct {
													IdStr string `json:"id_str"`
												} `json:"self_thread,omitempty"`
											} `json:"legacy"`
										} `json:"result"`
									} `json:"tweet_results"`
									TweetDisplayType string `json:"tweetDisplayType"`
								} `json:"itemContent"`
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
							} `json:"item"`
						} `json:"items,omitempty"`
						DisplayType     string `json:"displayType,omitempty"`
						ClientEventInfo struct {
							Details struct {
								ConversationDetails struct {
									ConversationSection string `json:"conversationSection"`
								} `json:"conversationDetails"`
							} `json:"details"`
						} `json:"clientEventInfo,omitempty"`
					} `json:"content"`
				} `json:"entries,omitempty"`
				Direction string `json:"direction,omitempty"`
			} `json:"instructions"`
			Metadata struct {
				ReaderModeConfig struct {
					IsReaderModeAvailable bool `json:"is_reader_mode_available"`
				} `json:"reader_mode_config"`
			} `json:"metadata"`
		} `json:"threaded_conversation_with_injections"`
	} `json:"data"`
}

type CollabControl struct {
	CollaboratorsResults []struct {
		Result struct {
			Typename                   string `json:"__typename"`
			Id                         string `json:"id"`
			RestId                     string `json:"rest_id"`
			AffiliatesHighlightedLabel struct {
			} `json:"affiliates_highlighted_label"`
			HasNftAvatar bool `json:"has_nft_avatar"`
			Legacy       struct {
				CreatedAt           string `json:"created_at"`
				DefaultProfile      bool   `json:"default_profile"`
				DefaultProfileImage bool   `json:"default_profile_image"`
				Description         string `json:"description"`
				Entities            struct {
					Description struct {
						Urls []struct {
							DisplayUrl  string `json:"display_url"`
							ExpandedUrl string `json:"expanded_url"`
							Url         string `json:"url"`
							Indices     []int  `json:"indices"`
						} `json:"urls"`
					} `json:"description"`
					Url struct {
						Urls []struct {
							DisplayUrl  string `json:"display_url"`
							ExpandedUrl string `json:"expanded_url"`
							Url         string `json:"url"`
							Indices     []int  `json:"indices"`
						} `json:"urls"`
					} `json:"url,omitempty"`
				} `json:"entities"`
				FastFollowersCount      int      `json:"fast_followers_count"`
				FavouritesCount         int      `json:"favourites_count"`
				FollowersCount          int      `json:"followers_count"`
				FriendsCount            int      `json:"friends_count"`
				HasCustomTimelines      bool     `json:"has_custom_timelines"`
				IsTranslator            bool     `json:"is_translator"`
				ListedCount             int      `json:"listed_count"`
				Location                string   `json:"location"`
				MediaCount              int      `json:"media_count"`
				Name                    string   `json:"name"`
				NormalFollowersCount    int      `json:"normal_followers_count"`
				PinnedTweetIdsStr       []string `json:"pinned_tweet_ids_str"`
				PossiblySensitive       bool     `json:"possibly_sensitive"`
				ProfileBannerExtensions struct {
					MediaColor struct {
						R struct {
							Ok struct {
								Palette []struct {
									Percentage float64 `json:"percentage"`
									Rgb        struct {
										Blue  int `json:"blue"`
										Green int `json:"green"`
										Red   int `json:"red"`
									} `json:"rgb"`
								} `json:"palette"`
							} `json:"ok"`
						} `json:"r"`
					} `json:"mediaColor"`
				} `json:"profile_banner_extensions"`
				ProfileBannerUrl       string `json:"profile_banner_url"`
				ProfileImageExtensions struct {
					MediaColor struct {
						R struct {
							Ok struct {
								Palette []struct {
									Percentage float64 `json:"percentage"`
									Rgb        struct {
										Blue  int `json:"blue"`
										Green int `json:"green"`
										Red   int `json:"red"`
									} `json:"rgb"`
								} `json:"palette"`
							} `json:"ok"`
						} `json:"r"`
					} `json:"mediaColor"`
				} `json:"profile_image_extensions"`
				ProfileImageUrlHttps    string        `json:"profile_image_url_https"`
				ProfileInterstitialType string        `json:"profile_interstitial_type"`
				Protected               bool          `json:"protected"`
				ScreenName              string        `json:"screen_name"`
				StatusesCount           int           `json:"statuses_count"`
				TranslatorType          string        `json:"translator_type"`
				Verified                bool          `json:"verified"`
				WithheldInCountries     []interface{} `json:"withheld_in_countries"`
				Url                     string        `json:"url,omitempty"`
			} `json:"legacy"`
		} `json:"result"`
	} `json:"collaborators_results"`
}
