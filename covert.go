package yutil

import (
	"fmt"
	"reflect"
	"strconv"
)

// v1.0.0: 将任意切片转成[]interface{}
func ConvertToInterface(slice interface{}) []interface{} {
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

// v1.0.0: 将字符串切片转成[]interface{}
func ConvertStrToInterface(slice []string) []interface{} {
	if len(slice) == 0 {
		return nil
	}
	params := make([]interface{}, len(slice))
	for i, v := range slice {
		params[i] = v
	}
	return params
}

// v1.0.0: 将整形切片转成[]interface{}
func ConvertIntToInterface(slice []int) []interface{} {
	if len(slice) == 0 {
		return nil
	}
	params := make([]interface{}, len(slice))
	for i, v := range slice {
		params[i] = v
	}
	return params
}

// v1.0.0: 将任意类型转string
func ConvertAnyToStr(v interface{}) string {
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

// v1.0.0: 将任意类型转int
func ConvertAnyToInt(v interface{}) int {
	if v == nil {
		return 0
	}
	switch d := v.(type) {
	case string:
		val, err := strconv.Atoi(d)
		_conf.log("conver", "ConvertAnyToInt", err)
		return val
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return int(reflect.ValueOf(v).Int())
	case []byte:
		val, err := strconv.Atoi(string(d))
		_conf.log("conver", "ConvertAnyToInt", err)
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

// v1.0.0: 将任意类型转interface
func ConvertAnyToInterface(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	return reflect.ValueOf(v).Interface()
}
