package instagrambulkscrapper

type ParsedPost struct {
	Data struct {
		ImageHd       string `json:"image_hd"`
		VideoHd       string `json:"video_hd"`
		MainMediaHd   string `json:"main_media_hd"`
		MainMediaType string `json:"main_media_type"`
		ChildMediasHd []struct {
			Url  string `json:"url"`
			Type string `json:"type"`
		} `json:"child_medias_hd"`
		Caption string `json:"caption"`
		Owner   struct {
			Id                        string `json:"id"`
			Username                  string `json:"username"`
			IsVerified                bool   `json:"is_verified"`
			ProfilePicUrl             string `json:"profile_pic_url"`
			BlockedByViewer           bool   `json:"blocked_by_viewer"`
			FollowedByViewer          bool   `json:"followed_by_viewer"`
			FullName                  string `json:"full_name"`
			HasBlockedViewer          bool   `json:"has_blocked_viewer"`
			IsEmbedsDisabled          bool   `json:"is_embeds_disabled"`
			IsPrivate                 bool   `json:"is_private"`
			IsUnpublished             bool   `json:"is_unpublished"`
			RequestedByViewer         bool   `json:"requested_by_viewer"`
			PassTieringRecommendation bool   `json:"pass_tiering_recommendation"`
			EdgeOwnerToTimelineMedia  struct {
				Count int `json:"count"`
			} `json:"edge_owner_to_timeline_media"`
			EdgeFollowedBy struct {
				Count int `json:"count"`
			} `json:"edge_followed_by"`
		} `json:"owner"`
	} `json:"data"`
	Status  string `json:"status"`
	Message string `json:"message"`
}
