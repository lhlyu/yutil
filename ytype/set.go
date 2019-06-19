package ytype

import (
	"fmt"
	"reflect"
	"strings"
)

type Set struct {
	m   map[interface{}]struct{}
	tp  string                        // 限制类型，默认不限制
}

type Seter interface {
	Clear()                                           // 清空
	Size()                       int                  // 尺寸

	Add(v ...interface{})                             // 添加
	Del(v ...interface{})                             // 删除

	Exist(key interface{})       bool                 // 是否存在
	Equals(set *Set)             bool                 // 判断两个集合是否相等
	Empty()                      bool                 // 是否为空

	Union(set *Set)              *Set                  // 并集
	Intersection(set *Set)       *Set                  // 交集
	Difference(set *Set)         *Set                  // 差集

	ToArray()                    []interface{}         // 转数组
	Print()                                            // 打印
}


func NewSet() *Set{
	return &Set{
		m: make(map[interface{}]struct{}),
	}
}

// 限制Set的值类型
func NewSetTypeLimit(t string) *Set{
	return &Set{
		m: make(map[interface{}]struct{}),
		tp:strings.ToLower(t),
	}
}


func (this *Set) Clear(){
	this.m = make(map[interface{}]struct{})
}

func (this *Set) Size() int{
	return len(this.m)
}

func (this *Set) Add(v ...interface{})  {
	length := len(v)
	for i := 0; i < length ; i++ {
		if this.tp != "" && this.tp != reflect.TypeOf(v[i]).String(){
			return
		}
		this.m[v[i]] = struct{}{}
	}
}

func (this *Set) Del(v ...interface{}) {
	length := len(v)
	for i := 0; i < length ; i++ {
		if this.Exist(v[i]){
			delete(this.m,v[i])
		}
	}
}

func (this *Set) Exist(key interface{}) bool {
	if _,ok := this.m[key];ok{
		return true
	}
	return false
}

func (this *Set) Equals(set *Set) bool {
	if this.Size() != set.Size(){
		return false
	}
	for s := range set.m{
		if !this.Exist(s){
			return false
		}
	}
	return true
}

func (this *Set) Empty() bool {
	if this.Size() == 0{
		return true
	}
	return false
}


func (this *Set) Union(set *Set) *Set {
	nSet := NewSetTypeLimit(this.tp)
	if this.Empty() || set.Empty(){
		return nSet
	}
	if this.tp != "" && set.tp != "" && this.tp != set.tp{
		return nSet
	}
	for k,_ := range this.m{
		nSet.Add(k)
	}
	for k,_ := range set.m{
		nSet.Add(k)
	}
	return nSet
}

// 交集
func (this *Set) Intersection(set *Set) *Set {
	nSet := NewSetTypeLimit(this.tp)
	if this.Empty() || set.Empty(){
		return nSet
	}
	if this.tp != "" && set.tp != "" && this.tp != set.tp{
		return nSet
	}
	for s := range set.m{
		if this.Exist(s){
			nSet.Add(s)
		}
	}
	return nSet
}

// 差集 this - set
func (this *Set) Difference(set *Set) *Set {
	nSet := NewSetTypeLimit(this.tp)
	if this.Empty() || set.Empty(){
		return nSet
	}
	if this.tp != "" && set.tp != "" && this.tp != set.tp{
		return nSet
	}
	for s := range this.m{
		if !set.Exist(s){
			nSet.Add(s)
		}
	}
	return nSet
}


func (this *Set) ToArray() []interface{}{
	if this.Empty(){
		return nil
	}
	var arr []interface{}
	for k,_ := range this.m{
		arr = append(arr,k)
	}
	return arr
}

func (this *Set) Print(){
	fmt.Println(this.ToArray())
}