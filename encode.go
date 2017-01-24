package str

import (
	"crypto/md5"
	"encoding/hex"
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
