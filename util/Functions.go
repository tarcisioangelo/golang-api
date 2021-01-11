package util

import (
	"crypto/md5"
	"encoding/hex"
)

// GetMD5Hash - Gerando Hash MD5
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
