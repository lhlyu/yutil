package yutil

import (
	"log"
	"sync"
	"time"
)

const _version = "v2.0.2"

const logTmplate = "[%v] yutil.%s.%s:%s\n"

type arc struct {
	ignore  bool                                       // 是否忽略错误,默认忽略 ，如果设置为 false ,错误将会以log方法打印
	handler func(fileName, funcName string, err error) // 自定义处理错误，如果为空则默认使用log包
	mx      sync.Mutex
}

func new() *arc {
	a := &arc{}
	a.ignore = true
	a.handler = func(fileName, funcName string, err error) {
		if a.ignore {
			return
		}
		if err != nil {
			log.Printf(logTmplate, time.Now().Format(DEFAULT_TIME_FORMAT), fileName, funcName, err.Error())
		}
	}
	return a
}

var _arc = new()

func (c *arc) setIgnore(ignore bool) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.ignore = ignore
}

func (c *arc) sethandler(f func(fileName, funcName string, err error)) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.handler = f
}

func (c *arc) log(fileName, funcName string, err error) {
	c.handler(fileName, funcName, err)
}

// 是否不忽略错误
func NotIgnore() {
	_arc.setIgnore(false)
}

// 自定义处理错误
func SetHandler(f func(fileName, funcName string, err error)) {
	_arc.sethandler(f)
}

func Version() string {
	return _version
}

// 工具分块
type (
	yBase64 struct{}
	yMd5    struct{}
	ySha1   struct{}
	yJson   struct{}

	yRegexp struct{}

	yFile struct{}

	yTime   struct{}
	yString struct{}
	yRandom struct{}

	yIp struct{}

	yRuntime struct{}
	yConvert struct{}
	ySlice   struct{}
	yOther   struct{}
)

var (
	Base64 = yBase64{}
	Md5    = yMd5{}
	Sha1   = ySha1{}

	Json   = yJson{}
	Regexp = yRegexp{}

	File    = yFile{}
	Time    = yTime{}
	String  = yString{}
	Random  = yRandom{}
	IP      = yIp{}
	Runtime = yRuntime{}
	Convert = yConvert{}
	Slice   = ySlice{}
	Other   = yOther{}
)
