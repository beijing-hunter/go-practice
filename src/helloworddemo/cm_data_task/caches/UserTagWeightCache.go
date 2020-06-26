package caches

import (
	"cm_data_task/dao"
	"strconv"
)

type UserTagWeightCache struct{}

var (
	cacheMap map[string]string //key:userid.tagid value:weightValue
	hdao     = dao.TagModuleDataDao{}
)

func init() {

	datas := hdao.FindAllUserTagWeights()
	cacheMap = make(map[string]string, len(datas))

	if len(datas) > 0 {

		for _, data := range datas {
			key := strconv.FormatInt(data.UserId, 10) + strconv.FormatInt(data.TagId, 10)
			cacheMap[key] = data.TagWeights
		}
	}
}

func (c UserTagWeightCache) GetValue(key string) (value string) {

	v, ok := cacheMap[key]

	if ok {
		value = v
	} else {
		value = ""
	}

	return value
}

func (c UserTagWeightCache) DelValue(key string) {

	delete(cacheMap, key) //删除操作
}
