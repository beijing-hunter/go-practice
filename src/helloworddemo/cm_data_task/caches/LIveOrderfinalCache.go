package caches

import (
	"cm_data_task/entitys"
	"strconv"
)

var (
	cacheLiveMap map[string][]entitys.LiveOrderfinal
)

type LiveOrderfinalCache struct{}

func init() {

	datas := hdao.FindAllLiveOrderfinalInfos()
	cacheLiveMap = make(map[string][]entitys.LiveOrderfinal, len(datas))

	if len(datas) > 0 {

		for _, data := range datas {

			key := strconv.FormatInt(data.ModuleId, 10) + strconv.FormatInt(data.CategoryId, 10) + strconv.FormatInt(data.TagId, 10)

			values, _ := cacheLiveMap[key]

			if len(values) < 30 {
				values = append(values, data)
				cacheLiveMap[key] = values
			}
		}
	}
}

func (c LiveOrderfinalCache) GetValue(key string) (values []entitys.LiveOrderfinal) {

	v, ok := cacheLiveMap[key]

	if ok {
		values = v
	} else {
		values = nil
	}

	return
}

func (c LiveOrderfinalCache) DelValue(key string) {

	delete(cacheLiveMap, key) //删除操作
}
