package yfunc

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/lhlyu/logger"
)

// @MD5加密
func Md5(value interface{}) string {
	byts,err := json.Marshal(value)
	if err != nil{
		logger.Error("Md5 marshal error = ",err)
		return ""
	}
	h := md5.New()
	h.Write(byts)
	return hex.EncodeToString(h.Sum(nil))
}