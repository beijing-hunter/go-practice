package service

import (
	"cm_data_task/caches"
	"cm_data_task/dao"
	"cm_data_task/entitys"
	"cm_data_task/utils"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	DataCollectChan        chan entitys.DataCollectNotifyInfo
	hdao                   = dao.TagModuleDataDao{}
	userTagCache           = caches.UserTagWeightCache{}
	liveOrderfinalCache    = caches.LiveOrderfinalCache{}
	moduleSettingLiveCache = caches.ModuleSettingLiveCache{}
	userUseLiveIdCache     = caches.UserUseLiveIdCache{}
)

//标签模块数据处理
type TagModuleDataHandlerService struct{}

func init() {
	DataCollectChan = make(chan entitys.DataCollectNotifyInfo, 1000)
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

		goroutineCount := 0
		tagIds := strings.Split(moduleInfo.TagIdStr, ",")

		for _, userCategoryInfo := range userCategoryDatas {

			go service.singleUserDataCollect(moduleInfo, userCategoryInfo, tagIds)
			goroutineCount++

			if goroutineCount > 500 {
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

	moduleSettingLiveInfo := moduleSettingLiveCache.GetValue(notifyInfo.ModuleId)
	categoryIds := strings.Split(notifyInfo.CategroyIdStr, ",")

	for _, categoryIdS := range categoryIds {

		categoryId, _ := strconv.ParseInt(categoryIdS, 10, 64)
		var recommendLiveDataResult []entitys.RecommendLiveInfo
		var recommentResultLiveId []int64
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
					recommentResultLiveId = append(recommentResultLiveId, result.LiveId)
				}
			}
		}

		collectRestulDatas, _ := notifyInfo.CollectResultData.(LiveDataCollectResultSort)

		if len(collectRestulDatas) == 0 {
			utils.InfoLogger.Printf("标签模块数据加工-用户订阅的学科没有收集到数据:ModuleId=%v,categoryId=%v,UserId=%v\n", notifyInfo.ModuleId, categoryId, notifyInfo.UserId)
		} else {

			for _, collectInfo := range collectRestulDatas {

				if collectInfo.CategoryId == categoryId {

					_, ok := resultLiveMap[collectInfo.LiveId]
					ok2 := userUseLiveIdCache.IsUseLiveId(notifyInfo.UserId, collectInfo.LiveId)

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
						recommentResultLiveId = append(recommentResultLiveId, result.LiveId)

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
			userUseLiveIdCache.AddUseLiveId(notifyInfo.UserId, recommentResultLiveId)

		} else {
			utils.InfoLogger.Printf("标签模块数据加工-用户订阅的学科没有匹配到数据:ModuleId=%v,categoryId=%v,UserId=%v\n", notifyInfo.ModuleId, categoryId, notifyInfo.UserId)
		}
	}
}

//单个用户数据收集
func (service TagModuleDataHandlerService) singleUserDataCollect(moduleInfo entitys.HomepageInfo, userCategoryInfo entitys.UserCategoryWeights, tagIds []string) {

	categoryIds := strings.Split(userCategoryInfo.CategroyIdStr, ",")
	categoryWeights := strings.Split(userCategoryInfo.CategroryWeightsStr, ",")
	dataCollectResultMap := make(map[int64]entitys.LiveDataCollectResult, 300)

	for cIndex, categoryIdS := range categoryIds {

		categoryId, _ := strconv.ParseInt(categoryIdS, 10, 64)

		for _, tagIdS := range tagIds {

			tagId, _ := strconv.ParseInt(tagIdS, 10, 64)

			tagKey := strconv.FormatInt(userCategoryInfo.UserId, 10) + strconv.FormatInt(tagId, 10)
			tagWeight := userTagCache.GetValue(tagKey)

			if tagWeight == "" {
				tagWeight = "0"
			}

			liveKey := strconv.FormatInt(moduleInfo.Id, 10) + strconv.FormatInt(categoryId, 10) + strconv.FormatInt(tagId, 10)
			liveOrderfinalDatas := liveOrderfinalCache.GetValue(liveKey)

			if len(liveOrderfinalDatas) == 0 {
				utils.InfoLogger.Printf("标签模块数据收集-学科与标签下没有获取到课程:moduleId=%v,categoryId=%v,tagId=%v\n", moduleInfo.Id, categoryId, tagId)
			} else {

				for _, liveOrderfinalInfo := range liveOrderfinalDatas {

					oValue, _ := strconv.ParseFloat(liveOrderfinalInfo.Orderfinal, 64)
					cWValue, _ := strconv.ParseFloat(categoryWeights[cIndex], 64)
					tWValue, _ := strconv.ParseFloat(tagWeight, 64)
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
					liveOrderfinalInfoValue, ok := dataCollectResultMap[liveOrderfinalInfo.LiveId]

					if !ok {
						dataCollectResultMap[liveOrderfinalInfo.LiveId] = result
					} else {

						if liveOrderfinalInfoValue.Orderfinal < result.Orderfinal {
							liveOrderfinalInfoValue.Orderfinal = result.Orderfinal
							dataCollectResultMap[liveOrderfinalInfo.LiveId] = liveOrderfinalInfoValue
						}
					}
				}

				utils.InfoLogger.Printf("标签模块数据收集-用户相关课程数据收集完成:moduleId=%v,userId=%v,categoryId=%v,tagId=%v\n", moduleInfo.Id, userCategoryInfo.UserId, categoryId, tagId)
			}
		}
	}

	var resultSort LiveDataCollectResultSort

	if len(dataCollectResultMap) > 0 {

		for _, value := range dataCollectResultMap {

			resultSort = append(resultSort, value)
		}

		sort.Sort(resultSort)
	}

	service.pushCollectChan(moduleInfo.Id, userCategoryInfo, resultSort)
}

//单个用户数据收集完成通知工厂加工数据
func (service TagModuleDataHandlerService) pushCollectChan(moduleId int64, userCategoryInfo entitys.UserCategoryWeights, resultSort LiveDataCollectResultSort) {
	notifyInfo := entitys.DataCollectNotifyInfo{
		UserId:            userCategoryInfo.UserId,
		ModuleId:          moduleId,
		CategroyIdStr:     userCategoryInfo.CategroyIdStr,
		CollectResultData: resultSort,
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

type LiveDataCollectResultSort []entitys.LiveDataCollectResult

func (s LiveDataCollectResultSort) Len() int {
	return len(s)
}

func (s LiveDataCollectResultSort) Less(i, j int) bool {
	return s[i].Orderfinal > s[j].Orderfinal
}

func (s LiveDataCollectResultSort) Swap(i, j int) {

	temp := s[i]
	s[i] = s[j]
	s[j] = temp
}
