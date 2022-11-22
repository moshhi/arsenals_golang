package utils

import "fmt"

type StrPager struct {
	Arr      []string
	PageSize int32
}

func (pager *StrPager) PageCnt() (pageCnt int32, err error) {
	if len(pager.Arr) == 0 {
		err = fmt.Errorf("data arr length is 0")
		return pageCnt, err
	}
	if pager.PageSize <= 0 {
		err = fmt.Errorf("page size must gt 0")
		return pageCnt, err
	}
	itemCnt := int32(len(pager.Arr))
	if itemCnt%pager.PageSize == 0 {
		pageCnt = itemCnt / pager.PageSize
	} else {
		pageCnt = itemCnt/pager.PageSize + 1
	}
	return pageCnt, nil
}

func (pager *StrPager) PageItems(pageNum int32) (pageItems []string, err error) {
	if pageNum < 0 {
		err = fmt.Errorf("pageNum must gte 0")
		return pageItems, err
	}
	itemCnt := int32(len(pager.Arr))
	cursor, maxCursor := pager.PageSize*pageNum, itemCnt-1
	if cursor > maxCursor {
		return pageItems, nil
	}
	if pager.PageSize > int32(len(pager.Arr[cursor:])) {
		return pager.Arr[cursor:], nil
	}
	pageItems = pager.Arr[cursor : cursor+pager.PageSize]
	return pageItems, nil
}
