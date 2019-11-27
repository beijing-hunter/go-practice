package entitys

type UserCategoryWeights struct {
	UserId              int64
	CategroyIdStr       string
	CategroryWeightsStr string
}

type UserTagWeights struct {
	UserId     int64
	TagId      int64
	TagWeights string
}

type DataCollectNotifyInfo struct {
	UserId        int64
	CategroyIdStr string
	ModuleId      int64
}
