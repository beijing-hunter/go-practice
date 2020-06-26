package caches

import "cm_data_task/entitys"

type ModuleSettingLiveCache struct{}

var (
	cacheSettingLive map[int64][]entitys.ModuleSettingLiveInfo
)

func init() {
	cacheSettingLive = make(map[int64][]entitys.ModuleSettingLiveInfo, 12)
}

func (m ModuleSettingLiveCache) GetValue(moduleId int64) []entitys.ModuleSettingLiveInfo {

	values, ok := cacheSettingLive[moduleId]

	if !ok {
		values = hdao.FindModuleSettingLiveInfo(moduleId)
		cacheSettingLive[moduleId] = values
	}

	return values
}
