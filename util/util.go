package util

import (
	"math/rand"
	"time"
)

// 随机生成用户名
func RandomString(random int) string {
	letters := []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
	result := make([]byte, random, random)
	rand.Seed(time.Now().UnixNano())
	for i, _ := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
