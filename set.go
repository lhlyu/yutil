package yutil

import (
	"fmt"
	"reflect"
	"strings"
)

// v1.0.3: Set
type Set struct {
	m  map[interface{}]struct{}
	tp string // 限制类型，默认不限制
}

type Seter interface {
	Clear()    // 清空
	Size() int // 尺寸

	Add(v ...interface{}) // 添加
	Del(v ...interface{}) // 删除

	Exist(key interface{}) bool // 是否存在
	Equals(set *Set) bool       // 判断两个集合是否相等
	Empty() bool                // 是否为空

	Union(set *Set) *Set        // 并集
	Intersection(set *Set) *Set // 交集
	Difference(set *Set) *Set   // 差集

	ToArray() []interface{} // 转数组
	ToIntArray() []int      // 转int数组
	ToStrArray() []string   // 转string数组
	Print()                 // 打印
}

// 新建一个Set
func NewSet() *Set {
	return &Set{
		m: make(map[interface{}]struct{}),
	}
}

// 新建一个Set并且限制Set的值类型(例:int string bool...)
func NewSetTypeLimit(t string) *Set {
	return &Set{
		m:  make(map[interface{}]struct{}),
		tp: strings.ToLower(t),
	}
}

// 清空Set中所有的元素
func (this *Set) Clear() {
	this.m = make(map[interface{}]struct{})
}

// 获取一个Set的长度
func (this *Set) Size() int {
	return len(this.m)
}

// 添加一个（或多个）元素
func (this *Set) Add(v ...interface{}) {
	length := len(v)
	for i := 0; i < length; i++ {
		if this.tp != "" && this.tp != reflect.TypeOf(v[i]).String() {
			return
		}
		this.m[v[i]] = struct{}{}
	}
}

// 删除Set中的一个（或多个）元素
func (this *Set) Del(v ...interface{}) {
	length := len(v)
	for i := 0; i < length; i++ {
		if this.Exist(v[i]) {
			delete(this.m, v[i])
		}
	}
}

// 判断一个元素是否在Set中
func (this *Set) Exist(key interface{}) bool {
	if _, ok := this.m[key]; ok {
		return true
	}
	return false
}

// 判断两个Set是否相等，相等返回true
func (this *Set) Equals(set *Set) bool {
	if this.Size() != set.Size() {
		return false
	}
	for s := range set.m {
		if !this.Exist(s) {
			return false
		}
	}
	return true
}

// 判断set是否为空，空为true
func (this *Set) Empty() bool {
	if this.Size() == 0 {
		return true
	}
	return false
}

// 并集  返回一个新的Set
func (this *Set) Union(set *Set) *Set {
	nSet := NewSetTypeLimit(this.tp)
	if this.Empty() || set.Empty() {
		return nSet
	}
	if this.tp != "" && set.tp != "" && this.tp != set.tp {
		return nSet
	}
	for k, _ := range this.m {
		nSet.Add(k)
	}
	for k, _ := range set.m {
		nSet.Add(k)
	}
	return nSet
}

// 交集  返回一个新的Set
func (this *Set) Intersection(set *Set) *Set {
	nSet := NewSetTypeLimit(this.tp)
	if this.Empty() || set.Empty() {
		return nSet
	}
	if this.tp != "" && set.tp != "" && this.tp != set.tp {
		return nSet
	}
	for s := range set.m {
		if this.Exist(s) {
			nSet.Add(s)
		}
	}
	return nSet
}

// 差集(前者减去后者)  返回一个新的Set
func (this *Set) Difference(set *Set) *Set {
	nSet := NewSetTypeLimit(this.tp)
	if this.Empty() || set.Empty() {
		return nSet
	}
	if this.tp != "" && set.tp != "" && this.tp != set.tp {
		return nSet
	}
	for s := range this.m {
		if !set.Exist(s) {
			nSet.Add(s)
		}
	}
	return nSet
}

// set 转 数组
func (this *Set) ToArray() []interface{} {
	if this.Empty() {
		return nil
	}
	var arr []interface{}
	for k, _ := range this.m {
		arr = append(arr, k)
	}
	return arr
}

// set 转 int数组
func (this *Set) ToIntArray() []int {
	if this.Empty() {
		return nil
	}
	var arr []int
	for k, _ := range this.m {
		v, ok := k.(int)
		if ok {
			arr = append(arr, v)
		}
	}
	return arr
}

// set 转 string数组
func (this *Set) ToStrArray() []string {
	if this.Empty() {
		return nil
	}
	var arr []string
	for k, _ := range this.m {
		v, ok := k.(string)
		if ok {
			arr = append(arr, v)
		}
	}
	return arr
}

// 打印set
func (this *Set) Print() {
	fmt.Println(this.ToArray())
}
