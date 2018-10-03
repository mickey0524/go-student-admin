package util

import (
	"strconv"
)

// ToInt8 string -> int8
func ToInt8(s string) int8 {
	res, _ := strconv.ParseInt(s, 10, 8)
	return int8(res)
}

// ToInt64 string -> int64
func ToInt64(s string) int64 {
	res, _ := strconv.ParseInt(s, 10, 64)
	return res
}
