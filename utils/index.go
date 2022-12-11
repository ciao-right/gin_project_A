package utils

import (
	"math/rand"
	"time"
)

func RandString(num int) string {
	rand.Seed(time.Now().Unix())
	letters := []byte("asdfasdhajksbhvajdsifuaiosdfaksndfkl")
	res := make([]byte, num)
	for i := range res {
		res[i] = letters[rand.Intn(len(letters))]
	}
	return string(res)
}

func IsEmptyString(str string) bool {
	return len(str) == 0
}
