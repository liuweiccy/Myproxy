package gateway

import (
	"net"
	"proxy"
	"config"
)

type ProxyServer struct {
	host		string
	port		uint16
	beattime	uint16
	listener	net.Listener
	proxy       proxy.Proxy
	on	        bool
}

func (proxyServer *ProxyServer) Init(config *config.Config)  {
	proxyServer.host = config.Host
	proxyServer.port = config.Port
	proxyServer.beattime = config.Heartbeat
	proxyServer.on = false
	proxyServer.setProxy(config)
}
func (proxyServer *ProxyServer) setProxy(config *config.Config) {
	proxyServer.proxy = new(proxy.EasyProxy)
	proxyServer.proxy.Init(config)
}

