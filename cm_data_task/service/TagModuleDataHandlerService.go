package service

import (
	"cm_data_task/dao"
	"cm_data_task/entitys"
	"cm_data_task/utils"
	"strconv"
	"strings"
	"time"
)

var (
	DataCollectChan chan entitys.DataCollectNotifyInfo
	hdao            = dao.TagModuleDataDao{}
)

//标签模块数据处理
type TagModuleDataHandlerService struct{}

func init() {
	DataCollectChan = make(chan entitys.DataCollectNotifyInfo, 200)
}

func (service TagModuleDataHandlerService) ClearDatas(isallClear bool) {

	hdao.ClearDatas(isallClear)
}

//标签模块数据采集
func (service TagModuleDataHandlerService) TagModuleDataCollect(isCollectEnd bool) {

	service.SetCollectSign(0)
	defer service.SetCollectSign(1)

	moduleInfos := hdao.FindModuleInfos()

	if len(moduleInfos) == 0 {
		utils.InfoLogger.Println("标签模块数据收集-没有需要处理的模块信息")
		return
	}

	userCategoryDatas := hdao.FindUserCategoryWeights()

	if len(userCategoryDatas) == 0 {
		utils.InfoLogger.Println("标签模块数据收集-没有需要处理的用户学科数据")
		return
	}

	for _, moduleInfo := range moduleInfos {

		var collectUserIds []int64 //收集过的userid
		goroutineCount := 0
		tagIds := strings.Split(moduleInfo.TagIdStr, ",")

		if isCollectEnd == false {
			collectUserIds = hdao.FindCollectModuleUserIds(moduleInfo.Id)
		}

		collectUserIdMap := utils.SliceToMapInt(collectUserIds)

		for _, userCategoryInfo := range userCategoryDatas {

			_, ok := collectUserIdMap[userCategoryInfo.UserId]

			if !ok { //未收集
				go service.singleUserDataCollect(moduleInfo, userCategoryInfo, tagIds)
				goroutineCount++
			} else {
				service.pushCollectChan(moduleInfo.Id, userCategoryInfo)
			}

			if goroutineCount > 50 || len(DataCollectChan) > 40 {
				time.Sleep(time.Second * 10) //休眠10秒
				goroutineCount = 0
			}
		}
	}
}

//标签模块数据加工厂
func (service TagModuleDataHandlerService) TagModuleDataMachiningFacatory() {

	for {
		select {
		case notifyInfo := <-DataCollectChan:
			utils.InfoLogger.Printf("标签模块数据加工-订单信息:moduleId=%v,userId=%v,categroyIdStr=%v\n", notifyInfo.ModuleId, notifyInfo.UserId, notifyInfo.CategroyIdStr)
			go service.tagModuleDataMachining(notifyInfo)
		default:

		}
	}
}

//标签模块数据加工
func (service TagModuleDataHandlerService) tagModuleDataMachining(notifyInfo entitys.DataCollectNotifyInfo) {

	moduleSettingLiveInfo := hdao.FindModuleSettingLiveInfo(notifyInfo.ModuleId)
	useLiveIdDatas := hdao.FindUseLiveId(notifyInfo.ModuleId, notifyInfo.UserId)
	useLiveIdMap := service.converMap(useLiveIdDatas)
	categoryIds := strings.Split(notifyInfo.CategroyIdStr, ",")

	for _, categoryIdS := range categoryIds {

		categoryId, _ := strconv.ParseInt(categoryIdS, 10, 64)
		var recommendLiveDataResult []entitys.RecommendLiveInfo
		resultLiveMap := make(map[int64]int64, 30)

		if len(moduleSettingLiveInfo) != 0 {

			for _, sliveInfo := range moduleSettingLiveInfo {

				if categoryId == sliveInfo.CategoryId {

					result := entitys.RecommendLiveInfo{
						MoudleId:   notifyInfo.ModuleId,
						UserId:     notifyInfo.UserId,
						LiveId:     sliveInfo.LiveId,
						LiveType:   sliveInfo.Livetype,
						CategoryId: sliveInfo.CategoryId,
						Orderfinal: 1,
						IsSysAuto:  1,
						Pos:        sliveInfo.DispalyOrder,
					}

					resultLiveMap[sliveInfo.LiveId] = sliveInfo.LiveId
					recommendLiveDataResult = append(recommendLiveDataResult, result)
				}
			}
		}

		collectRestulDatas := hdao.FindCollectResult(notifyInfo.ModuleId, categoryId, notifyInfo.UserId)

		if len(collectRestulDatas) == 0 {
			utils.InfoLogger.Printf("标签模块数据加工-用户订阅的学科没有收集到数据:ModuleId=%v,categoryId=%v,UserId=%v\n", notifyInfo.ModuleId, categoryId, notifyInfo.UserId)
		} else {

			for _, collectInfo := range collectRestulDatas {

				if collectInfo.CategoryId == categoryId {

					_, ok := resultLiveMap[collectInfo.LiveId]
					_, ok2 := useLiveIdMap[collectInfo.LiveId]

					if !ok && !ok2 {

						result := entitys.RecommendLiveInfo{
							MoudleId:   notifyInfo.ModuleId,
							UserId:     notifyInfo.UserId,
							LiveId:     collectInfo.LiveId,
							LiveType:   collectInfo.Livetype,
							CategoryId: collectInfo.CategoryId,
							Orderfinal: collectInfo.Orderfinal,
							IsSysAuto:  0,
							Pos:        18,
						}

						resultLiveMap[collectInfo.LiveId] = collectInfo.LiveId
						recommendLiveDataResult = append(recommendLiveDataResult, result)

						if len(recommendLiveDataResult) >= 10 {
							break
						}
					}
				}
			}
		}

		if len(recommendLiveDataResult) > 0 {
			utils.InfoLogger.Printf("标签模块数据加工-用户推荐课程数据加工完成:ModuleId=%v,categoryId=%v,UserId=%v\n", notifyInfo.ModuleId, categoryId, notifyInfo.UserId)
			hdao.AddRecommendLiveResult(recommendLiveDataResult)
		} else {
			utils.InfoLogger.Printf("标签模块数据加工-用户订阅的学科没有匹配到数据:ModuleId=%v,categoryId=%v,UserId=%v\n", notifyInfo.ModuleId, categoryId, notifyInfo.UserId)
		}
	}
}

