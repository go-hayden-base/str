package str

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"
)

// MD5 reutnr string's md5
func MD5(source string) string {
	if len(source) == 0 {
		return ""
	}

	md5Ctx := md5.New()
	md5Ctx.Write([]byte(source))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

const (
	OptionRandomStringNum   = 0 // 纯数字
	OptionRandomStringLower = 1 // 小写字母
	OptionRandomStringUpper = 2 // 大写字母
	OptionRandomStringAll   = 3 // 数字、大小写字母
)

// RandomMD5String return a md5 random string
func RandomMD5String() string {
	prefix := RandomString(64, OptionRandomStringAll)
	suffix := RandomString(64, OptionRandomStringAll)
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	return MD5(prefix + timestamp + suffix)
}

// RandomString return a random string
func RandomString(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}
