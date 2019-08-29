package sort

import (
	"encoding/json"
	"testing"
)

func TestSortLiveHall(t *testing.T) {
	source := Data()

	sourceBytes, _ := json.Marshal(source)
	t.Logf("source=%v", string(sourceBytes))
	SortLiveHall(source)
	resultBytes, _ := json.Marshal(source)
	t.Logf("result=%v", string(resultBytes))

}

type TestResult struct {
	IsLink      int64 `json:"is_link"`
	Top         int64 `json:"top"`           //是否置顶
	LinkRoomNum int64 `json:"link_room_num"` //连麦人数
	IsTurnplate int64 `json:"is_turnplate"`  //是否转盘
	IsReward    int64 `json:"is_reward"`     //是否送礼
	OnlineNum   int64 `json:"online_num"`    //房间人数
}

func (me TestResult) SortFeed() []SortCell {
	list := make([]SortCell, 0)
	list = append(list, SortCell{
		Level: StickSortLevel,
		Num:   me.Top,
	})

	if me.LinkRoomNum > 0 {
		list = append(list, SortCell{
			Level: LiveLinkSortLevel,
			Num:   me.LinkRoomNum,
		})
	}

	list = append(list, SortCell{
		Level: TurnplateSortLevel,
		Num:   me.IsTurnplate,
	})

	list = append(list, SortCell{
		Level: LiveRewardSortLevel,
		Num:   me.IsReward,
	})
	if me.OnlineNum > 0 {
		list = append(list, SortCell{
			Level: LiveRoomOnlineSortLevel,
			Num:   me.OnlineNum,
		})
	}
	if me.IsLink > 0 {
		list = append(list, SortCell{
			Level: IsLiveLinkSortLevel,
			Num:   me.IsLink,
		})
	}
	return list
}

func Data() []HallSort {
	testData := make([]HallSort, 0)
	testData = append(testData, TestResult{
		Top:         1,
		LinkRoomNum: 3,
		IsTurnplate: 1,
		IsReward:    1,
		OnlineNum:   100,
	})

	testData = append(testData, TestResult{
		Top:         0,
		LinkRoomNum: 3,
		IsTurnplate: 1,
		IsReward:    1,
		OnlineNum:   100,
	})

	testData = append(testData, TestResult{
		Top:         1,
		LinkRoomNum: 4,
		IsTurnplate: 1,
		IsReward:    1,
		OnlineNum:   100,
	})

	testData = append(testData, TestResult{
		Top:         1,
		LinkRoomNum: 2,
		IsTurnplate: 1,
		IsReward:    1,
		OnlineNum:   100,
	})

	testData = append(testData, TestResult{
		Top:         1,
		LinkRoomNum: 3,
		IsTurnplate: 0,
		IsReward:    1,
		OnlineNum:   100,
	})

	testData = append(testData, TestResult{
		Top:         1,
		LinkRoomNum: 3,
		IsTurnplate: 1,
		IsReward:    0,
		OnlineNum:   100,
	})

	testData = append(testData, TestResult{
		Top:         1,
		LinkRoomNum: 3,
		IsTurnplate: 1,
		IsReward:    0,
		OnlineNum:   100,
		IsLink:      1,
	})
	testData = append(testData, TestResult{
		Top:         0,
		LinkRoomNum: 3,
		IsTurnplate: 1,
		IsReward:    0,
		OnlineNum:   100,
		IsLink:      1,
	})

	testData = append(testData, TestResult{
		Top:         1,
		LinkRoomNum: 3,
		IsTurnplate: 1,
		IsReward:    1,
		OnlineNum:   100,
	})

	testData = append(testData, TestResult{
		Top:         0,
		LinkRoomNum: 3,
		IsTurnplate: 1,
		IsReward:    1,
		OnlineNum:   100,
	})

	testData = append(testData, TestResult{
		Top:         1,
		LinkRoomNum: 3,
		IsTurnplate: 1,
		IsReward:    1,
		OnlineNum:   99,
	})

	testData = append(testData, TestResult{
		Top:         0,
		LinkRoomNum: 3,
		IsTurnplate: 1,
		IsReward:    1,
		OnlineNum:   199,
	})

	testData = append(testData, TestResult{
		Top:         1,
		LinkRoomNum: 3,
		IsTurnplate: 1,
		IsReward:    1,
		OnlineNum:   101,
	})
	testData = append(testData, TestResult{
		Top:         0,
		LinkRoomNum: 5,
		IsTurnplate: 1,
		IsReward:    1,
		OnlineNum:   101,
	})
	testData = append(testData, TestResult{
		Top:         0,
		LinkRoomNum: 6,
		IsTurnplate: 1,
		IsReward:    1,
		OnlineNum:   101,
	})
	return testData
}
