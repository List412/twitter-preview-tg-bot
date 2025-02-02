package tiktok89

type VideoParsed struct {
	Ok           bool   `json:"ok"`
	ErrorMessage string `json:"errorMessage"`
	Author       struct {
		CreateTime               int    `json:"create_time"`
		CustomVerify             string `json:"custom_verify"`
		CvLevel                  string `json:"cv_level"`
		DownloadPromptTs         int    `json:"download_prompt_ts"`
		DownloadSetting          int    `json:"download_setting"`
		DuetSetting              int    `json:"duet_setting"`
		EnabledFilterAllComments bool   `json:"enabled_filter_all_comments"`
		EnterpriseVerifyReason   string `json:"enterprise_verify_reason"`

		FavoritingCount int           `json:"favoriting_count"`
		FbExpireTime    int           `json:"fb_expire_time"`
		FollowStatus    int           `json:"follow_status"`
		FollowerCount   int           `json:"follower_count"`
		FollowerStatus  int           `json:"follower_status"`
		FollowersDetail interface{}   `json:"followers_detail"`
		FollowingCount  int           `json:"following_count"`
		FriendsStatus   int           `json:"friends_status"`
		Gender          int           `json:"gender"`
		Geofencing      []interface{} `json:"geofencing"`

		HideSearch          bool        `json:"hide_search"`
		HomepageBottomToast interface{} `json:"homepage_bottom_toast"`

		ItemList interface{} `json:"item_list"`
		Language string      `json:"language"`

		MentionStatus         int         `json:"mention_status"`
		MutualRelationAvatars interface{} `json:"mutual_relation_avatars"`
		NeedPoints            interface{} `json:"need_points"`
		NeedRecommend         int         `json:"need_recommend"`
		Nickname              string      `json:"nickname"`

		Uid                string        `json:"uid"`
		UniqueId           string        `json:"unique_id"`
		UniqueIdModifyTime int           `json:"unique_id_modify_time"`
		UserCanceled       bool          `json:"user_canceled"`
		UserMode           int           `json:"user_mode"`
		UserPeriod         int           `json:"user_period"`
		UserProfileGuide   interface{}   `json:"user_profile_guide"`
		UserRate           int           `json:"user_rate"`
		UserSparkInfo      []interface{} `json:"user_spark_info"`
		UserTags           interface{}   `json:"user_tags"`
		VerificationType   int           `json:"verification_type"`
		VerifyInfo         string        `json:"verify_info"`
	} `json:"author"`
	AuthorUserId int64 `json:"author_user_id"`

	ContentDesc      string        `json:"content_desc"`
	ContentDescExtra []interface{} `json:"content_desc_extra"`

	ContentOriginalType int `json:"content_original_type"`
	ContentSizeType     int `json:"content_size_type"`

	CreateTime int `json:"create_time"`

	Desc string `json:"desc"`

	Duration              int `json:"duration"`
	FollowUpPublishFromId int `json:"follow_up_publish_from_id"`

	LabelTopText interface{} `json:"label_top_text"`
	LongVideo    interface{} `json:"long_video"`

	OriginalClientText struct {
		MarkupText string      `json:"markup_text"`
		TextExtra  interface{} `json:"text_extra"`
	} `json:"original_client_text"`
	PickedUsers             []interface{} `json:"picked_users"`
	PlaylistBlocked         bool          `json:"playlist_blocked"`
	PoiReTagSignal          int           `json:"poi_re_tag_signal"`
	Position                interface{}   `json:"position"`
	PreventDownload         bool          `json:"prevent_download"`
	ProductsInfo            interface{}   `json:"products_info"`
	QuestionList            interface{}   `json:"question_list"`
	Rate                    int           `json:"rate"`
	ReferenceTtsVoiceIds    interface{}   `json:"reference_tts_voice_ids"`
	ReferenceVoiceFilterIds interface{}   `json:"reference_voice_filter_ids"`
	Region                  string        `json:"region"`

	SearchHighlight interface{} `json:"search_highlight"`

	Statistics struct {
		AwemeId            string `json:"aweme_id"`
		CollectCount       int    `json:"collect_count"`
		CommentCount       int    `json:"comment_count"`
		DiggCount          int    `json:"digg_count"`
		DownloadCount      int    `json:"download_count"`
		ForwardCount       int    `json:"forward_count"`
		LoseCommentCount   int    `json:"lose_comment_count"`
		LoseCount          int    `json:"lose_count"`
		PlayCount          int    `json:"play_count"`
		RepostCount        int    `json:"repost_count"`
		ShareCount         int    `json:"share_count"`
		WhatsappShareCount int    `json:"whatsapp_share_count"`
	} `json:"statistics"`

	TextExtra            []interface{} `json:"text_extra"`
	TitleLanguage        string        `json:"title_language"`
	TtsVoiceIds          interface{}   `json:"tts_voice_ids"`
	TttProductRecallType int           `json:"ttt_product_recall_type"`

	UserDigged int   `json:"user_digged"`
	Video      Video `json:"video"`

	VideoLabels          []interface{} `json:"video_labels"`
	VideoText            []interface{} `json:"video_text"`
	VoiceFilterIds       interface{}   `json:"voice_filter_ids"`
	WithPromotionalMusic bool          `json:"with_promotional_music"`
	WithoutWatermark     bool          `json:"without_watermark"`
}
type Video struct {
	BigThumbs []interface{} `json:"big_thumbs"`
	BitRate   []struct {
		HDRBit    string      `json:"HDR_bit"`
		HDRType   string      `json:"HDR_type"`
		BitRate   int         `json:"bit_rate"`
		DubInfos  interface{} `json:"dub_infos"`
		GearName  string      `json:"gear_name"`
		IsBytevc1 int         `json:"is_bytevc1"`
		IsH265    int         `json:"is_h265"`
		PlayAddr  struct {
			DataSize  int         `json:"data_size"`
			FileCs    string      `json:"file_cs"`
			FileHash  string      `json:"file_hash"`
			Height    int         `json:"height"`
			Uri       string      `json:"uri"`
			UrlKey    string      `json:"url_key"`
			UrlList   []string    `json:"url_list"`
			UrlPrefix interface{} `json:"url_prefix"`
			Width     int         `json:"width"`
		} `json:"play_addr"`
		QualityType int    `json:"quality_type"`
		VideoExtra  string `json:"video_extra"`
	} `json:"bit_rate"`
	BitRateAudio  []interface{} `json:"bit_rate_audio"`
	CdnUrlExpired int           `json:"cdn_url_expired"`

	DownloadAddr struct {
		DataSize  int         `json:"data_size"`
		Height    int         `json:"height"`
		Uri       string      `json:"uri"`
		UrlList   []string    `json:"url_list"`
		UrlPrefix interface{} `json:"url_prefix"`
		Width     int         `json:"width"`
	} `json:"download_addr"`

	Duration int `json:"duration"`

	HasDownloadSuffixLogoAddr bool   `json:"has_download_suffix_logo_addr"`
	HasWatermark              bool   `json:"has_watermark"`
	Height                    int    `json:"height"`
	IsBytevc1                 int    `json:"is_bytevc1"`
	IsCallback                bool   `json:"is_callback"`
	IsH265                    int    `json:"is_h265"`
	Meta                      string `json:"meta"`
	MiscDownloadAddrs         string `json:"misc_download_addrs"`

	PlayAddr struct {
		DataSize  int         `json:"data_size"`
		FileCs    string      `json:"file_cs"`
		FileHash  string      `json:"file_hash"`
		Height    int         `json:"height"`
		Uri       string      `json:"uri"`
		UrlKey    string      `json:"url_key"`
		UrlList   []string    `json:"url_list"`
		UrlPrefix interface{} `json:"url_prefix"`
		Width     int         `json:"width"`
	} `json:"play_addr"`
	Ratio         string      `json:"ratio"`
	SourceHDRType int         `json:"source_HDR_type"`
	Tags          interface{} `json:"tags"`
	Width         int         `json:"width"`
}
