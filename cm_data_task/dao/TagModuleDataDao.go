package dao

import (
	"cm_data_task/entitys"
	"cm_data_task/utils"
	"strconv"
)

type TagModuleDataDao struct{}

//查询模块信息
func (home TagModuleDataDao) FindModuleInfos() (datas []entitys.HomepageInfo) {

	sql := "select hm.id,hm.title,hm.source_type,hm.module_type,GROUP_CONCAT(hmt.tag_id) as tagIdStr " +
		"from ods_cm_homepage_module hm " +
		"left join ods_cm_homepage_module_tag hmt on hm.id=hmt.module_id " +
		"WHERE hm.source_type=2 and hm.`status`=1 and hm.id not in (2,3,4,5) " +
		"GROUP BY hm.id; "
	rows, err := utils.Db.Query(sql)
	defer utils.ErrorCatch()
	defer rows.Close()

	if err != nil {
		panic(err)
		return
	}

	for rows.Next() {

		info := entitys.HomepageInfo{}
		rows.Scan(&info.Id, &info.Title, &info.SourceType, &info.ModuleType, &info.TagIdStr)
		datas = append(datas, info)
	}

	return
}

//查询课程排序值
func (home TagModuleDataDao) FindLiveOrderfinalInfos(moduleId int64, categoryId int64, tagId int64) (datas []entitys.LiveOrderfinal) {

	sql := "SELECT DISTINCT lt.liveId,lt.tagId,cep.categoryId,cep.orderfinal,cep.livetype " +
		"from ods_cm_homepage_module_tag hmt  " +
		"INNER join ods_cm_live_tag lt ON hmt.tag_id=lt.tagId " +
		"INNER JOIN dim_cmt_course_evaluation_profile cep ON lt.liveId=cep.subjectId " +
		"where hmt.module_id=? and cep.categoryId=? and lt.tagId=? " +
		"order by cep.orderfinal DESC " +
		"LIMIT 60;"

	rows, err := utils.Db.Query(sql, moduleId, categoryId, tagId)
	defer utils.ErrorCatch()
	defer rows.Close()

	if err != nil {
		panic(err)
		return
	}

	for rows.Next() {
		info := entitys.LiveOrderfinal{}
		rows.Scan(&info.LiveId, &info.TagId, &info.CategoryId, &info.Orderfinal, &info.Livetype)
		datas = append(datas, info)
	}

	return
}

//查询课程排序值
func (home TagModuleDataDao) FindAllLiveOrderfinalInfos() (datas []entitys.LiveOrderfinal) {

	sql := "SELECT  hmt.module_id as moduleId,lt.liveId,lt.tagId,cep.categoryId,cep.orderfinal,cep.livetype " +
		"from ods_cm_homepage_module_tag hmt  " +
		"INNER join ods_cm_live_tag lt ON hmt.tag_id=lt.tagId " +
		"INNER JOIN dim_cmt_course_evaluation_profile cep ON lt.liveId=cep.subjectId " +
		"order by cep.orderfinal DESC "

	rows, err := utils.Db.Query(sql)
	defer utils.ErrorCatch()
	defer rows.Close()

	if err != nil {
		panic(err)
		return
	}

	for rows.Next() {
		info := entitys.LiveOrderfinal{}
		rows.Scan(&info.ModuleId, &info.LiveId, &info.TagId, &info.CategoryId, &info.Orderfinal, &info.Livetype)
		datas = append(datas, info)
	}

	return
}

//用户学科权重
func (home TagModuleDataDao) FindUserCategoryWeights() (datas []entitys.UserCategoryWeights) {

	sql := "SELECT ucp.userId,GROUP_CONCAT(ucp.categroyId) as categroyIdStr, " +
		"GROUP_CONCAT(ucp.categroryWeights) as categroryWeightsStr " +
		"from dim_user_course_profile ucp where lastactDate is not null " +
		"GROUP BY ucp.userId"

	rows, err := utils.Db.Query(sql)
	defer utils.ErrorCatch()
	defer rows.Close()

	if err != nil {
		panic(err)
		return
	}

	for rows.Next() {

		info := entitys.UserCategoryWeights{}
		rows.Scan(&info.UserId, &info.CategroyIdStr, &info.CategroryWeightsStr)
		datas = append(datas, info)
	}

	return
}

func (home TagModuleDataDao) FindUserCategoryWeightsTest(userId int64) (datas []entitys.UserCategoryWeights) {

	sql := "SELECT ucp.userId,GROUP_CONCAT(ucp.categroyId) as categroyIdStr, " +
		"GROUP_CONCAT(ucp.categroryWeights) as categroryWeightsStr " +
		"from dim_user_course_profile ucp where ucp.userId in (" + strconv.FormatInt(userId, 10) + ")" +
		"GROUP BY ucp.userId"

	rows, err := utils.Db.Query(sql)
	defer utils.ErrorCatch()
	defer rows.Close()

	if err != nil {
		panic(err)
		return
	}

	for rows.Next() {

		info := entitys.UserCategoryWeights{}
		rows.Scan(&info.UserId, &info.CategroyIdStr, &info.CategroryWeightsStr)
		datas = append(datas, info)
	}

	return
}

