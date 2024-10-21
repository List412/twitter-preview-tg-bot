package instagramscrapper

type ParsedPost struct {
	Id             string `json:"id"`
	Pk             int64  `json:"pk"`
	Code           string `json:"code"`
	MediaType      int    `json:"media_type"`
	TakenAt        int    `json:"taken_at"`
	ImageVersions2 struct {
		Candidates []ImageCandidate `json:"candidates"`
	} `json:"image_versions2"`
	OriginalWidth      int             `json:"original_width"`
	OriginalHeight     int             `json:"original_height"`
	VideoVersions      []VideoVersion  `json:"video_versions"`
	VideoDuration      float32         `json:"video_duration"`
	CarouselMedia      []CarouselMedia `json:"carousel_media"`
	CarouselMediaCount int             `json:"carousel_media_count"`
	Caption            struct {
		Text               string `json:"text"`
		Pk                 string `json:"pk"`
		UserId             int64  `json:"user_id"`
		Type               int    `json:"type"`
		DidReportAsSpam    bool   `json:"did_report_as_spam"`
		CreatedAt          int    `json:"created_at"`
		CreatedAtUtc       int    `json:"created_at_utc"`
		ContentType        string `json:"content_type"`
		Status             string `json:"status"`
		BitFlags           int    `json:"bit_flags"`
		ShareEnabled       bool   `json:"share_enabled"`
		IsRankedComment    bool   `json:"is_ranked_comment"`
		IsCovered          bool   `json:"is_covered"`
		PrivateReplyStatus int    `json:"private_reply_status"`
		MediaId            int64  `json:"media_id"`
		HasTranslation     bool   `json:"has_translation"`
	} `json:"caption"`
	User struct {
		Pk            int64  `json:"pk"`
		Username      string `json:"username"`
		FullName      string `json:"full_name"`
		IsPrivate     bool   `json:"is_private"`
		ProfilePicUrl string `json:"profile_pic_url"`
		ProfilePicId  string `json:"profile_pic_id"`
		IsVerified    bool   `json:"is_verified"`
		AccountType   int    `json:"account_type"`
		FanClubInfo   struct {
		} `json:"fan_club_info"`
		FbidV2          int64  `json:"fbid_v2"`
		PkId            string `json:"pk_id"`
		StrongId        string `json:"strong_id__"`
		LatestReelMedia int    `json:"latest_reel_media"`
		AvatarStatus    struct {
			HasAvatar bool `json:"has_avatar"`
		} `json:"avatar_status"`
		ShowAccountTransparencyDetails bool `json:"show_account_transparency_details"`
	} `json:"user"`
	ProductType       string `json:"product_type"`
	VideoDashManifest string `json:"video_dash_manifest"`
	VideoManifest     struct {
		Audio interface{} `json:"audio"`
		Video interface{} `json:"video"`
	} `json:"video_manifest"`
	StrongId            string `json:"strong_id__"`
	CommercialityStatus string `json:"commerciality_status"`
	FbUserTags          struct {
		In interface{} `json:"in"`
	} `json:"fb_user_tags"`
	SharingFrictionInfo struct {
		ShouldHaveSharingFriction bool        `json:"should_have_sharing_friction"`
		BloksAppUrl               interface{} `json:"bloks_app_url"`
		SharingFrictionPayload    interface{} `json:"sharing_friction_payload"`
	} `json:"sharing_friction_info"`
	DeviceTimestamp                     int64  `json:"device_timestamp"`
	ClientCacheKey                      string `json:"client_cache_key"`
	IntegrityReviewDecision             string `json:"integrity_review_decision"`
	IsVisualReplyCommenterNoticeEnabled bool   `json:"is_visual_reply_commenter_notice_enabled"`
	IsOrganicProductTaggingEligible     bool   `json:"is_organic_product_tagging_eligible"`
	CommentInformTreatment              struct {
		ShouldHaveInformTreatment bool        `json:"should_have_inform_treatment"`
		Text                      string      `json:"text"`
		Url                       interface{} `json:"url"`
		ActionType                interface{} `json:"action_type"`
	} `json:"comment_inform_treatment"`
	FundraiserTag struct {
		HasStandaloneFundraiser bool `json:"has_standalone_fundraiser"`
	} `json:"fundraiser_tag"`
	OrganicTrackingToken         string `json:"organic_tracking_token"`
	CommentThreadingEnabled      bool   `json:"comment_threading_enabled"`
	MaxNumVisiblePreviewComments int    `json:"max_num_visible_preview_comments"`
	Owner                        struct {
		Pk            int64  `json:"pk"`
		Username      string `json:"username"`
		FullName      string `json:"full_name"`
		IsPrivate     bool   `json:"is_private"`
		ProfilePicUrl string `json:"profile_pic_url"`
		ProfilePicId  string `json:"profile_pic_id"`
		IsVerified    bool   `json:"is_verified"`
		AccountType   int    `json:"account_type"`
		FanClubInfo   struct {
		} `json:"fan_club_info"`
		FbidV2          int64  `json:"fbid_v2"`
		PkId            string `json:"pk_id"`
		StrongId        string `json:"strong_id__"`
		LatestReelMedia int    `json:"latest_reel_media"`
		AvatarStatus    struct {
			HasAvatar bool `json:"has_avatar"`
		} `json:"avatar_status"`
		ShowAccountTransparencyDetails bool `json:"show_account_transparency_details"`
	} `json:"owner"`
	CommentingDisabledForViewer bool   `json:"commenting_disabled_for_viewer"`
	ReshareCount                int    `json:"reshare_count"`
	VideoCodec                  string `json:"video_codec"`
	NumberOfQualities           int    `json:"number_of_qualities"`
}

