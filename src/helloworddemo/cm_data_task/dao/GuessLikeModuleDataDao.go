package dao

import (
	"cm_data_task/entitys"
	"cm_data_task/utils"
	"strconv"
)

type GuessLikeModuleDataDao struct{}

/**
获取观看课程用户id
**/
func (guess GuessLikeModuleDataDao) FindUserId() (datas []int64) {

	sql := "select DISTINCT d.userid " +
		"from ods_cm_event_course_play_view_day d "

	rows, err := utils.Db.Query(sql)
	defer utils.ErrorCatch()
	defer rows.Close()

	if err != nil {
		panic(err)
		return
	}

	for rows.Next() {

		id := int64(0)
		rows.Scan(&id)
		datas = append(datas, id)
	}

	return
}

func (guess GuessLikeModuleDataDao) FindUserIdTest() (datas []int64) {

	sql := "select DISTINCT d.userid " +
		"from ods_cm_event_course_play_view_day d where d.userid in (2538963)"

	rows, err := utils.Db.Query(sql)
	defer utils.ErrorCatch()
	defer rows.Close()

	if err != nil {
		panic(err)
		return
	}

	for rows.Next() {

		id := int64(0)
		rows.Scan(&id)
		datas = append(datas, id)
	}

	return
}

/**
获取用户完播率最大的课程uuid
**/
func (guess GuessLikeModuleDataDao) FindUserMaxPlayPctLiveUUID(userId int64) (datas []string) {

	sql := "select d.liveUUID,max(d.playPct) as pct,d.userid " +
		"from ods_cm_event_course_play_view_day d " +
		"where d.userid=? " +
		"GROUP BY d.liveUUID " +
		"order by d.playPct desc " +
		"LIMIT 5; "

	rows, err := utils.Db.Query(sql, userId)
	defer utils.ErrorCatch()
	defer rows.Close()

	if err != nil {
		panic(err)
		return
	}

	for rows.Next() {

		uuid := ""
		pct := ""
		userId := int64(0)
		rows.Scan(&uuid, &pct, &userId)
		datas = append(datas, uuid)
	}

	return
}

//获取观看过这些课程的用户id
func (guess GuessLikeModuleDataDao) FindOtherUserId(userId int64, liveUUIDs []string) (datas []int64) {

	var sql []byte
	sql = append(sql, " select DISTINCT d.userid"...)
	sql = append(sql, " from ods_cm_event_course_play_view_day d "...)
	sql = append(sql, " where d.userid!=? and d.liveUUID in ("...)

	for index, uuid := range liveUUIDs {

		if index == len(liveUUIDs)-1 {
			sql = append(sql, "'"+uuid+"') "...)
		} else {
			sql = append(sql, "'"+uuid+"',"...)
		}
	}

	sql = append(sql, " LIMIT 50"...)

	rows, err := utils.Db.Query(string(sql), userId)
	defer utils.ErrorCatch()
	defer rows.Close()

	if err != nil {
		panic(err)
		return
	}

	for rows.Next() {

		var uid int64 = 0
		rows.Scan(&uid)
		datas = append(datas, uid)
	}

	return
}

//获取其他用户完播率最大的课程uuid
func (guess GuessLikeModuleDataDao) FindOtherUserMaxPlayPctLiveInfo(userIds []int64) (datas []entitys.LivePlayPctInfo) {

	var sql []byte

	sql = append(sql, "select c.id as liveId,ct.liveUUID,ct.pct "...)
	sql = append(sql, "from ods_cm_live_classroom c INNER join  ( "...)
	sql = append(sql, "select  od.userid,od.liveUUID,max(od.playPct) as pct "...)
	sql = append(sql, "from ods_cm_event_course_play_view_day od  "...)
	sql = append(sql, "where od.userid in ( "...)

	for index, uid := range userIds {

		if index == len(userIds)-1 {
			sql = append(sql, ""+strconv.FormatInt(uid, 10)+") "...)
		} else {
			sql = append(sql, ""+strconv.FormatInt(uid, 10)+","...)
		}
	}

	sql = append(sql, "GROUP BY od.liveUUID ORDER BY pct desc LIMIT 50 "...)
	sql = append(sql, ") as ct "...)
	sql = append(sql, "on ct.liveUUID=c.liveUUID; "...)

	rows, err := utils.Db.Query(string(sql))
	defer utils.ErrorCatch()
	defer rows.Close()

	if err != nil {
		panic(err)
		return
	}

	for rows.Next() {

		info := entitys.LivePlayPctInfo{}
		rows.Scan(&info.LiveId, &info.LiveUUID, &info.Pct)
		datas = append(datas, info)
	}

	return

}

//获取用户完播率可以uuid
func (guess GuessLikeModuleDataDao) FindUserPlayPctLiveUUId(userId int64) (datas []string) {

	sql := "select DISTINCT d.liveUUID " +
		"from ods_cm_event_course_play_view_day d " +
		"where d.userid=?"

	rows, err := utils.Db.Query(sql, userId)
	defer utils.ErrorCatch()
	defer rows.Close()

	if err != nil {
		panic(err)
		return
	}

	for rows.Next() {

		uuid := ""
		rows.Scan(&uuid)
		datas = append(datas, uuid)
	}

	return
}

//获取课程的排序值信息
func (guess GuessLikeModuleDataDao) FindLiveOrderfinalInfo(userId int64, moduleId int64, liveIds []int64) (datas []entitys.RecommendLiveInfo) {

	var sql []byte
	sql = append(sql, "select cep.subjectId as liveId,cep.categoryId,cep.orderfinal,cep.livetype "...)
	sql = append(sql, "from dim_cmt_course_evaluation_profile cep "...)
	sql = append(sql, "WHERE cep.subjectId in ( "...)

	for index, liveId := range liveIds {

		if index == len(liveIds)-1 {
			sql = append(sql, ""+strconv.FormatInt(liveId, 10)+")"...)
		} else {
			sql = append(sql, ""+strconv.FormatInt(liveId, 10)+","...)
		}
	}

	sql = append(sql, "ORDER BY cep.orderfinal DESC  "...)
	sql = append(sql, "LIMIT 20 "...)

	rows, err := utils.Db.Query(string(sql))
	defer utils.ErrorCatch()
	defer rows.Close()

	if err != nil {
		panic(err)
		return
	}

	for rows.Next() {

		info := entitys.RecommendLiveInfo{}
		info.UserId = userId
		info.MoudleId = moduleId
		info.IsSysAuto = 0
		info.Pos = 18
		rows.Scan(&info.LiveId, &info.CategoryId, &info.Orderfinal, &info.LiveType)
		datas = append(datas, info)
	}

	return
}
