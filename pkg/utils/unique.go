package utils

import (
	"github.com/google/uuid"
	"math/rand"
	"unicode"
)

func MakeUniqueUUID() (result string) {
	for i := 0; i < 2; i++ {
		result += uuid.New().String()
	}
	for i, r := range result {
		if unicode.IsLetter(r) {
			if rand.Intn(100-0)+0 >= 50 {
				result = replaceAtIndex(result, unicode.ToUpper(r), i)
			}
		}
	}
	return
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
