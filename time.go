package yutil

import (
	"log"
	"time"
)

// 获取当前时间并格式化成 2006-01-02 15:04:05
func (yTime) Now() string {
	return time.Now().Format(DEFAULT_TIME_FORMAT)
}

// 获取当前时间并格式化成 2006-01-02
func (yTime) NowDate() string {
	return time.Now().Format(DEFAULT_DATE_FORMAT)
}

// 将2006-01-02 15:04:05字符串转换成时间
func (yTime) Parse(s string) time.Time {
	stamp, err := time.ParseInLocation(DEFAULT_TIME_FORMAT, s, time.Local)
	_arc.log("time", "TimeTimeParse", err)
	return stamp
}

// 将2006-01-02字符串转换成时间
func (yTime) DateParse(s string) time.Time {
	stamp, err := time.ParseInLocation(DEFAULT_DATE_FORMAT, s, time.Local)
	_arc.log("time", "TimeDateParse", err)
	return stamp
}

// 打印时间间隔 ,params任意参数都会导致函数重置
func (yTime) Interval(params ...interface{}) {
	f := timeInterval()
	f(params...)
}

func timeInterval() func(params ...interface{}) {
	start := int64(0)
	index := 1
	return func(params ...interface{}) {
		if len(params) > 0 {
			start = 0
			index = 1
		}
		if start == 0 {
			start = time.Now().UnixNano()
		} else {
			interval := time.Now().UnixNano() - start
			log.Printf("%d.cost: %fs\n", index, float64(interval)/1e9)
			index += 1
		}
		return
	}
}
