package entitys

type LiveOrderfinal struct {
	LiveId     int64
	TagId      int64
	CategoryId int64
	Orderfinal string
	Livetype   string
}

type LiveDataCollectResult struct {
	ModuleId   int64
	UserId     int64
	LiveId     int64
	TagId      int64
	CategoryId int64
	Orderfinal float64
	Livetype   int64
}

type RecommendLiveInfo struct {
	MoudleId   int64
	UserId     int64
	LiveId     int64
	LiveType   int64
	CategoryId int64
	Orderfinal float64
	IsSysAuto  int64
	Pos        int64
}
