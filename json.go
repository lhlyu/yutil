package yutil

import "encoding/json"

// 将对象转成json字符串
func (yJson) Marshal(v interface{}) string {
	bts, err := json.Marshal(v)
	_arc.log("Json", "Encode", err)
	return string(bts)
}

// 将字符串反序列化为对象
func (yJson) Unmarshal(s string, v interface{}) {
	err := json.Unmarshal([]byte(s), v)
	_arc.log("Json", "Decode", err)
}