type VideoVersion struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Url    string `json:"url"`
}

type ImageCandidate struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Url    string `json:"url"`
}

type CarouselMedia struct {
	Id             string `json:"id"`
	Pk             int64  `json:"pk"`
	MediaType      int    `json:"media_type"`
	TakenAt        int    `json:"taken_at"`
	ImageVersions2 struct {
		Candidates []ImageCandidate `json:"candidates"`
	} `json:"image_versions2"`
	OriginalWidth  int         `json:"original_width"`
	OriginalHeight int         `json:"original_height"`
	Caption        interface{} `json:"caption"`
	ProductType    string      `json:"product_type"`
	Usertags       struct {
		In []interface{} `json:"in"`
	} `json:"usertags"`
	VideoManifest struct {
		Audio interface{} `json:"audio"`
		Video interface{} `json:"video"`
	} `json:"video_manifest"`
	CarouselParentId    string `json:"carousel_parent_id"`
	StrongId            string `json:"strong_id__"`
	CommercialityStatus string `json:"commerciality_status"`
	Preview             string `json:"preview,omitempty"`
	FbUserTags          struct {
		In []interface{} `json:"in"`
	} `json:"fb_user_tags"`
	SharingFrictionInfo struct {
		ShouldHaveSharingFriction bool        `json:"should_have_sharing_friction"`
		BloksAppUrl               interface{} `json:"bloks_app_url"`
		SharingFrictionPayload    interface{} `json:"sharing_friction_payload"`
	} `json:"sharing_friction_info"`
	CommentInformTreatment struct {
		ShouldHaveInformTreatment bool        `json:"should_have_inform_treatment"`
		Text                      string      `json:"text"`
		Url                       interface{} `json:"url"`
		ActionType                interface{} `json:"action_type"`
	} `json:"comment_inform_treatment"`
	FundraiserTag struct {
		HasStandaloneFundraiser bool `json:"has_standalone_fundraiser"`
	} `json:"fundraiser_tag"`
	ReshareCount int `json:"reshare_count"`
}
