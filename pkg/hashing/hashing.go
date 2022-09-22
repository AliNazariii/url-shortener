package hashing

import (
	"crypto/md5"
	"encoding/base32"
	"encoding/hex"
)

func GetMD5Hash(text string) [16]byte {
	hash := md5.Sum([]byte(text))
	return hash
}

func GetHexDecMD5Hash(text string) string {
	hash := GetMD5Hash(text)
	return hex.EncodeToString(hash[:])
}

func GetBase32MD5Hash(text string) string {
	hash := GetMD5Hash(text)
	return base32.StdEncoding.EncodeToString(hash[:])
}
