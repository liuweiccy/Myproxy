package gateway

import (
	"net"
	"proxy"
	"config"
	"util"
	"log"
	"time"
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

func (proxyServer *ProxyServer) Start()  {
	local, err := net.Listen("tcp", proxyServer.Address())
	if err != nil {
		log.Panic("proxy server start error:", err)
	}
	log.Println("easyproxy server start ok")
	proxyServer.listener = local
	proxyServer.on = true
	// 检测心跳
	proxyServer.heartBeat()
	// 监听连接并转发连接
	for proxyServer.on {
		con, err := proxyServer.listener.Accept()
		if err == nil {
			go proxyServer.proxy.Dispatch(con)
		} else {
			log.Println("client connect server error:", err)
		}
	}
	defer proxyServer.listener.Close()
}

func (proxyServer ProxyServer)Address() string {
	return util.HostAndPortToAddress(proxyServer.host, proxyServer.port)
}

// 更具配置的固定时间，心跳检测
func (proxyServer *ProxyServer) heartBeat() {
	ticker := time.NewTicker(time.Second * time.Duration(proxyServer.beattime))
	go func() {
		for {
			select {
			case <-ticker.C:
				proxyServer.proxy.Check()
			}
		}
	}()
}
func (proxyServer *ProxyServer) Stop() {
    proxyServer.on = false
    proxyServer.proxy.Close()
    proxyServer.listener.Close()
    log.Println("easyproxy server stop ok")
}

