package str

import (
	"testing"
)

func TestMD5(t *testing.T) {
	md5 := MD5("abc123!@#")
	if md5 != "65af2bf5f5c7d6802d01bf967917e0cd" {
		t.Errorf("TestMD5 Fail! (md5: %s", md5)
	}
}
