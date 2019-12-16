package yfunc

import (
	"strings"
)

func CompareVersion(v1,v2 string) int{
	v1 = versionHandler(v1)
	v2 = versionHandler(v2)
	vArr1 := versionSplit(v1)
	vArr2 := versionSplit(v2)
	length := len(vArr1)
	if len(vArr2) > length{
		length = len(vArr2)
	}
	vArrN1 := make([]string,length)
	vArrN2 := make([]string,length)
	copy(vArrN1,vArr1)
	copy(vArrN2,vArr2)
	for i:=0;i<length;i++{
		if vArrN1[i] == ""{
			vArrN1[i] = "0"
		}
		if vArrN2[i] == ""{
			vArrN2[i] = "0"
		}
		result := strings.Compare(vArrN1[i],vArrN2[i])
		if result == 0{
			continue
		}
		return result
	}
	return 0
}

func versionHandler(v string) string{
	v = strings.ToLower(v)
	v = strings.TrimSpace(v)
	v = strings.TrimLeft(v,"v")
	return v
}

func versionSplit(v string) []string{
	return strings.Split(v,".")
}