package helper

import (
	"crypto/md5"
	"encoding/hex"
)

func GetHash(value string) string {
	hash := md5.Sum([]byte(value))
	return hex.EncodeToString(hash[:])
}
