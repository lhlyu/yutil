package yfunc

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/lhlyu/logger"
)

// @编码
func AnyEncode(v interface{}) string{
	if v == nil{
		logger.Error("AnyEncode","v is nil")
		return ""
	}
	byts,err := json.Marshal(v)
	if err != nil{
		logger.Error("AnyEncode","json marshal is err : ",err)
		return ""
	}
	return base64.StdEncoding.EncodeToString(byts)
}

// @解码
func AnyDecode(s string,v interface{}) error{
	if v == nil{
		logger.Error("AnyDecode","v is nil")
		return errors.New("v is nil")
	}
	byts,err := base64.URLEncoding.DecodeString(s)
	if err != nil{
		logger.Error("AnyDecode","decodeString is err : ",err)
		return err
	}
	err = json.Unmarshal(byts,v)
	if err != nil{
		logger.Error("AnyDecode","json unmarshal is err : ",err)
		return err
	}
	return nil
}


