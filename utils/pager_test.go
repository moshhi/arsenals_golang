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
