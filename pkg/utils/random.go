package utils

import (
	"math/rand"
	"strconv"
)

func MakeConfirmationCode() (result string) { // TODO: оптимизировать?
	for i := 0; i < 4; i++ {
		result += strconv.Itoa(rand.Intn(9-0) + 0)
	}
	return
}

func RandomInRightOutRangeInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func RandomInRangeInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}
