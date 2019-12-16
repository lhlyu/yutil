package yutil

import (
	"crypto/md5"
	"encoding/hex"
)

// md5加密
func Md5(data string) string {
	hash := md5.New()
	_, err := hash.Write([]byte(data))
	_conf.log("md5", "Md5", err)
	cipherStr := hash.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// md5 加密附带盐值
func Md5WithSalt(data, salt string) string {
	hash := md5.New()
	_, err := hash.Write([]byte(data + salt))
	_conf.log("md5", "Md5WithSalt", err)
	cipherStr := hash.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
