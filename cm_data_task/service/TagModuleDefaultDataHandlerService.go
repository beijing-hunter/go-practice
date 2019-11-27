package service

import (
	"cm_data_task/entitys"
	"cm_data_task/utils"
	"strconv"
	"strings"
)

//标签模块默认数据处理
type TagModuleDefaultDataHandlerService struct{}

func (service TagModuleDefaultDataHandlerService) DefaultDataCollect() {

	moduleInfos := hdao.FindModuleInfos()

	if len(moduleInfos) == 0 {
		utils.InfoLogger.Println("标签模块默认数据处理-没有需要处理的模块信息")
		return
	}

	categoryDatas := hdao.FindCategoryDatas()

	if len(categoryDatas) == 0 {
		utils.InfoLogger.Println("标签模块默认数据处理-没有获取到需要处理的学科")
		return
	}

	for _, moduleInfo := range moduleInfos {

		tagIds := strings.Split(moduleInfo.TagIdStr, ",")

		categoryInfo := categoryDatas[0]

		userId := int64(0)
		service.singleUserDefaultDataCollect(userId, moduleInfo, categoryInfo, tagIds)
	}

	utils.InfoLogger.Println("标签模块默认数据处理-处理完成")
}

//单个用户默认数据收集
func (service TagModuleDefaultDataHandlerService) singleUserDefaultDataCollect(userId int64, moduleInfo entitys.HomepageInfo, categoryInfo entitys.CategoryInfo, tagIds []string) {

	categoryIds := strings.Split(categoryInfo.CategoryIdStr, ",")

	for _, categoryIdS := range categoryIds {

		categoryId, _ := strconv.ParseInt(categoryIdS, 10, 64)

		for _, tagIdS := range tagIds {

			tagId, _ := strconv.ParseInt(tagIdS, 10, 64)
			liveOrderfinalDatas := hdao.FindLiveOrderfinalInfos(moduleInfo.Id, categoryId, tagId)

			if len(liveOrderfinalDatas) == 0 {
				utils.InfoLogger.Printf("标签模块默认数据处理-学科与标签下没有获取到课程:moduleId=%v,categoryId=%v,tagId=%v\n", moduleInfo.Id, categoryId, tagId)
			} else {

				var dataResults []entitys.LiveDataCollectResult

				for _, liveOrderfinalInfo := range liveOrderfinalDatas {

					oValue, _ := strconv.ParseFloat(liveOrderfinalInfo.Orderfinal, 64)
					resultValue := oValue

					result := entitys.LiveDataCollectResult{
						CategoryId: categoryId,
						TagId:      tagId,
						LiveId:     liveOrderfinalInfo.LiveId,
						ModuleId:   moduleInfo.Id,
						UserId:     userId,
						Orderfinal: resultValue,
					}
					result.Livetype, _ = strconv.ParseInt(liveOrderfinalInfo.Livetype, 10, 64)

					dataResults = append(dataResults, result)
				}

				hdao.AddCollectResult(dataResults)
			}
		}
	}

	notifyInfo := entitys.DataCollectNotifyInfo{
		UserId:        userId,
		ModuleId:      moduleInfo.Id,
		CategroyIdStr: categoryInfo.CategoryIdStr,
	}

	DataCollectChan <- notifyInfo
}
