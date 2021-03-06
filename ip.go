package yutil

import (
	"encoding/binary"
	"net"
	"net/http"
)

const (
	_XRealIP       = "X-Real-IP"
	_XForwardedFor = "X-Forwarded-For"
)

// 获取客户端Ip
func (yIp) ClientIp(req *http.Request) string {
	var err error
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get(_XRealIP); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get(_XForwardedFor); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, err = net.SplitHostPort(remoteAddr)
	}
	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}
	_arc.log("ip", "ClientIp", err)
	return remoteAddr
}

// 将Ip转成uint32
func (yIp) Tolong(ipstr string) uint32 {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip)
}
