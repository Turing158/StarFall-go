package util

import "math/rand"

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