//用户标签权重
func (home TagModuleDataDao) FindUserTagWeights(userId int64, tagId int64) (datas []entitys.UserTagWeights) {

	sql := "select uctp.userid,uctp.tagId,uctp.tagWeights " +
		"from dim_user_content_tag_profile uctp " +
		"where uctp.userid=? and uctp.tagId=?"

	rows, err := utils.Db.Query(sql, userId, tagId)
	defer utils.ErrorCatch()
	defer rows.Close()

	if err != nil {
		panic(err)
		return
	}

	for rows.Next() {

		info := entitys.UserTagWeights{}
		rows.Scan(&info.UserId, &info.TagId, &info.TagWeights)
		datas = append(datas, info)
	}

	return
}

//用户标签权重
func (home TagModuleDataDao) FindAllUserTagWeights() (datas []entitys.UserTagWeights) {

	sql := "select uctp.userid,uctp.tagId,uctp.tagWeights " +
		"from dim_user_content_tag_profile uctp "

	rows, err := utils.Db.Query(sql)
	defer utils.ErrorCatch()
	defer rows.Close()

	if err != nil {
		panic(err)
		return
	}

	for rows.Next() {

		info := entitys.UserTagWeights{}
		rows.Scan(&info.UserId, &info.TagId, &info.TagWeights)
		datas = append(datas, info)
	}

	return
}

//查询模块配置的课程
func (home TagModuleDataDao) FindModuleSettingLiveInfo(moduleId int64) (datas []entitys.ModuleSettingLiveInfo) {

	sql := "select hs.classId as liveId,hs.dispaly_order as dispalyOrder,cep.categoryId,cep.livetype " +
		"from ods_cm_homepage_set hs " +
		"INNER JOIN dim_cmt_course_evaluation_profile cep on hs.classId=cep.subjectId " +
		"where hs.status=1 and hs.module_id=?;"

	rows, err := utils.Db.Query(sql, moduleId)
	defer utils.ErrorCatch()
	defer rows.Close()

	if err != nil {
		panic(err)
		return
	}

	for rows.Next() {

		info := entitys.ModuleSettingLiveInfo{}
		rows.Scan(&info.LiveId, &info.DispalyOrder, &info.CategoryId, &info.Livetype)
		datas = append(datas, info)
	}

	return
}

//查询用户其他模块已经有的课程id
func (home TagModuleDataDao) FindUseLiveId(moduleId int64, userId int64) (datas []entitys.RecommendLiveInfo) {

	sql := "SELECT DISTINCT liveId " +
		"from st_cm_live_recommendation_sets " +
		"where moudleId!=? and userId=? GROUP BY liveId"

	rows, err := utils.Db.Query(sql, moduleId, userId)
	defer utils.ErrorCatch()
	defer rows.Close()

	if err != nil {
		panic(err)
		return
	}

	for rows.Next() {

		info := entitys.RecommendLiveInfo{}
		rows.Scan(&info.LiveId)
		datas = append(datas, info)
	}

	return
}

//查询数据收集结果
func (home TagModuleDataDao) FindCollectResult(moduleId int64, categoryId int64, userId int64) (datas []entitys.LiveDataCollectResult) {

	sql := "SELECT module_id,userId,category_id,live_id,max(orderfinal) as orderfinal,livetype " +
		"from cm_homepage_recommend_temple " +
		"where module_id=? and category_id=? and userId=? " +
		"GROUP BY live_id " +
		"order by orderfinal desc"

	rows, err := utils.Db.Query(sql, moduleId, categoryId, userId)
	defer utils.ErrorCatch()
	defer rows.Close()

	if err != nil {
		panic(err)
		return
	}

	for rows.Next() {

		info := entitys.LiveDataCollectResult{}
		rows.Scan(&info.ModuleId, &info.UserId, &info.CategoryId, &info.LiveId, &info.Orderfinal, &info.Livetype)
		datas = append(datas, info)
	}

	return
}

//删除数据收集结果
func (home TagModuleDataDao) DelCollectResult(moduleId int64, categoryId int64, userId int64) {

	sql := "delete from cm_homepage_recommend_temple where module_id=? and category_id=? and userId=?;"
	_, error := utils.Db.Exec(string(sql), moduleId, categoryId, userId)

	if error != nil {
		panic(error)
	}

	defer utils.ErrorCatch()
}

