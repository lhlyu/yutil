package yutil

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
)

// 字符串base64编码
func (yBase64) Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

// 对象进行base64编码
func (yBase64) Encode4Obj(data interface{}) string {
	bytes, err := json.Marshal(data)
	_arc.log("Base64", "Encode4Obj", err)
	return base64.StdEncoding.EncodeToString([]byte(bytes))
}

// base64字符串解码
func (yBase64) Decode(data string) string {
	bytes, err := base64.StdEncoding.DecodeString(data)
	_arc.log("Base64", "Decode", err)
	return string(bytes)
}

// base64字符串反序列化对象
func (self yBase64) Decode2Obj(data string, v interface{}) {
	bts := self.Decode(data)
	err := json.Unmarshal([]byte(bts), v)
	_arc.log("Base64", "Decode4Obj", err)
}

// ---------------------------------------------------------

// md5加密
func (yMd5) Encode(data string) string {
	hash := md5.New()
	_, err := hash.Write([]byte(data))
	_arc.log("Md5", "Encode", err)
	cipherStr := hash.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// md5 加密附带盐值
func (yMd5) EncodeWithSalt(data, salt string) string {
	hash := md5.New()
	_, err := hash.Write([]byte(data + salt))
	_arc.log("Md5", "EncodeWithSalt", err)
	cipherStr := hash.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// ------------------------------------------------------------

// sha1 加密
func (ySha1) Encode(data string) string {
	hash := sha1.New()
	hash.Write([]byte(data))
	cipherStr := hash.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// sha1 加密附带盐值
func (ySha1) EncodeWithSalt(data, salt string) string {
	hash := sha1.New()
	hash.Write([]byte(data + salt))
	cipherStr := hash.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
