package commons

import (
	"strconv"
	"time"
	"math/rand"
	"crypto/md5"
	"encoding/hex"
)

/**
产生一个随机文件名
 */
func GenerateFileName(fileExt string) string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(int(time.Now().UnixNano())) + strconv.Itoa(rand.Intn(1000)) + fileExt
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

/**
时间类型
 */
func GetTimeLayout(layout string) string {
	switch layout {
	case "day":
		return "2019-04-06"
		break
	case "hour":
		return "2019-10-11 03:02"
		break
	case "min":
		return "2006-01-02 15:04:02"
		break
	}
	return "2006-01-02 15:04:02"
}

// 返回md5值
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}