//添加数据收集结果
func (home TagModuleDataDao) AddCollectResult(datas []entitys.LiveDataCollectResult) {

	var sql []byte
	sql = append(sql, "insert into cm_homepage_recommend_temple(module_id,userId,category_id,tag_id,live_id,orderfinal,livetype,create_time) values "...)

	for index, info := range datas {

		if index == len(datas)-1 {
			sql = append(sql, "("+strconv.FormatInt(info.ModuleId, 10)+","+strconv.FormatInt(info.UserId, 10)+","+strconv.FormatInt(info.CategoryId, 10)+","+strconv.FormatInt(info.TagId, 10)+","+strconv.FormatInt(info.LiveId, 10)+","+strconv.FormatFloat(info.Orderfinal, 'f', 4, 64)+","+strconv.FormatInt(info.Livetype, 10)+",now());"...)
		} else {
			sql = append(sql, "("+strconv.FormatInt(info.ModuleId, 10)+","+strconv.FormatInt(info.UserId, 10)+","+strconv.FormatInt(info.CategoryId, 10)+","+strconv.FormatInt(info.TagId, 10)+","+strconv.FormatInt(info.LiveId, 10)+","+strconv.FormatFloat(info.Orderfinal, 'f', 4, 64)+","+strconv.FormatInt(info.Livetype, 10)+",now()),"...)
		}
	}

	_, error := utils.Db.Exec(string(sql))

	if error != nil {
		panic(error)
	}

	defer utils.ErrorCatch()

}

//添加推荐课程
func (home TagModuleDataDao) AddRecommendLiveResult(datas []entitys.RecommendLiveInfo) {

	var sql []byte
	sql = append(sql, "insert into st_cm_live_recommendation_sets(moudleId,userId,liveId,liveType,categoryId,orderfinal,isSysAuto,pos) values "...)

	for index, info := range datas {

		if index == len(datas)-1 {
			sql = append(sql, "("+strconv.FormatInt(info.MoudleId, 10)+","+strconv.FormatInt(info.UserId, 10)+","+strconv.FormatInt(info.LiveId, 10)+","+strconv.FormatInt(info.LiveType, 10)+","+strconv.FormatInt(info.CategoryId, 10)+","+strconv.FormatFloat(info.Orderfinal, 'f', 4, 64)+","+strconv.FormatInt(info.IsSysAuto, 10)+","+strconv.FormatInt(info.Pos, 10)+");"...)
		} else {
			sql = append(sql, "("+strconv.FormatInt(info.MoudleId, 10)+","+strconv.FormatInt(info.UserId, 10)+","+strconv.FormatInt(info.LiveId, 10)+","+strconv.FormatInt(info.LiveType, 10)+","+strconv.FormatInt(info.CategoryId, 10)+","+strconv.FormatFloat(info.Orderfinal, 'f', 4, 64)+","+strconv.FormatInt(info.IsSysAuto, 10)+","+strconv.FormatInt(info.Pos, 10)+"),"...)
		}
	}

	_, error := utils.Db.Exec(string(sql))

	if error != nil {
		panic(error)
	}

	defer utils.ErrorCatch()
}

//获取学科数据
func (home TagModuleDataDao) FindCategoryDatas() (datas []entitys.CategoryInfo) {

	sql := "SELECT GROUP_CONCAT(id) as categoryIdStr,GROUP_CONCAT(category) as categoryNameStr " +
		"from ods_cm_category c " +
		"where c.available=1;"

	rows, err := utils.Db.Query(sql)
	defer utils.ErrorCatch()
	defer rows.Close()

	if err != nil {
		panic(err)
		return
	}

	for rows.Next() {

		info := entitys.CategoryInfo{}
		rows.Scan(&info.CategoryIdStr, &info.CategoryNameStr)
		datas = append(datas, info)
	}

	return
}

func (home TagModuleDataDao) ClearDatas(isallClear bool) {

	sql := "delete from st_cm_live_recommendation_sets;"
	_, error := utils.Db.Exec(string(sql))

	if error != nil {
		panic(error)
	}

	if isallClear {
		sql = "DELETE from cm_homepage_recommend_temple;"
		_, error = utils.Db.Exec(string(sql))
	}

	defer utils.ErrorCatch()
}

//查询收集标记
func (home TagModuleDataDao) FindCollectSign() (signCount int64) {

	sql := "select count(*) as number " +
		"from cm_homepage_recommend_temple ht " +
		"where ht.module_id=-888"

	rows, err := utils.Db.Query(sql)
	defer utils.ErrorCatch()
	defer rows.Close()

	if err != nil {
		panic(err)
		return
	}

	for rows.Next() {

		rows.Scan(&signCount)
	}

	return
}

//查询模块已经收集完成的userid
func (home TagModuleDataDao) FindCollectModuleUserIds(moduleId int64) (userIds []int64) {

	sql := "select ht.module_id,ht.userId " +
		"from cm_homepage_recommend_temple ht " +
		"where ht.module_id=? " +
		"GROUP BY ht.userId;"

	rows, err := utils.Db.Query(sql, moduleId)
	defer utils.ErrorCatch()
	defer rows.Close()

	if err != nil {
		panic(err)
		return
	}

	for rows.Next() {

		userId := int64(0)
		m := int64(0)
		rows.Scan(&m, &userId)
		userIds = append(userIds, userId)
	}

	return
}
