package service

import (
	"cm_data_task/dao"
	"cm_data_task/utils"
	"time"
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

	goroutineCount := 0

	for i := 0; i < len(userIds); i++ {
		userId := userIds[i]
		go service.SingleUserDataHandler(userId, moduleId, tservice)
		goroutineCount++

		if goroutineCount > 30 {
			time.Sleep(time.Second * 20) //休眠10秒
			goroutineCount = 0
		}

		if (len(userIds) - i) == 10 { //当剩余处理用户还有10个时：休眠40秒，略微等待其它goroutine数据处理
			time.Sleep(time.Second * 40) //休眠40秒
		}
	}

	utils.InfoLogger.Println("猜你喜欢模块-处理完成")
}

func (service GuessLikeModuleDataHandlerService) SingleUserDataHandler(userId int64, moduleId int64, tservice TagModuleDataHandlerService) {

	liveUUIDs := guessDao.FindUserMaxPlayPctLiveUUID(userId)
	livePctInfos := guessDao.FindOtherUserMaxPlayPctLiveInfo(userId, liveUUIDs)

	if len(livePctInfos) == 0 {
		utils.InfoLogger.Printf("猜你喜欢模块-观看此课程的用户，没有观看过其他课程:userId=%v,liveUUIDS=%v\n", userId, liveUUIDs)
		return
	}

	var complementLiveIds []int64                             //取补集课程id
	userLiveUUIDs := guessDao.FindUserPlayPctLiveUUId(userId) //用户观看过的课程uuid
	userLiveUUIDMap := utils.SliceToMapString(userLiveUUIDs)

	useLiveIdDatas := hdao.FindUseLiveId(moduleId, userId) //自定义模块推荐给用户的课程id
	useLiveIdMap := tservice.converMap(useLiveIdDatas)

	for _, livePctInfo := range livePctInfos {

		_, ok := userLiveUUIDMap[livePctInfo.LiveUUID]
		_, ok2 := useLiveIdMap[livePctInfo.LiveId]

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
		}
	}

}
