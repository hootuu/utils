package strs

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	hashValue := hash.Sum(nil)
	md5Str := hex.EncodeToString(hashValue)
	return md5Str
}
