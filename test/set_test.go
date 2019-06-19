package test

import (
	"fmt"
	"testing"
	"yutil/ytype"
)

func TestSet(t *testing.T) {
	set := ytype.NewSet()
	set.Add("1")
	set.Add("2")
	set.Add("3")
	fmt.Println(set.Size())
	set.Del("3")
	fmt.Println(set.Size())
	fmt.Println(set.Exist("4"))
	fmt.Println(set.Exist("2"))
	set.Clear()
	fmt.Println(set.Size(),set.Empty())
	set.Add("1")
	set.Add("2")
	set.Add("3")
	set2 := ytype.NewSet()
	set2.Add("3")
	set2.Add("5")
	set2.Add("6")
	unset := set.Union(set2)
	unset.Print()
	set3 := ytype.NewSetTypeLimit("int")
	set3.Add("3")
	set3.Add("5")
	set3.Add(1)
	set3.Print()
}
