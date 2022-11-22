package utils

import (
	"strconv"
	"strings"
)

func RmDupNumArr(numbers []int64) []int64 {
	numSet := Int64Set{}
	for _, i := range numbers {
		numSet.Add(i)
	}
	return numSet.Items()
}

func JoinNumArr(numbers []int64, sep string) string {
	strArr := make([]string, 0, len(numbers))
	for _, num := range numbers {
		s := strconv.FormatInt(num, 10)
		strArr = append(strArr, s)
	}
	return strings.Join(strArr, sep)
}
