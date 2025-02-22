package socialapi1instagram

type ParsedPost struct {
	Data struct {
		Caption                               *Caption        `json:"caption"`
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

		ImageVersions *ImageVersion `json:"image_versions"`

		IsVideo bool `json:"is_video"`

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
		OpenCarouselShowFollowButton bool   `json:"open_carousel_show_follow_button"`
		OpenCarouselSubmissionState  string `json:"open_carousel_submission_state"`
		OriginalHeight               int    `json:"original_height"`
		OriginalWidth                int    `json:"original_width"`
		Owner                        *User  `json:"owner"`

		ProductType string `json:"product_type"`

		TakenAt               int             `json:"taken_at"`
		ThumbnailUrl          string          `json:"thumbnail_url"`
		TimelinePinnedUserIds []interface{}   `json:"timeline_pinned_user_ids"`
		TopLikers             []string        `json:"top_likers"`
		User                  *User           `json:"user"`
		VideoUrl              string          `json:"video_url"`
		VideoVersions         []*VideoVersion `json:"video_versions"`
	} `json:"data"`
	Detail string `json:"detail"`
}

type Caption struct {
	Text string `json:"text"`
}

type User struct {
	FullName                   string `json:"full_name"`
	HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
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
	ImageVersions       *ImageVersion `json:"image_versions"`
	IsVideo             bool          `json:"is_video"`
	MediaName           string        `json:"media_name"`
	MediaType           int           `json:"media_type"`
	OriginalHeight      int           `json:"original_height"`
	OriginalWidth       int           `json:"original_width"`
	ProductSuggestions  []interface{} `json:"product_suggestions"`
	ProductType         string        `json:"product_type"`
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
