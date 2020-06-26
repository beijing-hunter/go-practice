package service

import (
	"cm_data_task/dao"
	"cm_data_task/utils"
	"sync"
)

var (
	guessDao = dao.GuessLikeModuleDataDao{}
)

type GuessLikeModuleDataHandlerService struct{}

//处理猜你喜欢模块数据
func (service GuessLikeModuleDataHandlerService) Handler() {

	userIds := guessDao.FindUserId()
	moduleId := int64(1)
	tservice := TagModuleDataHandlerService{}

	if len(userIds) == 0 {
		utils.InfoLogger.Println("猜你喜欢模块-没有需要处理的用户信息")
		return
	}

	userCount := len(userIds)
	groupCount := 1000
	groupLength := userCount / groupCount

	for i := 0; i <= groupCount; i++ { //分组执行goruntine

		j := i * groupLength
		groupMaxLength := j + groupLength

		if groupMaxLength > userCount {

			groupMaxLength = userCount
			groupLength = userCount - j
		}

		if groupLength <= 0 {
			break
		}

		wg := sync.WaitGroup{}
		wg.Add(groupLength) //每组执行groupLength个goruntine则进入等待

		for ; j < groupMaxLength; j++ {

			userId := userIds[j]

			go func(uid int64) {
				defer wg.Done()
				service.SingleUserDataHandler(uid, moduleId, tservice)
			}(userId)
		}

		wg.Wait()
	}

	utils.InfoLogger.Println("猜你喜欢模块-处理完成")
}

func (service GuessLikeModuleDataHandlerService) SingleUserDataHandler(userId int64, moduleId int64, tservice TagModuleDataHandlerService) {

	liveUUIDs := guessDao.FindUserMaxPlayPctLiveUUID(userId)
	otherUserIds := guessDao.FindOtherUserId(userId, liveUUIDs)

	if otherUserIds == nil || len(otherUserIds) == 0 {
		utils.InfoLogger.Printf("猜你喜欢模块-没有其他用户观看此课程:userId=%v,liveUUIDS=%v\n", userId, liveUUIDs)
		return
	}

	livePctInfos := guessDao.FindOtherUserMaxPlayPctLiveInfo(otherUserIds)

	if len(livePctInfos) == 0 {
		utils.InfoLogger.Printf("猜你喜欢模块-观看此课程的用户，没有观看过其他课程:userId=%v,liveUUIDS=%v\n", userId, liveUUIDs)
		return
	}

	var complementLiveIds []int64                             //取补集课程id
	userLiveUUIDs := guessDao.FindUserPlayPctLiveUUId(userId) //用户观看过的课程uuid
	userLiveUUIDMap := utils.SliceToMapString(userLiveUUIDs)

	for _, livePctInfo := range livePctInfos {

		_, ok := userLiveUUIDMap[livePctInfo.LiveUUID]
		ok2 := userUseLiveIdCache.IsUseLiveId(userId, livePctInfo.LiveId)

		if !ok && !ok2 { //取补集
			complementLiveIds = append(complementLiveIds, livePctInfo.LiveId)
		}
	}

	if len(complementLiveIds) == 0 {
		utils.InfoLogger.Printf("猜你喜欢模块-没有筛选到可用的补集课程id:userId=%v\n", userId)
		return
	} else {

		liveOrderfinalDatas := guessDao.FindLiveOrderfinalInfo(userId, moduleId, complementLiveIds)

		if len(liveOrderfinalDatas) == 0 {
			utils.InfoLogger.Printf("猜你喜欢模块-没有查询到课程的排序值:userId=%v\n", userId)
			return
		} else {
			utils.InfoLogger.Printf("猜你喜欢模块-筛选到用户喜欢的课程:userId=%v\n", userId)
			hdao.AddRecommendLiveResult(liveOrderfinalDatas)
			userUseLiveIdCache.AddUseLiveId(userId, complementLiveIds)
		}
	}

}
