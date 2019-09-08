package commons

import (
	"strconv"
	"time"
	"math/rand"
)

/**
产生一个随机文件名
 */
func GenerateFileName(fileExt string) string {
	rand.Seed(time.Now().UnixNano())
	return  strconv.Itoa(int(time.Now().UnixNano())) + strconv.Itoa(rand.Intn(1000))+ fileExt
}
