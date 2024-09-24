package tiktok_scaraper2

type Author struct {
	AvatarLarger    string `json:"avatarLarger"`
	AvatarMedium    string `json:"avatarMedium"`
	AvatarThumb     string `json:"avatarThumb"`
	CommentSetting  int    `json:"commentSetting"`
	DownloadSetting int    `json:"downloadSetting"`
	DuetSetting     int    `json:"duetSetting"`
	Ftc             bool   `json:"ftc"`
	Id              string `json:"id"`
	IsADVirtual     bool   `json:"isADVirtual"`
	IsEmbedBanned   bool   `json:"isEmbedBanned"`
	Nickname        string `json:"nickname"`
	OpenFavorite    bool   `json:"openFavorite"`
	PrivateAccount  bool   `json:"privateAccount"`
	Relation        int    `json:"relation"`
	SecUid          string `json:"secUid"`
	Secret          bool   `json:"secret"`
	Signature       string `json:"signature"`
	StitchSetting   int    `json:"stitchSetting"`
	UniqueId        string `json:"uniqueId"`
	Verified        bool   `json:"verified"`
}

type Stats struct {
	CollectCount string `json:"collectCount"`
	CommentCount string `json:"commentCount"`
	DiggCount    string `json:"diggCount"`
	PlayCount    string `json:"playCount"`
	RepostCount  string `json:"repostCount"`
	ShareCount   string `json:"shareCount"`
}

type Video struct {
	VQScore     string `json:"VQScore"`
	Bitrate     int    `json:"bitrate"`
	BitrateInfo []struct {
		Bitrate   int    `json:"Bitrate"`
		CodecType string `json:"CodecType"`
		GearName  string `json:"GearName"`
		MVMAF     string `json:"MVMAF"`
		PlayAddr  struct {
			DataSize int      `json:"DataSize"`
			FileCs   string   `json:"FileCs"`
			FileHash string   `json:"FileHash"`
			Height   int      `json:"Height"`
			Uri      string   `json:"Uri"`
			UrlKey   string   `json:"UrlKey"`
			UrlList  []string `json:"UrlList"`
			Width    int      `json:"Width"`
		} `json:"PlayAddr"`
		QualityType int `json:"QualityType"`
	} `json:"bitrateInfo"`
	CodecType     string `json:"codecType"`
	Cover         string `json:"cover"`
	Definition    string `json:"definition"`
	DownloadAddr  string `json:"downloadAddr"`
	Duration      int    `json:"duration"`
	DynamicCover  string `json:"dynamicCover"`
	EncodeUserTag string `json:"encodeUserTag"`
	EncodedType   string `json:"encodedType"`
	Format        string `json:"format"`
	Height        int    `json:"height"`
	Id            string `json:"id"`
	OriginCover   string `json:"originCover"`
	PlayAddr      string `json:"playAddr"`
	Ratio         string `json:"ratio"`
	VideoQuality  string `json:"videoQuality"`
	VolumeInfo    struct {
		Loudness float64 `json:"Loudness"`
		Peak     float64 `json:"Peak"`
	} `json:"volumeInfo"`
	Width     int `json:"width"`
	ZoomCover struct {
		Field1 string `json:"240"`
		Field2 string `json:"480"`
		Field3 string `json:"720"`
		Field4 string `json:"960"`
	} `json:"zoomCover"`
}

type VideoParsed struct {
	ItemInfo struct {
		ItemStruct struct {
			Author                     Author `json:"author"`
			BackendSourceEventTracking string `json:"backendSourceEventTracking"`
			Collected                  bool   `json:"collected"`
			CreateTime                 int    `json:"createTime"`
			Desc                       string `json:"desc"`
			Digged                     bool   `json:"digged"`
			DuetDisplay                int    `json:"duetDisplay"`
			DuetEnabled                bool   `json:"duetEnabled"`
			ForFriend                  bool   `json:"forFriend"`
			Id                         string `json:"id"`
			ItemCommentStatus          int    `json:"itemCommentStatus"`
			ItemControl                struct {
				CanRepost bool `json:"can_repost"`
			} `json:"item_control"`
			StatsV2 Stats `json:"statsV2"`
			Video   Video `json:"video"`
		} `json:"itemStruct"`
	} `json:"itemInfo"`
	StatusCode  int    `json:"statusCode"`
	StatusCode1 int    `json:"status_code"`
	StatusMsg   string `json:"status_msg"`
}
