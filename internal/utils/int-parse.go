package utils

import (
	"strconv"
)

func IntParse(s string) int {
	v, _ := strconv.ParseUint(s, 16, 0)
	return int(v)
}
