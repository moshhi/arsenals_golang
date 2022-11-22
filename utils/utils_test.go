package utils

import (
	"fmt"
	"testing"
)

func TestStrPager(t *testing.T) {
	strArr := []string{"word", "name", "age", "jack", "rose", "trump"}
	strPager := StrPager{
		Arr:      strArr,
		PageSize: 2,
	}
	pageCnt, _ := strPager.PageCnt()
	for i := int32(0); i < pageCnt; i++ {
		pageItems, _ := strPager.PageItems(i)
		fmt.Println(pageItems)
	}
}

func TestRmDupNumArr(t *testing.T) {
	arr := []int64{1, 1, 2, 2, 3, 4, 4, 5, 5}
	r := RmDupNumArr(arr)
	fmt.Println(r)
}

func TestJoinNumArr(t *testing.T) {
	arr := []int64{1, 1, 2, 2, 3, 4, 4, 5, 5}
	r := JoinNumArr(arr, ",")
	fmt.Println(r)
}
