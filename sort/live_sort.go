package sort

//author 赵伦翔
import (
	"fmt"
	Sort "sort"
)

//直播排序大厅

type SortLevel int

//以下变量是从上到下按照产品优先级排序，如果后续排序优先级有变化  只需要调整先后既可
//以下是排序的优先级 配置到下面既可
const (
	min                     SortLevel = iota
	IsLiveLinkSortLevel               //是否在麦
	StickSortLevel                    //置顶
	LiveLinkSortLevel                 //在麦人数
	TurnplateSortLevel                //转盘
	LiveRewardSortLevel               //送礼
	LiveRoomOnlineSortLevel           //房间在线人数
	//todo 此处增加其他排序维度

	max
)

//排序维度
type SortCell struct {
	Level SortLevel
	Num   int64 //权重值   是（1）/否（0）  数量 同样适用
}

type HallSort interface {
	SortFeed() []SortCell
}

//按照权重排序
func SortLiveHall(slice []HallSort) {
	Sort.Slice(slice, func(i, j int) bool {
		iList := slice[i].SortFeed()
		jList := slice[j].SortFeed()
		Sort.Slice(iList, func(k, m int) bool {
			return iList[k].Level < iList[m].Level
		})
		Sort.Slice(jList, func(k, m int) bool {
			return jList[k].Level < jList[m].Level
		})

		if len(iList) > 0 && len(jList) > 0 {
			//同时都有排序维度的情况
			maxLen := len(iList)
			if maxLen > len(jList) {
				maxLen = len(jList)
			}
			for index := 0; index < maxLen; index++ {
				if iList[index].Level == jList[index].Level && iList[index].Num == jList[index].Num {
					//排序级别相同 且排序权重也相同
					continue
				}
				if iList[index].Level == jList[index].Level {
					fmt.Printf("index_i_num=%d,j_index_num:%d\n", iList[index].Num, jList[index].Num)
					//排序级别相同的情况
					return iList[index].Num > jList[index].Num
				} else {
					//级别越小 越靠前
					fmt.Printf("index_i=%d,j_index:%d\n", iList[index].Level, jList[index].Level)
					return iList[index].Level < jList[index].Level
				}
			}
		} else {
			//其他有一个没有排序维度的情况
			return len(iList) > len(jList)
		}
		return true

	})
}