//单个用户数据收集
func (service TagModuleDataHandlerService) singleUserDataCollect(moduleInfo entitys.HomepageInfo, userCategoryInfo entitys.UserCategoryWeights, tagIds []string) {

	categoryIds := strings.Split(userCategoryInfo.CategroyIdStr, ",")
	categoryWeights := strings.Split(userCategoryInfo.CategroryWeightsStr, ",")

	for cIndex, categoryIdS := range categoryIds {

		categoryId, _ := strconv.ParseInt(categoryIdS, 10, 64)

		for _, tagIdS := range tagIds {

			tagId, _ := strconv.ParseInt(tagIdS, 10, 64)
			userTagDatas := hdao.FindUserTagWeights(userCategoryInfo.UserId, tagId)

			if len(userTagDatas) == 0 {

				tagInfo := entitys.UserTagWeights{UserId: userCategoryInfo.UserId, TagId: tagId, TagWeights: "0"}
				userTagDatas = append(userTagDatas, tagInfo)
			}

			liveOrderfinalDatas := hdao.FindLiveOrderfinalInfos(moduleInfo.Id, categoryId, tagId)

			if len(liveOrderfinalDatas) == 0 {
				utils.InfoLogger.Printf("标签模块数据收集-学科与标签下没有获取到课程:moduleId=%v,categoryId=%v,tagId=%v\n", moduleInfo.Id, categoryId, tagId)
			} else {

				var dataResults []entitys.LiveDataCollectResult

				for _, liveOrderfinalInfo := range liveOrderfinalDatas {

					oValue, _ := strconv.ParseFloat(liveOrderfinalInfo.Orderfinal, 64)
					cWValue, _ := strconv.ParseFloat(categoryWeights[cIndex], 64)
					tWValue, _ := strconv.ParseFloat(userTagDatas[0].TagWeights, 64)
					resultValue := oValue * cWValue * tWValue

					result := entitys.LiveDataCollectResult{
						CategoryId: categoryId,
						TagId:      tagId,
						LiveId:     liveOrderfinalInfo.LiveId,
						ModuleId:   moduleInfo.Id,
						UserId:     userCategoryInfo.UserId,
						Orderfinal: resultValue,
					}
					result.Livetype, _ = strconv.ParseInt(liveOrderfinalInfo.Livetype, 10, 64)

					dataResults = append(dataResults, result)
				}

				utils.InfoLogger.Printf("标签模块数据收集-用户相关课程数据收集完成:moduleId=%v,userId=%v,categoryId=%v,tagId=%v\n", moduleInfo.Id, userCategoryInfo.UserId, categoryId, tagId)
				hdao.AddCollectResult(dataResults)
			}
		}
	}

	service.pushCollectChan(moduleInfo.Id, userCategoryInfo)
}

//单个用户数据收集完成通知工厂加工数据
func (service TagModuleDataHandlerService) pushCollectChan(moduleId int64, userCategoryInfo entitys.UserCategoryWeights) {
	notifyInfo := entitys.DataCollectNotifyInfo{
		UserId:        userCategoryInfo.UserId,
		ModuleId:      moduleId,
		CategroyIdStr: userCategoryInfo.CategroyIdStr,
	}
	DataCollectChan <- notifyInfo
}

func (service TagModuleDataHandlerService) converMap(liveInfoDatas []entitys.RecommendLiveInfo) (liveIdMap map[int64]int64) {

	if len(liveInfoDatas) == 0 {
		liveIdMap = make(map[int64]int64, 1)
	} else {
		liveIdMap = make(map[int64]int64, len(liveInfoDatas))

		for _, liveInfo := range liveInfoDatas {
			liveIdMap[liveInfo.LiveId] = liveInfo.LiveId
		}
	}
	return
}

//设置数据收集标记 0：开始，1：结束
func (service TagModuleDataHandlerService) SetCollectSign(isStart int64) {

	result := entitys.LiveDataCollectResult{
		CategoryId: -888,
		TagId:      -888,
		LiveId:     -888,
		ModuleId:   -888,
		UserId:     isStart,
		Orderfinal: 0.00,
	}

	result.Livetype, _ = strconv.ParseInt("0", 10, 64)

	var dataResults []entitys.LiveDataCollectResult
	dataResults = append(dataResults, result)
	hdao.AddCollectResult(dataResults)
}

func (service TagModuleDataHandlerService) IsCollectEnd() (result bool) {

	count := hdao.FindCollectSign()
	result = count >= 2
	return
}
