package socialapi1instagram

type ParsedPost struct {
	Data struct {
		AccessibilityCaption                  interface{}     `json:"accessibility_caption"`
		Caption                               *Caption        `json:"caption"`
		CaptionIsEdited                       bool            `json:"caption_is_edited"`
		CarouselMedia                         []CarouselMedia `json:"carousel_media"`
		CarouselMediaCount                    int             `json:"carousel_media_count"`
		CarouselMediaIds                      []int64         `json:"carousel_media_ids"`
		CarouselMediaPendingPostCount         int             `json:"carousel_media_pending_post_count"`
		ClipsTabPinnedUserIds                 []interface{}   `json:"clips_tab_pinned_user_ids"`
		CoauthorProducerCanSeeOrganicInsights bool            `json:"coauthor_producer_can_see_organic_insights"`
		CoauthorProducers                     []interface{}   `json:"coauthor_producers"`
		Code                                  string          `json:"code"`

		FeaturedProducts []interface{} `json:"featured_products"`
		FilterType       int           `json:"filter_type"`
		FundraiserTag    struct {
			HasStandaloneFundraiser bool `json:"has_standalone_fundraiser"`
		} `json:"fundraiser_tag"`
		GenAiDetectionMethod struct {
			DetectionMethod string `json:"detection_method"`
		} `json:"gen_ai_detection_method"`
		HasHighRiskGenAiInformTreatment        bool          `json:"has_high_risk_gen_ai_inform_treatment"`
		HasLiked                               bool          `json:"has_liked"`
		HasMoreComments                        bool          `json:"has_more_comments"`
		HasPrivatelyLiked                      bool          `json:"has_privately_liked"`
		HasSharedToFb                          int           `json:"has_shared_to_fb"`
		HasViewsFetching                       bool          `json:"has_views_fetching"`
		Id                                     string        `json:"id"`
		IgMediaSharingDisabled                 bool          `json:"ig_media_sharing_disabled"`
		IgbioProduct                           interface{}   `json:"igbio_product"`
		ImageVersions                          *ImageVersion `json:"image_versions"`
		InlineComposerDisplayCondition         string        `json:"inline_composer_display_condition"`
		InlineComposerImpTriggerTime           int           `json:"inline_composer_imp_trigger_time"`
		IntegrityReviewDecision                string        `json:"integrity_review_decision"`
		InvitedCoauthorProducers               []interface{} `json:"invited_coauthor_producers"`
		IsCommentsGifComposerEnabled           bool          `json:"is_comments_gif_composer_enabled"`
		IsCutoutStickerAllowed                 bool          `json:"is_cutout_sticker_allowed"`
		IsEligibleForMediaNoteRecsNux          bool          `json:"is_eligible_for_media_note_recs_nux"`
		IsInProfileGrid                        bool          `json:"is_in_profile_grid"`
		IsOpenToPublicSubmission               bool          `json:"is_open_to_public_submission"`
		IsOrganicProductTaggingEligible        bool          `json:"is_organic_product_tagging_eligible"`
		IsPaidPartnership                      bool          `json:"is_paid_partnership"`
		IsPinned                               bool          `json:"is_pinned"`
		IsPostLiveClipsMedia                   bool          `json:"is_post_live_clips_media"`
		IsQuietPost                            bool          `json:"is_quiet_post"`
		IsReshareOfTextPostAppMediaInIg        bool          `json:"is_reshare_of_text_post_app_media_in_ig"`
		IsReuseAllowed                         bool          `json:"is_reuse_allowed"`
		IsSocialUfiDisabled                    bool          `json:"is_social_ufi_disabled"`
		IsTaggedMediaSharedToViewerProfileGrid bool          `json:"is_tagged_media_shared_to_viewer_profile_grid"`
		IsUnifiedVideo                         bool          `json:"is_unified_video"`
		IsVideo                                bool          `json:"is_video"`
		Lat                                    float64       `json:"lat"`
		LikeAndViewCountsDisabled              bool          `json:"like_and_view_counts_disabled"`
		Lng                                    float64       `json:"lng"`
		Location                               struct {
			Address             string  `json:"address"`
			City                string  `json:"city"`
			ExternalSource      string  `json:"external_source"`
			FacebookPlacesId    int64   `json:"facebook_places_id"`
			IsEligibleForGuides bool    `json:"is_eligible_for_guides"`
			Lat                 float64 `json:"lat"`
			Lng                 float64 `json:"lng"`
			Name                string  `json:"name"`
			Pk                  int64   `json:"pk"`
			ShortName           string  `json:"short_name"`
		} `json:"location"`
		MaxNumVisiblePreviewComments int    `json:"max_num_visible_preview_comments"`
		MediaName                    string `json:"media_name"`
		MediaNotes                   struct {
			Items []interface{} `json:"items"`
		} `json:"media_notes"`
		MediaType              int           `json:"media_type"`
		MetaAiSuggestedPrompts []interface{} `json:"meta_ai_suggested_prompts"`
		Metrics                *Metrics      `json:"metrics"`
		MusicMetadata          struct {
			AudioCanonicalId  string      `json:"audio_canonical_id"`
			AudioType         interface{} `json:"audio_type"`
			MusicInfo         interface{} `json:"music_info"`
			OriginalSoundInfo interface{} `json:"original_sound_info"`
			PinnedMediaIds    interface{} `json:"pinned_media_ids"`
		} `json:"music_metadata"`
		OpenCarouselShowFollowButton bool          `json:"open_carousel_show_follow_button"`
		OpenCarouselSubmissionState  string        `json:"open_carousel_submission_state"`
		OriginalHeight               int           `json:"original_height"`
		OriginalWidth                int           `json:"original_width"`
		Owner                        *User         `json:"owner"`
		Pk                           int64         `json:"pk"`
		PreviewComments              []interface{} `json:"preview_comments"`
		ProductSuggestions           []interface{} `json:"product_suggestions"`
		ProductType                  string        `json:"product_type"`
		ShareCountDisabled           bool          `json:"share_count_disabled"`

		ShopRoutingUserId                                    interface{}     `json:"shop_routing_user_id"`
		ShouldShowAuthorPogForTaggedMediaSharedToProfileGrid bool            `json:"should_show_author_pog_for_tagged_media_shared_to_profile_grid"`
		SubscribeCtaVisible                                  bool            `json:"subscribe_cta_visible"`
		TaggedUsers                                          interface{}     `json:"tagged_users"`
		TakenAt                                              int             `json:"taken_at"`
		ThumbnailUrl                                         string          `json:"thumbnail_url"`
		TimelinePinnedUserIds                                []interface{}   `json:"timeline_pinned_user_ids"`
		TopLikers                                            []string        `json:"top_likers"`
		User                                                 *User           `json:"user"`
		VideoUrl                                             string          `json:"video_url"`
		VideoVersions                                        []*VideoVersion `json:"video_versions"`
	} `json:"data"`
}

