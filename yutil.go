package yutil

import (
	"log"
	"sync"
	"time"
)

const _version = "v1.0.7"

const logTmplate = "[%v] yutil.%s.%s:%s\n"

type conf struct {
	ignore bool // 是否忽略错误,默认忽略 ，如果设置为 false ,错误将会以log方法打印
	mx     sync.Mutex
}

var _conf = &conf{
	ignore: true,
}

func (c *conf) setIgnore(ignore bool) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.ignore = ignore
}

func (c *conf) log(fileName, funcName string, err error) {
	if c.ignore {
		return
	}
	if err != nil {
		log.Printf(logTmplate, time.Now().Format(DEFAULT_TIME_FORMAT), fileName, funcName, err.Error())
	}
}

// 是否不忽略错误
func NotIgnore() {
	_conf.setIgnore(false)
}
