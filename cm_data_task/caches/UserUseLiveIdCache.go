package caches

import (
	"cm_data_task/utils"
	"strconv"
	"time"
)

type UserUseLiveIdCache struct{}

func (u UserUseLiveIdCache) AddUseLiveId(userId int64, liveIds []int64) {

	key := strconv.FormatInt(userId, 10)
	fieldMaps := make(map[string]interface{}, len(liveIds))

	for _, liveId := range liveIds {
		fieldMaps[strconv.FormatInt(liveId, 10)] = true
	}

	utils.RedisPoolClient.Expire(key, time.Hour*2)
	utils.RedisPoolClient.HMSet(key, fieldMaps)
}

func (u UserUseLiveIdCache) IsUseLiveId(userId int64, liveId int64) bool {

	key := strconv.FormatInt(userId, 10)
	field := strconv.FormatInt(liveId, 10)
	return utils.RedisPoolClient.HExists(key, field).Val()
}
