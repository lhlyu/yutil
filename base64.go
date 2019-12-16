package yutil

import (
	"encoding/base64"
	"encoding/json"
)

// 编码:将Base64[]byte转成字符串
func Base64EncodeBytesToStr(bts []byte) string {
	result := base64.StdEncoding.EncodeToString(bts)
	return result
}

// 编码:将普通string转成base64字符串
func Base64EncodeStrToStr(data string) string {
	return Base64EncodeBytesToStr([]byte(data))
}

// 编码:将一个对象转成bas64字符串
func Base64EncodeObjToStr(v interface{}) string {
	bytes, err := json.Marshal(v)
	_conf.log("base64","Base64EncodeObjToStr",err)
	return Base64EncodeBytesToStr(bytes)
}


// 解码:将base64字符串解码成[]byte
func Base64DecodeStrToBytes(data string) []byte {
	bytes, err := base64.StdEncoding.DecodeString(data)
	_conf.log("base64","Base64DecodeStrToBytes",err)
	return bytes
}

// 解码:将base64字符串解码成string
func Base64DecodeStrToStr(data string) string {
	bts := Base64DecodeStrToBytes(data)
	return string(bts)
}

// 解码:将base64字符串反序列化为对象
func Base64DecodeStrToObj(data string, v interface{}) {
	bts := Base64DecodeStrToBytes(data)
	err := json.Unmarshal(bts, v)
	_conf.log("base64","Base64DecodeStrToObj",err)
}