type Caption struct {
	CreatedAt          int           `json:"created_at"`
	CreatedAtUtc       int           `json:"created_at_utc"`
	DidReportAsSpam    bool          `json:"did_report_as_spam"`
	HasTranslation     bool          `json:"has_translation"`
	Hashtags           []interface{} `json:"hashtags"`
	Id                 int64         `json:"id"`
	IsCovered          bool          `json:"is_covered"`
	IsRankedComment    bool          `json:"is_ranked_comment"`
	Mentions           []interface{} `json:"mentions"`
	Pk                 string        `json:"pk"`
	PrivateReplyStatus int           `json:"private_reply_status"`
	ShareEnabled       bool          `json:"share_enabled"`
	Text               string        `json:"text"`
	Type               int           `json:"type"`
}

type User struct {
	FullName                   string `json:"full_name"`
	HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
	Id                         string `json:"id"`
	IsFavorite                 bool   `json:"is_favorite"`
	IsPrivate                  bool   `json:"is_private"`
	IsUnpublished              bool   `json:"is_unpublished"`
	IsVerified                 bool   `json:"is_verified"`
	LatestReelMedia            int    `json:"latest_reel_media"`
	ProfilePicId               string `json:"profile_pic_id"`
	ProfilePicUrl              string `json:"profile_pic_url"`
	Username                   string `json:"username"`
}

type CarouselMedia struct {
	CarouselParentId    string `json:"carousel_parent_id"`
	CommercialityStatus string `json:"commerciality_status"`
	ExplorePivotGrid    bool   `json:"explore_pivot_grid"`
	FbUserTags          struct {
		In []interface{} `json:"in"`
	} `json:"fb_user_tags"`
	FeaturedProducts    []interface{} `json:"featured_products"`
	Id                  string        `json:"id"`
	ImageVersions       *ImageVersion `json:"image_versions"`
	IsVideo             bool          `json:"is_video"`
	MediaName           string        `json:"media_name"`
	MediaType           int           `json:"media_type"`
	OriginalHeight      int           `json:"original_height"`
	OriginalWidth       int           `json:"original_width"`
	Pk                  int64         `json:"pk"`
	ProductSuggestions  []interface{} `json:"product_suggestions"`
	ProductType         string        `json:"product_type"`
	SharingFrictionInfo struct {
		BloksAppUrl               interface{} `json:"bloks_app_url"`
		SharingFrictionPayload    interface{} `json:"sharing_friction_payload"`
		ShouldHaveSharingFriction bool        `json:"should_have_sharing_friction"`
	} `json:"sharing_friction_info"`
	ShopRoutingUserId   interface{}   `json:"shop_routing_user_id"`
	TakenAt             int           `json:"taken_at"`
	ThumbnailUrl        string        `json:"thumbnail_url"`
	VideoStickerLocales []interface{} `json:"video_sticker_locales"`
}

type VideoVersion struct {
	Height int    `json:"height"`
	Id     string `json:"id"`
	Type   int    `json:"type"`
	Url    string `json:"url"`
	Width  int    `json:"width"`
}

type ImageVersion struct {
	Items []struct {
		Height int    `json:"height"`
		Url    string `json:"url"`
		Width  int    `json:"width"`
	} `json:"items"`
}

type Metrics struct {
	CommentCount      int         `json:"comment_count"`
	FbLikeCount       interface{} `json:"fb_like_count"`
	FbPlayCount       interface{} `json:"fb_play_count"`
	LikeCount         int         `json:"like_count"`
	PlayCount         interface{} `json:"play_count"`
	SaveCount         interface{} `json:"save_count"`
	ShareCount        int         `json:"share_count"`
	UserFollowerCount interface{} `json:"user_follower_count"`
	UserMediaCount    interface{} `json:"user_media_count"`
	ViewCount         int         `json:"view_count"`
}
