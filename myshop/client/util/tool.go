package util

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	ss := hex.EncodeToString(h.Sum(nil))
	return ss
}