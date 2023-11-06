package crypto

import (
	"crypto/sha256"
	"encoding/hex"
)

func SHA256(str string) string {
	hash := sha256.Sum256([]byte(str))
	return hex.EncodeToString(hash[:])
}

func SHA256Bytes(byteData []byte) string {
	hash := sha256.Sum256(byteData)
	return hex.EncodeToString(hash[:])
}
