package yutil

import (
	"fmt"
	"reflect"
	"strconv"
)

// 将任意切片转成[]interface{}
func (yConvert) ToSlinceInterface(slice interface{}) []interface{} {
	val := reflect.ValueOf(slice)
	if val.Kind() != reflect.Slice {
		return nil
	}
	sliceLen := val.Len()
	if sliceLen == 0 {
		return nil
	}
	params := make([]interface{}, sliceLen)
	for i := 0; i < sliceLen; i++ {
		params[i] = val.Index(i).Interface()
	}
	return params
}

// 将任意类型转string
func (yConvert) ToString(v interface{}) string {
	if v == nil {
		return ""
	}
	switch d := v.(type) {
	case string:
		return d
	case int, int8, int16, int32, int64:
		return strconv.FormatInt(reflect.ValueOf(v).Int(), 10)
	case uint, uint8, uint16, uint32, uint64:
		return strconv.FormatUint(reflect.ValueOf(v).Uint(), 10)
	case []byte:
		return string(d)
	case float32, float64:
		return strconv.FormatFloat(reflect.ValueOf(v).Float(), 'f', -1, 64)
	case bool:
		return strconv.FormatBool(d)
	default:
		return fmt.Sprint(v)
	}
}

// 将任意类型转int
func (yConvert) ToInt(v interface{}) int {
	if v == nil {
		return 0
	}
	switch d := v.(type) {
	case string:
		val, err := strconv.Atoi(d)
		_arc.log("Conver", "ToInt", err)
		return val
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return int(reflect.ValueOf(v).Int())
	case []byte:
		val, err := strconv.Atoi(string(d))
		_arc.log("Conver", "ToInt", err)
		return val
	case float32, float64:
		return int(reflect.ValueOf(v).Float())
	case bool:
		if d {
			return 0
		} else {
			return 1
		}
	case complex64:
		return int(real(d))
	case complex128:
		return int(real(d))
	default:
		return int(reflect.ValueOf(v).Float())
	}
}

// 将任意类型转interface
func (yConvert) ToInterface(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return reflect.ValueOf(v).Interface()
}

// 将Int切片转成字符串切片
func (yConvert) IntSliceToStrSlice(v []int) []string {
	var arr []string
	for _, v := range v {
		arr = append(arr, strconv.Itoa(v))
	}
	return arr
}

// 将字符串切片转成int切片
func (yConvert) StrSliceToIntSlice(v []string) []int {
	var arr []int
	for _, v := range v {
		a, err := strconv.Atoi(v)
		_arc.log("Conver", "StrSliceToIntSlice", err)
		arr = append(arr, a)
	}
	return arr
}
