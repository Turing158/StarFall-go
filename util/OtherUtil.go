package util

import (
	"math/rand"
	"time"
)

const UpperStr = "ABCDEFGHIJCKLMNOPQRSTUVWXYZ"
const lowerStr = "abcdefghijklmnopqrstuvwxyz"
const numStr = "0123456789"

func RandomStr(length int) string {
	str := UpperStr + numStr
	b := make([]byte, length)
	for i := range b {
		b[i] = str[rand.Intn(len(str))]
	}
	return string(b)
}

func IntToBool(i int) bool {
	if i > 0 {
		return true
	}
	return false
}

func Int64ToBool(i int64) bool {
	return IntToBool(int(i))
}

func IsContinualDate(date1, date2 time.Time) bool {
	//原理：通过time.Sub函数，计算两个时间之间的秒数，然后
	return date1.Sub(date2)/(24*time.Hour) == 1
}
