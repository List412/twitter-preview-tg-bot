package scrapper

type SelfReplay struct {
	Text string
	Id   string
}

type Collab struct {
	Name       string
	ScreenName string
}

type PageScrapperResult struct {
	SelfReplay  []SelfReplay
	CollabUsers []Collab
}
