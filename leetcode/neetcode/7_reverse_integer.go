package neetcode

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	var (
		max       = math.MaxInt32
		min       = math.MinInt32
		xStr      = strconv.Itoa(x)
		resultStr string
	)

	for i := len(xStr) - 1; i >= 1; i-- {
		resultStr += string(xStr[i])
	}
	if string(xStr[0]) == "-" {
		resultStr = "-" + resultStr
	} else {
		resultStr += string(xStr[0])
	}

	resultI, _ := strconv.Atoi(resultStr)
	if resultI > max || resultI < min {
		return 0
	}
	return resultI
}
