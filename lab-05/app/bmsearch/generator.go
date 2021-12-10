package bmsearch

import (
	"math/rand"
	"time"
)


func GenerateRune(size int) []rune {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
    b := make([]rune, size)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return b
}

func GeneratePattern(src []rune, size int) []rune {
	rand.Seed(time.Now().UnixNano())
	start := rand.Intn(len(src) - size)
	length := len(src)

	return src[start : length - 1]
}