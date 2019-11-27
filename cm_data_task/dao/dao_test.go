package dao

import (
	"fmt"
	"testing"
)

func TestHomeInfo(t *testing.T) {

	homedao := TagModuleDataDao{}

	fmt.Println(homedao.FindCollectSign())
}

func festGuessDao(t *testing.T) {

	g := GuessLikeModuleDataDao{}
	//fmt.Println(g.FindUserId())
	//uuids := g.FindUserMaxPlayPctLiveUUID(int64(2499739))
	//fmt.Println(g.FindOtherUserMaxPlayPctLiveInfo(int64(2499739), uuids))
	fmt.Println(g.FindUserPlayPctLiveUUId(int64(2499739)))
}
