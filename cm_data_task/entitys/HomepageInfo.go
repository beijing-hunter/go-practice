package entitys

type HomepageInfo struct {
	Id         int64
	Title      string
	SourceType int8
	ModuleType int8
	TagIdStr   string
}

type ModuleSettingLiveInfo struct {
	LiveId       int64
	DispalyOrder int64
	CategoryId   int64
	Livetype     int64
}

type CategoryInfo struct {
	CategoryIdStr   string
	CategoryNameStr string
}

type LivePlayPctInfo struct {
	LiveId   int64
	LiveUUID string
	Pct      float32
}
