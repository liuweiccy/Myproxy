package proxy

import (
	"config"
	"net"
)

// 代理接口
// 定义了 初始化、检测、清理、恢复、转发、关闭
type Proxy interface {
	Init(config *config.Config)
	Check()
	Clean(url string)
	Recover(url string)
	Dispatch(con net.Conn)
	Close()
}

