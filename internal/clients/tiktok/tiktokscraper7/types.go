package tiktokscraper7

type VideoParsed struct {
	Code          int     `json:"code"`
	Msg           string  `json:"msg"`
	ProcessedTime float64 `json:"processed_time"`
	Data          struct {
		AwemeId        string `json:"aweme_id"`
		Id             string `json:"id"`
		Region         string `json:"region"`
		Title          string `json:"title"`
		Cover          string `json:"cover"`
		AiDynamicCover string `json:"ai_dynamic_cover"`
		OriginCover    string `json:"origin_cover"`
		Duration       int    `json:"duration"`
		Play           string `json:"play"`
		Wmplay         string `json:"wmplay"`
		Hdplay         string `json:"hdplay"`
		Size           int    `json:"size"`
		WmSize         int    `json:"wm_size"`
		HdSize         int    `json:"hd_size"`
		Music          string `json:"music"`
		MusicInfo      struct {
			Id       string `json:"id"`
			Title    string `json:"title"`
			Play     string `json:"play"`
			Cover    string `json:"cover"`
			Author   string `json:"author"`
			Original bool   `json:"original"`
			Duration int    `json:"duration"`
			Album    string `json:"album"`
		} `json:"music_info"`
		PlayCount     int         `json:"play_count"`
		DiggCount     int         `json:"digg_count"`
		CommentCount  int         `json:"comment_count"`
		ShareCount    int         `json:"share_count"`
		DownloadCount int         `json:"download_count"`
		CollectCount  int         `json:"collect_count"`
		CreateTime    int         `json:"create_time"`
		Anchors       interface{} `json:"anchors"`
		AnchorsExtras string      `json:"anchors_extras"`
		IsAd          bool        `json:"is_ad"`
		CommerceInfo  struct {
			AdvPromotable          bool `json:"adv_promotable"`
			AuctionAdInvited       bool `json:"auction_ad_invited"`
			BrandedContentType     int  `json:"branded_content_type"`
			WithCommentFilterWords bool `json:"with_comment_filter_words"`
		} `json:"commerce_info"`
		CommercialVideoInfo string `json:"commercial_video_info"`
		ItemCommentSettings int    `json:"item_comment_settings"`
		MentionedUsers      string `json:"mentioned_users"`
		Author              struct {
			Id       string `json:"id"`
			UniqueId string `json:"unique_id"`
			Nickname string `json:"nickname"`
			Avatar   string `json:"avatar"`
		} `json:"author"`
	} `json:"data"`
}
