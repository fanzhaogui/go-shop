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

/**
生成一个唯一的ID
@return int
 */
func GenerateId() int {
	rand.Seed(time.Now().UnixNano())
	id, _ := strconv.Atoi(strconv.Itoa(rand.Intn(10000)) + strconv.Itoa(int(time.Now().Unix())))
	return id
}